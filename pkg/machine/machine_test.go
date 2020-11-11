package machine

import (
	"testing"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type testInstance struct {
	state v1alpha1.StateType
}

func (t *testInstance) GetState() v1alpha1.StateType {
	return t.state
}

func (t *testInstance) SetState(state v1alpha1.StateType) {
	t.state = state
}

var out string

func handlerTest0(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	out = "Hello"
	return "test1", ctrl.Result{}, nil
}

func handlerTest1(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	out += " world"
	return "test2", ctrl.Result{}, nil
}

func handlerTest2(client *client.Client, instance interface{}) (nextState v1alpha1.StateType, result ctrl.Result, err error) {
	out += "!"
	return "", ctrl.Result{}, nil
}

func TestMachine(t *testing.T) {
	defer func() {
		out = ""
	}()
	var instance testInstance
	m := New(
		nil,
		&instance,
		&Handlers{
			"":      handlerTest0,
			"test1": handlerTest1,
			"test2": handlerTest2,
		},
	)
	m.Reconcile()
	if out != "Hello" {
		t.Fatal(out)
	}
	m.Reconcile()
	if out != "Hello world" {
		t.Fatal(out)
	}
	m.Reconcile()
	if out != "Hello world!" {
		t.Fatal(out)
	}
}
