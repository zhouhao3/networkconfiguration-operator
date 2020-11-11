package device

import (
	"errors"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// New ...
func New(client *client.Client, deviceRef *v1alpha1.DeviceRef) (Device, error) {
	switch deviceRef.Kind {
	case "Switch":
		return newSwitch(client, deviceRef)
	}

	return nil, errors.New("")
}

// Device ...
type Device interface {
	ConfigurePort()
	DeConfigurePort()
}
