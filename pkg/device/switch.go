package device

import (
	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// newSwitch ...
func newSwitch(client *client.Client, deviceRef *v1alpha1.DeviceRef) (*Switch, error) {
	var instance interface{}

	// Get SwitchDevice CR
	instance, err := deviceRef.Fetch(client)
	if err != nil {
		return nil, err
	}

	return &Switch{
		instance: instance.(v1alpha1.Switch),
	}, nil
}

// Switch is a kind of network device
type Switch struct {
	instance v1alpha1.Switch
}

// ConfigurePort set the network configure to the port
func (s *Switch) ConfigurePort(port v1alpha1.NetworkBindingSpecPort) error {
	return nil
}

// DeConfigurePort remove the network configure from the port
func (s *Switch) DeConfigurePort(port v1alpha1.NetworkBindingSpecPort) error {
	return nil
}

// PortState return the port's state of the device
func (s *Switch) PortState(portID string) v1alpha1.StateType {
	return ""
}
