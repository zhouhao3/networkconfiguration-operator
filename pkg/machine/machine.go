package machine

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Handler ...
type Handler func(ctx context.Context, info *Information, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error)

// Handlers ...
type Handlers map[v1alpha1.StateType]Handler

// Instance ...
type Instance interface {
	GetState() v1alpha1.StateType
	SetState(state v1alpha1.StateType)
}

// Information ...
type Information struct {
	Client *client.Client
	Logger *logr.Logger
}

// Machine ...
type Machine struct {
	ctx      context.Context
	info     *Information
	instance Instance
	handlers *Handlers
}

// New create state machine
func New(ctx context.Context, info *Information, instance Instance, handlers *Handlers) Machine {
	return Machine{
		ctx:      ctx,
		info:     info,
		instance: instance,
		handlers: handlers,
	}
}

// Reconcile ...
func (m *Machine) Reconcile() (ctrl.Result, error) {
	nextState, result, err := (*m.handlers)[m.instance.GetState()](m.ctx, m.info, m.instance)
	m.instance.SetState(nextState)
	return result, err
}
