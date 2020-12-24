/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// StateType is the type of .status.state
type StateType string

// NicHint describes the requirements for the network card
type NicHint struct {
	// The name of the network card for this NicHint
	Name string `json:"name"`
	// True if smart network card is required, false otherwise.
	SmartNic bool `json:"smartNic"`
}

// PortRef is the reference for Port CR
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

// PortSpec defines the desired state of Port
type PortSpec struct {
	// Reference for PortConfiguration CR
	PortConfigurationRef PortConfigurationRef `json:"portConfigurationRef"`
	// Describes the port number on the device
	PortID string `json:"portID"`
	// Reference for Device CR
	DeviceRef DeviceRef `json:"deviceRef"`
	// ????
	SmartNic bool `json:"smartNic"`
}

// PortConfigurationRef is the reference for PortConfiguration CR
type PortConfigurationRef struct {
	Name string `json:"name"`

	NameSpace string `json:"nameSpace"`

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

// VLANID is a 12-bit 802.1Q VLAN identifier
type VLANID int32

// VLAN represents the name and ID of a VLAN
type VLAN struct {
	ID VLANID `json:"id"`

	Name string `json:"name,omitempty"`
}

// PortStatus defines the observed state of Port
type PortStatus struct {
	// The current configuration status of the port
	State StateType `json:"state,omitempty"`
	// The current portConfiguration of the port
	PortConfigurationRef PortConfigurationRef `json:"portConfigurationRef"`
	// VLAN information to which the port belongs
	VLANs []VLAN `json:"vlans"`
}

const (
	// PortNone ...
	PortNone StateType = ""

	// PortCreating ...
	PortCreating StateType = "Creating"

	// PortConfiguring ...
	PortConfiguring StateType = "Configuring"

	// PortConfigured ...
	PortConfigured StateType = "Configured"

	// PortDeleting ...
	PortDeleting StateType = "Deleting"

	// PortDeleted ...
	PortDeleted StateType = "Deleted"
)

// GetState gets the current state of the port
func (n *Port) GetState() StateType {
	return n.Status.State
}

// SetState sets the state of the port
func (n *Port) SetState(state StateType) {
	n.Status.State = state
}

// +kubebuilder:object:root=true

// Port is the Schema for the ports API
type Port struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PortSpec   `json:"spec,omitempty"`
	Status PortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PortList contains a list of Port
type PortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Port `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Port{}, &PortList{})
}
