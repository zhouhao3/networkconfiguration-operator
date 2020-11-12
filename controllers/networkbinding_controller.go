/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"github.com/metal3-io/networkconfiguration-operator/pkg/machine"
	"github.com/metal3-io/networkconfiguration-operator/pkg/util/finalizer"
)

const finalizerKey string = "metal3.io.v1alpha1"

// NetworkBindingReconciler reconciles a NetworkBinding object
type NetworkBindingReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=metal3.io.my.domain,resources=networkbindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=metal3.io.my.domain,resources=networkbindings/status,verbs=get;update;patch

// SetupWithManager ...
func (r *NetworkBindingReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.NetworkBinding{}).
		Complete(r)
}

// Reconcile ...
func (r *NetworkBindingReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("networkbinding", req.NamespacedName)

	// Fetch the instance
	instance := &v1alpha1.NetworkBinding{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	m := machine.New(
		&r.Client,
		instance,
		&machine.Handlers{
			v1alpha1.NetworkBindingNone:            r.NoneHandler,
			v1alpha1.NetworkBindingCreated:         r.CreateHandler,
			v1alpha1.NetworkBindingConfiguring:     r.ConfiguringHandler,
			v1alpha1.NetworkBindingConfigured:      r.ConfiguredHandler,
			v1alpha1.NetworkBindingConfigureFailed: r.ConfigureFailedHandler,
			v1alpha1.NetworkBindingDeleting:        r.DeletingHandler,
		},
	)

	switch {
	// On object created
	case instance.DeletionTimestamp.IsZero() && len(instance.Finalizers) == 0:
		// Add finalizer
		err = finalizer.AddFinalizer(&instance.Finalizers, finalizerKey)
		// Do something
		_, _ = m.Reconcile()
	// On object updated
	case instance.DeletionTimestamp.IsZero() && len(instance.Finalizers) != 0:
		// Do something

	// On object delete
	case !instance.DeletionTimestamp.IsZero():
		// Remove finalizers, you can do something by finalizer hook
		err = finalizer.RemoveFinalizer(&instance, &instance.Finalizers, finalizerKey)
	}

	// Update object
	err = r.Update(context.Background(), instance)

	return ctrl.Result{}, err
}

// NoneHandler ...
func (r *NetworkBindingReconciler) NoneHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Set Network Configure
	return v1alpha1.NetworkBindingCreated, ctrl.Result{RequeueAfter: 10}, nil
}

// CreateHandler ...
func (r *NetworkBindingReconciler) CreateHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// Set Network Configure
	return v1alpha1.NetworkBindingConfiguring, ctrl.Result{RequeueAfter: 10}, nil
}

// ConfiguringHandler ...
func (r *NetworkBindingReconciler) ConfiguringHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	return v1alpha1.NetworkBindingConfigured, ctrl.Result{}, nil
}

// ConfiguredHandler ...
func (r *NetworkBindingReconciler) ConfiguredHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	return v1alpha1.NetworkBindingDeleting, ctrl.Result{RequeueAfter: 10}, nil
}

// ConfigureFailedHandler ...
func (r *NetworkBindingReconciler) ConfigureFailedHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	// 删除Device CR上的配置
	// 判断配置完成
	// 若牌子未完成 return v1alpha1.ConfigureFailed, ctrl.Result{RequeueAfter: 10}, nil
	return v1alpha1.NetworkBindingCreated, ctrl.Result{RequeueAfter: 10}, nil
}

// DeletingHandler ...
func (r *NetworkBindingReconciler) DeletingHandler(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	_ = instance.(*v1alpha1.NetworkBinding)

	return v1alpha1.NetworkBindingNone, ctrl.Result{}, nil
}
