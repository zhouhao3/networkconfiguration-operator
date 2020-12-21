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
	corev1 "k8s.io/api/core/v1"
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
	OS     string                  `json:"os"`
	IP     string                  `json:"ip"`
	MAC    string                  `json:"mac"`
	Secret *corev1.SecretReference `json:"secret"`
	Ports  []SwitchSpecPort        `json:"ports,omitempty"`
}

// SwitchSpecPort ...
type SwitchSpecPort struct {
	PortID string `json:"portID,omitempty"`

	Disabled bool `json:"disabled,omitempty"`

	TrunkDisabled bool `json:"trunkDisable,omitempty"`

	// +kubebuilder:validation:Pattern=`([0-9]{1,})|([0-9]{1,}-[0-9]{1,})(,([0-9]{1,})|([0-9]{1,}-[0-9]{1,}))*`
	VlanRange string `json:"vlanRange,omitempty"`
}

// SwitchStatus defines the observed state of Switch
type SwitchStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}
