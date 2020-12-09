package device

import (
	"context"
	"errors"

	"github.com/metal3-io/networkconfiguration-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PortState ...
type PortState string

const (
	// NotConfigured ...
	NotConfigured PortState = "not configured"

	// Configuring ...
	Configuring PortState = "configuring"

	// Configured ...
	Configured PortState = "configured"

	// ConfigureFailed ...
	ConfigureFailed PortState = "configure failed"

	// Deleting ...
	Deleting PortState = "deleting"

	// Deleted ...
	Deleted PortState = "deleted"

	// DeleteFailed ...
	DeleteFailed PortState = "delete failed"
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
	ConfigurePort(ctx context.Context, networkConfiguration *v1alpha1.NetworkConfiguration, port *v1alpha1.NetworkBindingSpecPort) error

	// DeConfigurePort remove the network configure from the port
	DeConfigurePort(ctx context.Context, port *v1alpha1.NetworkBindingSpecPort) error

	// PortState return the port's state of the device
	PortState(ctx context.Context, portID string) PortState
}
