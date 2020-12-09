package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"github.com/metal3-io/networkconfiguration-operator/pkg/device"
	"github.com/metal3-io/networkconfiguration-operator/pkg/machine"
	"github.com/metal3-io/networkconfiguration-operator/pkg/util/finalizer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const finalizerKey string = "metal3.io.v1alpha1"

// noneHandler will be called when networkBinding is created
func (r *NetworkBindingReconciler) noneHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	// Add finalizer
	err := finalizer.AddFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err == nil {
		result.Requeue = true
	}

	return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true}, err
}

// createHandler will be called when networkBinding was created
func (r *NetworkBindingReconciler) creatingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	// Initialize device
	dev, err := device.New(&info.Client, &i.Spec.Port.DeviceRef)
	if err != nil {
		return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	// Get port's state
	switch dev.PortState(ctx, i.Spec.Port.PortID) {
	case device.NotConfigured, device.Deleted:
		// Go to `Configuring` state
		return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true}, nil

	case device.Deleting:
		// Just wait

	default:
		return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, fmt.Errorf("port(%s) have been used", i.Spec.Port.PortID)
	}

	return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
}

// configuringHandler will be called when configuring network
func (r *NetworkBindingReconciler) configuringHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	dev, err := device.New(&info.Client, &i.Spec.Port.DeviceRef)
	if err != nil {
		return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(ctx, i.Spec.Port.PortID) {
	case device.Configured:
		// If configure network success, we just need to set next state to `Configured`, but not Reconcile
		return v1alpha1.NetworkBindingConfigured, ctrl.Result{Requeue: false}, nil

	case device.NotConfigured, device.Deleted, device.ConfigureFailed:
		// Fetch network configuration
		networkConfiguration, err := i.Spec.NetworkConfigurationRef.Fetch(info.Client)
		if err != nil {
			return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
		}
		// Configure network
		err = dev.ConfigurePort(ctx, networkConfiguration, &i.Spec.Port)

	default:
		// Just wait
	}

	return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
}

// configuredHandler will be called when the user want to delete the network configuration for the port be configured
func (r *NetworkBindingReconciler) configuredHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// `Configured` state just show user: this port has been configured

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true}, nil
}

// deletingHandler will be called when deleting network configuration
func (r *NetworkBindingReconciler) deletingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	dev, err := device.New(&info.Client, &i.Spec.Port.DeviceRef)
	if err != nil {
		return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(ctx, i.Spec.Port.PortID) {
	case device.NotConfigured, device.Deleted:
		return v1alpha1.NetworkBindingDeleted, ctrl.Result{Requeue: true}, nil

	case device.Configured, device.ConfigureFailed, device.DeleteFailed:
		// Delete network
		err = dev.DeConfigurePort(ctx, &i.Spec.Port)

	default:
		// Just wait
	}

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
}

// deletedHandler will be called when the network configuration has been deleted
func (r *NetworkBindingReconciler) deletedHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	// Remove finalizer
	err := finalizer.RemoveFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err != nil {
		result.Requeue = true
	}

	return v1alpha1.NetworkBindingDeleted, result, err
}
