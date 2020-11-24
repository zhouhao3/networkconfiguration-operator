package controllers

import (
	"context"
	"time"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"github.com/metal3-io/networkconfiguration-operator/pkg/machine"
	"github.com/metal3-io/networkconfiguration-operator/pkg/util/finalizer"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const finalizerKey string = "metal3.io.v1alpha1"

// NoneHandler ...
func (r *NetworkBindingReconciler) NoneHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	// Add finalizer
	err := finalizer.AddFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err == nil {
		result.Requeue = true
	}

	return v1alpha1.NetworkBindingCreated, result, err
}

// CreateHandler ...
func (r *NetworkBindingReconciler) CreateHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Configure network

	return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true}, nil
}

// ConfiguringHandler ...
func (r *NetworkBindingReconciler) ConfiguringHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Check network has been configured or not
	var configureResutl string
	switch configureResutl {
	case "success":
		return v1alpha1.NetworkBindingConfigured, ctrl.Result{Requeue: false}, nil
	case "failed":
		return v1alpha1.NetworkBindingConfigureFailed, ctrl.Result{Requeue: true}, nil
	}

	return v1alpha1.NetworkBindingConfiguring, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
}

// ConfiguredHandler ...
func (r *NetworkBindingReconciler) ConfiguredHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Delete network

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true}, nil
}

// ConfigureFailedHandler ...
func (r *NetworkBindingReconciler) ConfigureFailedHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Reconfigure network

	return v1alpha1.NetworkBindingCreated, ctrl.Result{Requeue: true}, nil
}

// DeletingHandler ...
func (r *NetworkBindingReconciler) DeletingHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Check network has been deleted or not
	var deleteResutl string
	switch deleteResutl {
	case "success":
		return v1alpha1.NetworkBindingDeleted, ctrl.Result{Requeue: true}, nil
	case "failed":
		// Delete network again
	}

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{Requeue: true, RequeueAfter: time.Second * 10}, nil
}

// DeletedHandler ...
func (r *NetworkBindingReconciler) DeletedHandler(ctx context.Context, info *machine.Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	i := instance.(*v1alpha1.NetworkBinding)

	// Remove finalizers
	err := finalizer.RemoveFinalizer(&i.Finalizers, finalizerKey)
	result := reconcile.Result{}
	if err != nil {
		result.Requeue = true
	}

	return v1alpha1.NetworkBindingDeleted, result, err
}
