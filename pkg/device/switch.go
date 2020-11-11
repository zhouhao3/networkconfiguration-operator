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

// Switch ...
type Switch struct {
	instance v1alpha1.Switch
}

// ConfigurePort ...
func (s *Switch) ConfigurePort() {

}

// DeConfigurePort ...
func (s *Switch) DeConfigurePort() {

}
