package v1alpha1

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NetworkConfigurationRef is the reference for NetworkConfiguration CR
type NetworkConfigurationRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
}

// Fetch the instance
func (n *NetworkConfigurationRef) Fetch(client client.Client) (*NetworkConfiguration, error) {
	var instance NetworkConfiguration
	err := client.Get(
		context.Background(),
		types.NamespacedName{
			Name:      n.Name,
			Namespace: n.NameSpace,
		},
		&instance,
	)
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

// DeviceRef is the reference for Device CR
type DeviceRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
	Kind      string `json:"kind"`
}

// Fetch the instance
func (d *DeviceRef) Fetch(client *client.Client) (interface{}, error) {

	switch d.Kind {
	case "Switch":
		var instance Switch
		err := (*client).Get(
			context.Background(),
			types.NamespacedName{
				Name:      d.Name,
				Namespace: d.NameSpace,
			},
			&instance,
		)
		if err != nil {
			return nil, err
		}
		return instance, nil
	}

	return nil, errors.New("")
}

// StateType is the type of .status.state
type StateType string
