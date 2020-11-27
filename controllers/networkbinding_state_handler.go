package controllers

import (
	"context"
	"errors"
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

	dev, err := device.New(&info.Client, &i.Spec.Port.DeviceRef)
	if err != nil {
		return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(i.Spec.Port.PortID) {
	case "not found":
		return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true}, nil

	case "deleting":
		// Just wait

	default:
		return v1alpha1.NetworkBindingCreating, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, errors.New("port have been used")
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

	switch dev.PortState(i.Spec.Port.PortID) {
	case "configure success":
		// If configure network success, we just need to set next state to configured, but not Reconcile
		return v1alpha1.NetworkBindingConfigured, ctrl.Result{Requeue: false}, nil

	case "not found":
		dev.ConfigurePort()
		// Configure network

	case "configure failed":
		// Configure network again
		return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 120}, nil

	default:
		// Just wait
	}

	return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
}

// configuredHandler will be called when the user want to delete the network configuration for the port be configured
func (r *NetworkBindingReconciler) configuredHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true}, nil
}

// deletingHandler will be called when deleting network configuration
func (r *NetworkBindingReconciler) deletingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	dev, err := device.New(&info.Client, &i.Spec.Port.DeviceRef)
	if err != nil {
		return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, err
	}

	switch dev.PortState(i.Spec.Port.PortID) {
	case "deleting success":
		return v1alpha1.NetworkBindingDeleted, ctrl.Result{Requeue: true}, nil

	case "configure success":
		// Delete network

	case "delete failed":
		// Delete network again
		return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 120}, nil

	default:
		// Just wait
	}

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
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
