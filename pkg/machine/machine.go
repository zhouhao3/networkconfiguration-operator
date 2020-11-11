package machine

import (
	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Handler ...
type Handler func(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error)

// Handlers ...
type Handlers map[v1alpha1.StateType]Handler

// Object ...
type Object interface {
	GetState() v1alpha1.StateType
	SetState(state v1alpha1.StateType)
}

// Machine ...
type Machine struct {
	client   *client.Client
	instance Object
	handlers *Handlers
}

// New create state machine
func New(client *client.Client, instance Object, handlers *Handlers) Machine {
	return Machine{
		client:   client,
		instance: instance,
		handlers: handlers,
	}
}

// Reconcile ...
func (m *Machine) Reconcile() (ctrl.Result, error) {
	nextState, result, err := (*m.handlers)[m.instance.GetState()](m.client, m.instance)
	m.instance.SetState(nextState)
	return result, err
}
