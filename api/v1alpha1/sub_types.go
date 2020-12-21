package v1alpha1

import (
	"context"
	"fmt"
	"reflect"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// StateType is the type of .status.state
type StateType string

// PortRef is the reference for NetworkBinding CR
type PortRef struct {
	Name      string `json:"name"`
	NameSpace string `json:"nameSpace"`
}

// Fetch the instance
func (ref *PortRef) Fetch(ctx context.Context, client client.Client) (instance *Port, err error) {
	err = client.Get(
		ctx,
		types.NamespacedName{
			Name:      ref.Name,
			Namespace: ref.NameSpace,
		},
		instance,
	)

	return
}

// PortConfigurationRef is the reference for NetworkConfiguration CR
type PortConfigurationRef struct {
	Name string `json:"name"`

	NameSpace string `json:"nameSpace"`

	// +kubebuilder:validation:Enum="SwitchPortConfiguration"
	Kind string `json:"kind"`
}

// Fetch the instance
func (ref *PortConfigurationRef) Fetch(ctx context.Context, client client.Client) (instance interface{}, err error) {
	switch ref.Kind {
	case "SwitchPortConfiguration":
		err = client.Get(
			ctx,
			types.NamespacedName{
				Name:      ref.Name,
				Namespace: ref.NameSpace,
			},
			instance.(*SwitchPortConfiguration),
		)
	default:
		err = fmt.Errorf("no instance for the ref")
	}

	return
}

// DeviceRef is the reference for Device CR
type DeviceRef struct {
	Name string `json:"name"`

	NameSpace string `json:"nameSpace"`

	// +kubebuilder:validation:Enum="Switch"
	Kind string `json:"kind"`
}

// Fetch the instance
func (ref *DeviceRef) Fetch(ctx context.Context, client client.Client) (instance interface{}, err error) {
	switch ref.Kind {
	case "Switch":
		err = client.Get(
			ctx,
			types.NamespacedName{
				Name:      ref.Name,
				Namespace: ref.NameSpace,
			},
			instance.(*Switch),
		)
	default:
		err = fmt.Errorf("no instance for the ref")
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
