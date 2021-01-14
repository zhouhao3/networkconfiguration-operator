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

// noneHandler will be called when Port is created
func (r *PortReconciler) noneHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	// Add finalizer
	err := finalizer.AddFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err == nil {
		result.Requeue = true
	}

	return v1alpha1.PortCreating, ctrl.Result{Requeue: true}, err
}

// createHandler will be called when Port was created
func (r *PortReconciler) creatingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	// Initialize device
	dev, err := device.New(ctx, info.Client, &i.Spec.DeviceRef)
	if err != nil {
		return v1alpha1.PortCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	// Get port's state
	switch dev.PortState(ctx, i.Spec.PortID) {
	case device.None, device.Deleted:
		// Go to `Configuring` state
		return v1alpha1.PortConfiguring, ctrl.Result{Requeue: true}, nil

	case device.Deleting:
		// Just wait

	default:
		return v1alpha1.PortCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, fmt.Errorf("port(%s) have been used", i.Spec.PortID)
	}

	return v1alpha1.PortCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
}

// configuringHandler will be called when configuring network
func (r *PortReconciler) configuringHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	dev, err := device.New(ctx, info.Client, &i.Spec.DeviceRef)
	if err != nil {
		return v1alpha1.PortConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(ctx, i.Spec.PortID) {
	case device.Configured:
		// If configure network success, we just need to set next state to `Configured`, but not Reconcile
		return v1alpha1.PortConfigured, ctrl.Result{Requeue: false}, nil

	case device.None, device.Deleted, device.ConfigureFailed:
		// Fetch network configuration
		configuration, err := i.Spec.PortConfigurationRef.Fetch(ctx, info.Client)
		if err != nil {
			return v1alpha1.PortConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
		}
		// Configure network
		err = dev.ConfigurePort(ctx, configuration, i.Spec.PortID)

	default:
		// Just wait
	}

	return v1alpha1.PortConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
}

// configuredHandler will be called when the user want to delete the network configuration for the port be configured
func (r *PortReconciler) configuredHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	// Reconfigure port when user update CR.
	if i.DeletionTimestamp.IsZero() {
		return v1alpha1.PortConfiguring, ctrl.Result{Requeue: true}, nil
	}

	return v1alpha1.PortDeleting, ctrl.Result{Requeue: true}, nil
}

// deletingHandler will be called when deleting network configuration
func (r *PortReconciler) deletingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	dev, err := device.New(ctx, info.Client, &i.Spec.DeviceRef)
	if err != nil {
		return v1alpha1.PortDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(ctx, i.Spec.PortID) {
	case device.None, device.Deleted:
		return v1alpha1.PortDeleted, ctrl.Result{Requeue: true}, nil

	case device.Configured, device.ConfigureFailed, device.DeleteFailed:
		// Delete network
		err = dev.DeConfigurePort(ctx, i.Spec.PortID)

	default:
		// Just wait
	}

	return v1alpha1.PortDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
}

// deletedHandler will be called when the network configuration has been deleted
func (r *PortReconciler) deletedHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.Port)

	// Remove finalizer
	err := finalizer.RemoveFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err != nil {
		result.Requeue = true
	}

	return v1alpha1.PortDeleted, result, err
}
