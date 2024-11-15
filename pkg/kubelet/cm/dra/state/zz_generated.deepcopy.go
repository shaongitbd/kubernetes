//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package state

import (
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClaimInfoState) DeepCopyInto(out *ClaimInfoState) {
	*out = *in
	if in.PodUIDs != nil {
		in, out := &in.PodUIDs, &out.PodUIDs
		*out = make(sets.Set[string], len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DriverState != nil {
		in, out := &in.DriverState, &out.DriverState
		*out = make(map[string]DriverState, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClaimInfoState.
func (in *ClaimInfoState) DeepCopy() *ClaimInfoState {
	if in == nil {
		return nil
	}
	out := new(ClaimInfoState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	if in.RequestNames != nil {
		in, out := &in.RequestNames, &out.RequestNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CDIDeviceIDs != nil {
		in, out := &in.CDIDeviceIDs, &out.CDIDeviceIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriverState) DeepCopyInto(out *DriverState) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]Device, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriverState.
func (in *DriverState) DeepCopy() *DriverState {
	if in == nil {
		return nil
	}
	out := new(DriverState)
	in.DeepCopyInto(out)
	return out
}
