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
)

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
func (r *NetworkBindingReconciler) Reconcile(req ctrl.Request) (result ctrl.Result, err error) {
	_ = context.Background()
	_ = r.Log.WithValues("networkbinding", req.NamespacedName)

	// Fetch the instance
	instance := &v1alpha1.NetworkBinding{}
	err = r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		// Error reading the object - requeue the request
		return reconcile.Result{}, err
	}

	// Initialize state machine
	m := machine.New(
		&machine.Information{
			Client: r.Client,
			Logger: r.Log,
		},
		instance,
		&machine.Handlers{
			v1alpha1.NetworkBindingNone:        r.noneHandler,
			v1alpha1.NetworkBindingCreating:    r.creatingHandler,
			v1alpha1.NetworkBindingConfiguring: r.configuringHandler,
			v1alpha1.NetworkBindingConfigured:  r.configuredHandler,
			v1alpha1.NetworkBindingDeleting:    r.deletingHandler,
			v1alpha1.NetworkBindingDeleted:     r.deletedHandler,
		},
	)

	var merr machine.Error
	switch {
	// On object created
	case instance.DeletionTimestamp.IsZero() && len(instance.Finalizers) == 0:
		// Reconcile state
		result, merr = m.Reconcile(context.TODO())
		if merr != nil {
			err = merr.Error()
			switch merr.Type() {
			case machine.ReconcileError:
				// Do something
			case machine.HandlerError:
				// Do something
			}
		}

	// On object updated
	case instance.DeletionTimestamp.IsZero() && len(instance.Finalizers) != 0:
		// Reconcile state
		result, merr = m.Reconcile(context.TODO())
		if merr != nil {
			err = merr.Error()
			switch merr.Type() {
			case machine.ReconcileError:
				// Do something
			case machine.HandlerError:
				// Do something
			}
		}

	// On object delete
	case !instance.DeletionTimestamp.IsZero():
		instance.SetState(v1alpha1.NetworkBindingDeleting)
		// Reconcile state
		result, merr = m.Reconcile(context.TODO())
		if merr != nil {
			err = merr.Error()
			switch merr.Type() {
			case machine.ReconcileError:
				// Do something
			case machine.HandlerError:
				// Do something
			}
		}
	}

	// Update object
	err = r.Update(context.Background(), instance)

	return ctrl.Result{}, err
}
