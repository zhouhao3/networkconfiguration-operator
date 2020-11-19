package machine

import (
	"context"
	"testing"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type testInstance struct {
	out   string
	state v1alpha1.StateType
}

func (t *testInstance) GetState() v1alpha1.StateType {
	return t.state
}

func (t *testInstance) SetState(state v1alpha1.StateType) {
	t.state = state
}

func handlerTest0(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	instance.(*testInstance).out = "Hello"
	return "test1", ctrl.Result{}, nil
}

func handlerTest1(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	instance.(*testInstance).out += " world"
	return "test2", ctrl.Result{}, nil
}

func handlerTest2(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	instance.(*testInstance).out += "!"
	return "", ctrl.Result{}, nil
}

func TestMachine(t *testing.T) {
	var instance testInstance
	m := New(
		context.TODO(),
		nil,
		&instance,
		&Handlers{
			"":      handlerTest0,
			"test1": handlerTest1,
			"test2": handlerTest2,
		},
	)
	m.Reconcile()
	if instance.out != "Hello" {
		t.Fatal(instance.out)
	}
	m.Reconcile()
	if instance.out != "Hello world" {
		t.Fatal(instance.out)
	}
	m.Reconcile()
	if instance.out != "Hello world!" {
		t.Fatal(instance.out)
	}
}

func BenchmarkMachine(b *testing.B) {
	var instance testInstance
	m := New(
		context.TODO(),
		nil,
		&instance,
		&Handlers{
			"":      handlerTest0,
			"test1": handlerTest1,
			"test2": handlerTest2,
		},
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Reconcile()
	}
}

var tInstance testInstance

func handlerTest00(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	tInstance.out = "Hello"
	return "test1", ctrl.Result{}, nil
}

func handlerTest11(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	tInstance.out += " world"
	return "test2", ctrl.Result{}, nil
}

func handlerTest22(ctx context.Context, info *Information, instance interface{}) (v1alpha1.StateType, ctrl.Result, error) {
	tInstance.out += "!"
	return "", ctrl.Result{}, nil
}

func BenchmarkMachineNoAssert(b *testing.B) {
	m := New(
		context.TODO(),
		nil,
		&tInstance,
		&Handlers{
			"":      handlerTest00,
			"test1": handlerTest11,
			"test2": handlerTest22,
		},
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Reconcile()
	}
}
