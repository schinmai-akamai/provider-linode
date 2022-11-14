//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigList) DeepCopyInto(out *ConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigList.
func (in *ConfigList) DeepCopy() *ConfigList {
	if in == nil {
		return nil
	}
	out := new(ConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
	if in.Booted != nil {
		in, out := &in.Booted, &out.Booted
		*out = new(bool)
		**out = **in
	}
	if in.Comments != nil {
		in, out := &in.Comments, &out.Comments
		*out = new(string)
		**out = **in
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]DevicesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Helpers != nil {
		in, out := &in.Helpers, &out.Helpers
		*out = make([]HelpersParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = make([]InterfaceParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kernel != nil {
		in, out := &in.Kernel, &out.Kernel
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.MemoryLimit != nil {
		in, out := &in.MemoryLimit, &out.MemoryLimit
		*out = new(float64)
		**out = **in
	}
	if in.RootDevice != nil {
		in, out := &in.RootDevice, &out.RootDevice
		*out = new(string)
		**out = **in
	}
	if in.RunLevel != nil {
		in, out := &in.RunLevel, &out.RunLevel
		*out = new(string)
		**out = **in
	}
	if in.VirtMode != nil {
		in, out := &in.VirtMode, &out.VirtMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigStatus) DeepCopyInto(out *ConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigStatus.
func (in *ConfigStatus) DeepCopy() *ConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicesObservation) DeepCopyInto(out *DevicesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicesObservation.
func (in *DevicesObservation) DeepCopy() *DevicesObservation {
	if in == nil {
		return nil
	}
	out := new(DevicesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicesParameters) DeepCopyInto(out *DevicesParameters) {
	*out = *in
	if in.Sda != nil {
		in, out := &in.Sda, &out.Sda
		*out = make([]SdaParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdb != nil {
		in, out := &in.Sdb, &out.Sdb
		*out = make([]SdbParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdc != nil {
		in, out := &in.Sdc, &out.Sdc
		*out = make([]SdcParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdd != nil {
		in, out := &in.Sdd, &out.Sdd
		*out = make([]SddParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sde != nil {
		in, out := &in.Sde, &out.Sde
		*out = make([]SdeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdf != nil {
		in, out := &in.Sdf, &out.Sdf
		*out = make([]SdfParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdg != nil {
		in, out := &in.Sdg, &out.Sdg
		*out = make([]SdgParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sdh != nil {
		in, out := &in.Sdh, &out.Sdh
		*out = make([]SdhParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicesParameters.
func (in *DevicesParameters) DeepCopy() *DevicesParameters {
	if in == nil {
		return nil
	}
	out := new(DevicesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelpersObservation) DeepCopyInto(out *HelpersObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelpersObservation.
func (in *HelpersObservation) DeepCopy() *HelpersObservation {
	if in == nil {
		return nil
	}
	out := new(HelpersObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelpersParameters) DeepCopyInto(out *HelpersParameters) {
	*out = *in
	if in.DevtmpfsAutomount != nil {
		in, out := &in.DevtmpfsAutomount, &out.DevtmpfsAutomount
		*out = new(bool)
		**out = **in
	}
	if in.Distro != nil {
		in, out := &in.Distro, &out.Distro
		*out = new(bool)
		**out = **in
	}
	if in.ModulesDep != nil {
		in, out := &in.ModulesDep, &out.ModulesDep
		*out = new(bool)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(bool)
		**out = **in
	}
	if in.UpdatedbDisabled != nil {
		in, out := &in.UpdatedbDisabled, &out.UpdatedbDisabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelpersParameters.
func (in *HelpersParameters) DeepCopy() *HelpersParameters {
	if in == nil {
		return nil
	}
	out := new(HelpersParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceObservation) DeepCopyInto(out *InterfaceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceObservation.
func (in *InterfaceObservation) DeepCopy() *InterfaceObservation {
	if in == nil {
		return nil
	}
	out := new(InterfaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceParameters) DeepCopyInto(out *InterfaceParameters) {
	*out = *in
	if in.IpamAddress != nil {
		in, out := &in.IpamAddress, &out.IpamAddress
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Purpose != nil {
		in, out := &in.Purpose, &out.Purpose
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceParameters.
func (in *InterfaceParameters) DeepCopy() *InterfaceParameters {
	if in == nil {
		return nil
	}
	out := new(InterfaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdaObservation) DeepCopyInto(out *SdaObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdaObservation.
func (in *SdaObservation) DeepCopy() *SdaObservation {
	if in == nil {
		return nil
	}
	out := new(SdaObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdaParameters) DeepCopyInto(out *SdaParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdaParameters.
func (in *SdaParameters) DeepCopy() *SdaParameters {
	if in == nil {
		return nil
	}
	out := new(SdaParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdbObservation) DeepCopyInto(out *SdbObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdbObservation.
func (in *SdbObservation) DeepCopy() *SdbObservation {
	if in == nil {
		return nil
	}
	out := new(SdbObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdbParameters) DeepCopyInto(out *SdbParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdbParameters.
func (in *SdbParameters) DeepCopy() *SdbParameters {
	if in == nil {
		return nil
	}
	out := new(SdbParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdcObservation) DeepCopyInto(out *SdcObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdcObservation.
func (in *SdcObservation) DeepCopy() *SdcObservation {
	if in == nil {
		return nil
	}
	out := new(SdcObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdcParameters) DeepCopyInto(out *SdcParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdcParameters.
func (in *SdcParameters) DeepCopy() *SdcParameters {
	if in == nil {
		return nil
	}
	out := new(SdcParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddObservation) DeepCopyInto(out *SddObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddObservation.
func (in *SddObservation) DeepCopy() *SddObservation {
	if in == nil {
		return nil
	}
	out := new(SddObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SddParameters) DeepCopyInto(out *SddParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SddParameters.
func (in *SddParameters) DeepCopy() *SddParameters {
	if in == nil {
		return nil
	}
	out := new(SddParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdeObservation) DeepCopyInto(out *SdeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdeObservation.
func (in *SdeObservation) DeepCopy() *SdeObservation {
	if in == nil {
		return nil
	}
	out := new(SdeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdeParameters) DeepCopyInto(out *SdeParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdeParameters.
func (in *SdeParameters) DeepCopy() *SdeParameters {
	if in == nil {
		return nil
	}
	out := new(SdeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdfObservation) DeepCopyInto(out *SdfObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdfObservation.
func (in *SdfObservation) DeepCopy() *SdfObservation {
	if in == nil {
		return nil
	}
	out := new(SdfObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdfParameters) DeepCopyInto(out *SdfParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdfParameters.
func (in *SdfParameters) DeepCopy() *SdfParameters {
	if in == nil {
		return nil
	}
	out := new(SdfParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdgObservation) DeepCopyInto(out *SdgObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdgObservation.
func (in *SdgObservation) DeepCopy() *SdgObservation {
	if in == nil {
		return nil
	}
	out := new(SdgObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdgParameters) DeepCopyInto(out *SdgParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdgParameters.
func (in *SdgParameters) DeepCopy() *SdgParameters {
	if in == nil {
		return nil
	}
	out := new(SdgParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdhObservation) DeepCopyInto(out *SdhObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdhObservation.
func (in *SdhObservation) DeepCopy() *SdhObservation {
	if in == nil {
		return nil
	}
	out := new(SdhObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SdhParameters) DeepCopyInto(out *SdhParameters) {
	*out = *in
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.VolumeID != nil {
		in, out := &in.VolumeID, &out.VolumeID
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SdhParameters.
func (in *SdhParameters) DeepCopy() *SdhParameters {
	if in == nil {
		return nil
	}
	out := new(SdhParameters)
	in.DeepCopyInto(out)
	return out
}