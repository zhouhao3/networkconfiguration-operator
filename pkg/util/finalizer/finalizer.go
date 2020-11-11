package finalizer

import (
	"errors"

	"github.com/metal3-io/networkbinding-operator/pkg/util/stringslice"
)

type hookType func(interface{}) error

var finalizersHooks map[string]hookType

func init() {
	finalizersHooks = make(map[string]hookType)
}

// AddFinalizer create finalizer, must unique for every object.
func AddFinalizer(finalizers *[]string, finalizer string) (err error) {
	if stringslice.Contains(*finalizers, finalizer) {
		return errors.New("the finalizer of object must be unique")
	}
	*finalizers = append(*finalizers, finalizer)
	return
}

// RemoveFinalizer run finalizers hook and remove finalizer, return the result of hook.
// Note: the input of object must be a pointer.
func RemoveFinalizer(object interface{}, finalizers *[]string, finalizer string) (err error) {
	if !stringslice.Delete(finalizers, finalizer) {
		return errors.New("haven't find finalizer in finalizers")
	}

	hook, exist := finalizersHooks[finalizer]
	if !exist {
		return nil
	}

	return hook(object)
}

// SetHook set hook for finalizer.
func SetHook(finalizer string, hook hookType) (err error) {
	_, exist := finalizersHooks[finalizer]
	if exist {
		return errors.New("the hook of finalizer must be unique")
	}
	finalizersHooks[finalizer] = hook
	return
}
