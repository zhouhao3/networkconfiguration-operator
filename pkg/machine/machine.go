package machine

import (
	"context"
	"fmt"
	"time"

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
	Client client.Client
	Logger logr.Logger
}

// Machine ...
type Machine struct {
	ctx          context.Context
	info         *Information
	instance     Instance
	handlers     *Handlers
	requeueAfter time.Duration
}

// ErrorType ...
type ErrorType string

const (
	// ReconcileError means have error when reconcile
	ReconcileError ErrorType = "reconcile error"

	// HandlerError means have error in the handler for a state
	HandlerError ErrorType = "handler error"
)

// Error ...
type Error interface {
	Type() ErrorType
	Error() error
}

type machineError struct {
	errType ErrorType
	err     error
}

func (me *machineError) Type() ErrorType {
	return me.errType
}

func (me *machineError) Error() error {
	return me.err
}

// New create state machine, the paramater of instance must be a pointer
func New(ctx context.Context, info *Information, instance Instance, handlers *Handlers) Machine {
	return Machine{
		ctx:      ctx,
		info:     info,
		instance: instance,
		handlers: handlers,
	}
}

// Reconcile ...
func (m *Machine) Reconcile() (ctrl.Result, Error) {
	handler, exist := (*m.handlers)[m.instance.GetState()]
	if !exist {
		return ctrl.Result{}, &machineError{
			errType: ReconcileError,
			err:     fmt.Errorf("no handler for %s state", m.instance.GetState()),
		}
	}

	nextState, result, err := handler(m.ctx, m.info, m.instance)
	m.instance.SetState(nextState)

	if err == nil {
		return result, nil
	}
	return result, &machineError{
		errType: HandlerError,
		err:     err,
	}
}
