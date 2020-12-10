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

// Handler is a state handle function
type Handler func(ctx context.Context, info *Information, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error)

// Handlers includes a lot of handler
type Handlers map[v1alpha1.StateType]Handler

// Instance is a object for the CR need be reconcile
// NOTE: Instance must be a pointer
type Instance interface {
	GetState() v1alpha1.StateType
	SetState(state v1alpha1.StateType)
}

// Information ...
type Information struct {
	Client client.Client
	Logger logr.Logger
}

// Machine is a state machine
type Machine struct {
	info         *Information
	instance     Instance
	handlers     *Handlers
	requeueAfter time.Duration
}

// ErrorType is the error when reconcile state machine
type ErrorType string

const (
	// ReconcileError means have error when reconcile
	ReconcileError ErrorType = "reconcile error"

	// HandlerError means have error in the handler for a state
	HandlerError ErrorType = "handler error"
)

// Error include error type and error message from state machine
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

// New a state machine
// NOTE: The paramater of instance must be a pointer
func New(info *Information, instance Instance, handlers *Handlers) Machine {
	return Machine{
		info:     info,
		instance: instance,
		handlers: handlers,
	}
}

// Reconcile state machine
func (m *Machine) Reconcile(ctx context.Context) (ctrl.Result, Error) {
	handler, exist := (*m.handlers)[m.instance.GetState()]
	if !exist {
		return ctrl.Result{}, &machineError{
			errType: ReconcileError,
			err:     fmt.Errorf("no handler for the state(%v)", m.instance.GetState()),
		}
	}

	nextState, result, err := handler(ctx, m.info, m.instance)
	m.instance.SetState(nextState)
	if err != nil {
		return result, &machineError{
			errType: HandlerError,
			err:     err,
		}
	}

	return result, nil
}
