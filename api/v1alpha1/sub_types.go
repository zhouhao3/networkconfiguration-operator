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
func (n *NetworkConfigurationRef) Fetch(ctx context.Context, client client.Client) (instance *NetworkConfiguration, err error) {
	err = client.Get(
		context.Background(),
		types.NamespacedName{
			Name:      n.Name,
			Namespace: n.NameSpace,
		},
		instance,
	)

	return
}

// NetworkBindingRef is the reference for NetworkBinding CR
type NetworkBindingRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
}

// Fetch the instance
func (n *NetworkBindingRef) Fetch(ctx context.Context, client client.Client) (instance *NetworkBinding, err error) {
	err = client.Get(
		ctx,
		types.NamespacedName{
			Name:      n.Name,
			Namespace: n.NameSpace,
		},
		instance,
	)

	return
}

// DeviceRef is the reference for Device CR
type DeviceRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
	Kind      string `json:"kind"`
}

// Fetch the instance
func (d *DeviceRef) Fetch(ctx context.Context, client *client.Client) (instance interface{}, err error) {
	switch d.Kind {
	case "Switch":
		err = (*client).Get(
			ctx,
			types.NamespacedName{
				Name:      d.Name,
				Namespace: d.NameSpace,
			},
			instance.(*Switch),
		)
	default:
		err = errors.New("no instance for the ref")
	}

	return
}

// NICHint ...
type NICHint struct {
	Speed    uint `json:"speed,omitempty"`
	SmartNIC bool `json:"smartNIC,omitempty"`
}

// Assess capability and return a score.
func (hint NICHint) Assess(capability NICHint) float64 {

	hintV := reflect.ValueOf(hint)
	capabilityV := reflect.ValueOf(capability)

	score := 0.0
	for i := 0; i < hintV.NumField(); i++ {
		switch hintV.Field(i).Kind() {
		case reflect.Bool:
			if hintV.Field(i).Bool() && !capabilityV.Field(i).Bool() {
				return 0
			} else if hintV.Field(i).Bool() == capabilityV.Field(i).Bool() {
				score += 10
			}

		case reflect.String:
			if hintV.Field(i).String() != capabilityV.Field(i).String() {
				return 0
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if hintV.Field(i).Int() > capabilityV.Field(i).Int() {
				return 0
			}
			score += float64(hintV.Field(i).Int()) / float64(capabilityV.Field(i).Int()) * 100

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if hintV.Field(i).Uint() > capabilityV.Field(i).Uint() {
				return 0
			}
			score += float64(hintV.Field(i).Uint()) / float64(capabilityV.Field(i).Uint()) * 100

		case reflect.Float32, reflect.Float64:
			if hintV.Field(i).Float() > capabilityV.Field(i).Float() {
				return 0
			}
			score += hintV.Field(i).Float() / capabilityV.Field(i).Float() * 100
		}
	}

	return score
}
