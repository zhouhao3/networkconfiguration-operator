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
	SchemeBuilder.Register(&NetworkBinding{}, &NetworkBindingList{})
}

// +kubebuilder:object:root=true

// NetworkBindingList contains a list of NetworkBinding
type NetworkBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkBinding `json:"items"`
}

// +kubebuilder:object:root=true

// NetworkBinding is the Schema for the networkbindings API
type NetworkBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkBindingSpec   `json:"spec,omitempty"`
	Status NetworkBindingStatus `json:"status,omitempty"`
}

// NetworkBindingSpec defines the desired state of NetworkBinding
type NetworkBindingSpec struct {
	NetworkConfigurationRef NetworkConfigurationRef `json:"networkConfigurationRef"`

	Port NetworkBindingSpecPort `json:"port"`
}

// NetworkBindingSpecPort ...
type NetworkBindingSpecPort struct {
	PortID    string    `json:"portID"`
	LagWith   string    `json:"lagWith,omitempty"`
	DeviceRef DeviceRef `json:"deviceRef"`
}

// NetworkBindingStatus defines the observed state of NetworkBinding
type NetworkBindingStatus struct {
	State StateType `json:"state,omitempty"`
}

const (
	// NetworkBindingNone ...
	NetworkBindingNone StateType = ""

	// NetworkBindingCreating ...
	NetworkBindingCreating StateType = "Creating"

	// NetworkBindingConfiguring ...
	NetworkBindingConfiguring StateType = "Configuring"

	// NetworkBindingConfigured ...
	NetworkBindingConfigured StateType = "Configured"

	// NetworkBindingDeleting ...
	NetworkBindingDeleting StateType = "Deleting"

	// NetworkBindingDeleted ...
	NetworkBindingDeleted StateType = "Deleted"
)

// GetState ...
func (n *NetworkBinding) GetState() StateType {
	return n.Status.State
}

// SetState ...
func (n *NetworkBinding) SetState(state StateType) {
	n.Status.State = state
}
