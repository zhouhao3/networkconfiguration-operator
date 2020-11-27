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
	SchemeBuilder.Register(&Switch{}, &SwitchList{})
}

// +kubebuilder:object:root=true

// SwitchList contains a list of Switch
type SwitchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Switch `json:"items"`
}

// +kubebuilder:object:root=true

// Switch is the Schema for the switches API
type Switch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SwitchSpec   `json:"spec,omitempty"`
	Status SwitchStatus `json:"status,omitempty"`
}

// SwitchSpec defines the desired state of Switch
type SwitchSpec struct {
	OS    string           `json:"os"`
	IP    string           `json:"ip"`
	MAC   string           `json:"mac"`
	Ports []SwitchSpecPort `json:"ports,omitempty"`
}

// SwitchSpecPort ...
type SwitchSpecPort struct {
	PortID                  string                  `json:"portID"`
	LagWith                 string                  `json:"lagWith,omitempty"`
	NetworkConfigurationRef NetworkConfigurationRef `json:"networkConfigurationRef"`
}

// SwitchStatus defines the observed state of Switch
type SwitchStatus struct {
	Ports []SwitchStatusPort `json:"ports,omitempty"`
}

// SwitchStatusPort ...
type SwitchStatusPort struct {
	PortID  string    `json:"portID,omitempty"`
	LagWith string    `json:"lagWith,omitempty"`
	State   StateType `json:"state,omitempty"`
}
