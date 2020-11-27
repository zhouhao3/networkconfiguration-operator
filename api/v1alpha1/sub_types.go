package v1alpha1

import (
	"context"
	"errors"
	"reflect"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// StateType is the type of .status.state
type StateType string

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

// NetworkBindingRef is the reference for NetworkBinding CR
type NetworkBindingRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
}

// Fetch the instance
func (n *NetworkBindingRef) Fetch(client client.Client) (*NetworkBinding, error) {
	var instance NetworkBinding
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

	return nil, errors.New("can't find instance for the kind")
}

// NICHint ...
type NICHint struct {
	Speed    uint `json:"speed,omitempty"`
	SmartNIC bool `json:"smartNIC,omitempty"`
}

// Assess nicHints and return a score.
func (n NICHint) Assess(nicHint NICHint) float64 {
	rsv := reflect.ValueOf(n)
	rav := reflect.ValueOf(nicHint)

	score := 0.0
	for i := 0; i < rsv.NumField(); i++ {
		switch rsv.Field(i).Kind() {
		case reflect.Bool:
			if rsv.Field(i).Bool() && !rav.Field(i).Bool() {
				return 0
			} else if rsv.Field(i).Bool() == rav.Field(i).Bool() {
				score += 10
			}

		case reflect.String:
			if rsv.Field(i).String() != rav.Field(i).String() {
				return 0
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if rsv.Field(i).Int() > rav.Field(i).Int() {
				return 0
			}
			score += float64(rsv.Field(i).Int()) / float64(rav.Field(i).Int()) * 100

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if rsv.Field(i).Uint() > rav.Field(i).Uint() {
				return 0
			}
			score += float64(rsv.Field(i).Uint()) / float64(rav.Field(i).Uint()) * 100

		case reflect.Float32, reflect.Float64:
			if rsv.Field(i).Float() > rav.Field(i).Float() {
				return 0
			}
			score += rsv.Field(i).Float() / rav.Field(i).Float() * 100
		}
	}

	return score
}
