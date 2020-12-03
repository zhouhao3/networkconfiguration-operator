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
	// ConfigurePort set the network configure to the port
	ConfigurePort(port v1alpha1.NetworkBindingSpecPort) error

	// DeConfigurePort remove the network configure from the port
	DeConfigurePort(port v1alpha1.NetworkBindingSpecPort) error

	// PortState return the port's state of the device
	PortState(portID string) v1alpha1.StateType
}
