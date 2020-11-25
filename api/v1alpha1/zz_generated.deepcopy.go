// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACL) DeepCopyInto(out *ACL) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACL.
func (in *ACL) DeepCopy() *ACL {
	if in == nil {
		return nil
	}
	out := new(ACL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceRef) DeepCopyInto(out *DeviceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceRef.
func (in *DeviceRef) DeepCopy() *DeviceRef {
	if in == nil {
		return nil
	}
	out := new(DeviceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NICHint) DeepCopyInto(out *NICHint) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NICHint.
func (in *NICHint) DeepCopy() *NICHint {
	if in == nil {
		return nil
	}
	out := new(NICHint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkBinding) DeepCopyInto(out *NetworkBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkBinding.
func (in *NetworkBinding) DeepCopy() *NetworkBinding {
	if in == nil {
		return nil
	}
	out := new(NetworkBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkBindingList) DeepCopyInto(out *NetworkBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkBindingList.
func (in *NetworkBindingList) DeepCopy() *NetworkBindingList {
	if in == nil {
		return nil
	}
	out := new(NetworkBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkBindingRef) DeepCopyInto(out *NetworkBindingRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkBindingRef.
func (in *NetworkBindingRef) DeepCopy() *NetworkBindingRef {
	if in == nil {
		return nil
	}
	out := new(NetworkBindingRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkBindingSpec) DeepCopyInto(out *NetworkBindingSpec) {
	*out = *in
	out.NetworkConfigurationRef = in.NetworkConfigurationRef
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkBindingSpec.
func (in *NetworkBindingSpec) DeepCopy() *NetworkBindingSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkBindingStatus) DeepCopyInto(out *NetworkBindingStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkBindingStatus.
func (in *NetworkBindingStatus) DeepCopy() *NetworkBindingStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfiguration) DeepCopyInto(out *NetworkConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfiguration.
func (in *NetworkConfiguration) DeepCopy() *NetworkConfiguration {
	if in == nil {
		return nil
	}
	out := new(NetworkConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationList) DeepCopyInto(out *NetworkConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationList.
func (in *NetworkConfigurationList) DeepCopy() *NetworkConfigurationList {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationRef) DeepCopyInto(out *NetworkConfigurationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationRef.
func (in *NetworkConfigurationRef) DeepCopy() *NetworkConfigurationRef {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationSpec) DeepCopyInto(out *NetworkConfigurationSpec) {
	*out = *in
	if in.ACLs != nil {
		in, out := &in.ACLs, &out.ACLs
		*out = make([]ACL, len(*in))
		copy(*out, *in)
	}
	if in.Vlans != nil {
		in, out := &in.Vlans, &out.Vlans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NICHint != nil {
		in, out := &in.NICHint, &out.NICHint
		*out = new(NICHint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationSpec.
func (in *NetworkConfigurationSpec) DeepCopy() *NetworkConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigurationStatus) DeepCopyInto(out *NetworkConfigurationStatus) {
	*out = *in
	if in.NetworkBindingRefs != nil {
		in, out := &in.NetworkBindingRefs, &out.NetworkBindingRefs
		*out = make([]NetworkBindingRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigurationStatus.
func (in *NetworkConfigurationStatus) DeepCopy() *NetworkConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	out.DeviceRef = in.DeviceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Port.
func (in *Port) DeepCopy() *Port {
	if in == nil {
		return nil
	}
	out := new(Port)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Switch) DeepCopyInto(out *Switch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Switch.
func (in *Switch) DeepCopy() *Switch {
	if in == nil {
		return nil
	}
	out := new(Switch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Switch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchList) DeepCopyInto(out *SwitchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Switch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchList.
func (in *SwitchList) DeepCopy() *SwitchList {
	if in == nil {
		return nil
	}
	out := new(SwitchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSpec) DeepCopyInto(out *SwitchSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]SwitchSpecPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSpec.
func (in *SwitchSpec) DeepCopy() *SwitchSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSpecPort) DeepCopyInto(out *SwitchSpecPort) {
	*out = *in
	if in.NetworkConfigurationRef != nil {
		in, out := &in.NetworkConfigurationRef, &out.NetworkConfigurationRef
		*out = new(NetworkConfigurationRef)
		**out = **in
	}
	out.LagWith = in.LagWith
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchSpecPort.
func (in *SwitchSpecPort) DeepCopy() *SwitchSpecPort {
	if in == nil {
		return nil
	}
	out := new(SwitchSpecPort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchStatus) DeepCopyInto(out *SwitchStatus) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]SwitchStatusPort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchStatus.
func (in *SwitchStatus) DeepCopy() *SwitchStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchStatusPort) DeepCopyInto(out *SwitchStatusPort) {
	*out = *in
	if in.NetworkConfiguration != nil {
		in, out := &in.NetworkConfiguration, &out.NetworkConfiguration
		*out = new(NetworkConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchStatusPort.
func (in *SwitchStatusPort) DeepCopy() *SwitchStatusPort {
	if in == nil {
		return nil
	}
	out := new(SwitchStatusPort)
	in.DeepCopyInto(out)
	return out
}
