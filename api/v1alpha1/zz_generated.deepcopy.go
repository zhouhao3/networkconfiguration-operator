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
func (in *Port) DeepCopyInto(out *Port) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
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

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Port) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortConfigurationRef) DeepCopyInto(out *PortConfigurationRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortConfigurationRef.
func (in *PortConfigurationRef) DeepCopy() *PortConfigurationRef {
	if in == nil {
		return nil
	}
	out := new(PortConfigurationRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortList) DeepCopyInto(out *PortList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Port, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortList.
func (in *PortList) DeepCopy() *PortList {
	if in == nil {
		return nil
	}
	out := new(PortList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PortList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortRef) DeepCopyInto(out *PortRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortRef.
func (in *PortRef) DeepCopy() *PortRef {
	if in == nil {
		return nil
	}
	out := new(PortRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSpec) DeepCopyInto(out *PortSpec) {
	*out = *in
	out.PortConfigurationRef = in.PortConfigurationRef
	out.DeviceRef = in.DeviceRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSpec.
func (in *PortSpec) DeepCopy() *PortSpec {
	if in == nil {
		return nil
	}
	out := new(PortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortStatus) DeepCopyInto(out *PortStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortStatus.
func (in *PortStatus) DeepCopy() *PortStatus {
	if in == nil {
		return nil
	}
	out := new(PortStatus)
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
func (in *SwitchPortConfiguration) DeepCopyInto(out *SwitchPortConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfiguration.
func (in *SwitchPortConfiguration) DeepCopy() *SwitchPortConfiguration {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPortConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationList) DeepCopyInto(out *SwitchPortConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwitchPortConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationList.
func (in *SwitchPortConfigurationList) DeepCopy() *SwitchPortConfigurationList {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwitchPortConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationSpec) DeepCopyInto(out *SwitchPortConfigurationSpec) {
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationSpec.
func (in *SwitchPortConfigurationSpec) DeepCopy() *SwitchPortConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchPortConfigurationStatus) DeepCopyInto(out *SwitchPortConfigurationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwitchPortConfigurationStatus.
func (in *SwitchPortConfigurationStatus) DeepCopy() *SwitchPortConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(SwitchPortConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwitchSpec) DeepCopyInto(out *SwitchSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]SwitchSpecPort, len(*in))
		copy(*out, *in)
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
	out.PortConfigurationRef = in.PortConfigurationRef
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
		copy(*out, *in)
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
