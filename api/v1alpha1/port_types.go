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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	SchemeBuilder.Register(&Port{}, &PortList{})
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

// PortSpec defines the desired state of Port
type PortSpec struct {
	PortConfigurationRef PortConfigurationRef `json:"portConfigurationRef"`
	PortID               string               `json:"portID"`
	LagWith              string               `json:"lagWith,omitempty"`
	DeviceRef            DeviceRef            `json:"deviceRef"`
}

// PortStatus defines the observed state of Port
type PortStatus struct {
	State StateType `json:"state,omitempty"`
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

// GetState ...
func (n *Port) GetState() StateType {
	return n.Status.State
}

// SetState ...
func (n *Port) SetState(state StateType) {
	n.Status.State = state
}
