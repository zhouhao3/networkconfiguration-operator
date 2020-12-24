package device

import (
	"context"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// newSwitch ...
func newSwitch(ctx context.Context, client client.Client, deviceRef *v1alpha1.DeviceRef) (*Switch, error) {
	var instance interface{}

	// Get SwitchDevice CR
	instance, err := deviceRef.Fetch(ctx, client)
	if err != nil {
		return nil, err
	}

	return &Switch{
		client:   client,
		instance: instance.(v1alpha1.Switch),
	}, nil
}

// Switch is a kind of network device
type Switch struct {
	client   client.Client
	instance v1alpha1.Switch
}

// ConfigurePort set the network configure to the port
func (s *Switch) ConfigurePort(ctx context.Context, configuration interface{}, portID string) error {
	return nil
}

// DeConfigurePort remove the network configure from the port
func (s *Switch) DeConfigurePort(ctx context.Context, portID string) error {
	return nil
}

// PortState return the port's state of the device
func (s *Switch) PortState(ctx context.Context, portID string) PortState {
	return None
}

// CheckPortConfigutation checks whether the configuration is configured on the port
func (s *Switch) CheckPortConfigutation(ctx context.Context, configuration interface{}, portID string) bool {
	return false
}
