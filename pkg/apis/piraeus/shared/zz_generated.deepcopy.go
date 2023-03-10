//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Piraeus Operator
Copyright 2019 LINBIT USA, LLC.

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

package shared

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonPhysicalStorageOptions) DeepCopyInto(out *CommonPhysicalStorageOptions) {
	*out = *in
	if in.DevicePaths != nil {
		in, out := &in.DevicePaths, &out.DevicePaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonPhysicalStorageOptions.
func (in *CommonPhysicalStorageOptions) DeepCopy() *CommonPhysicalStorageOptions {
	if in == nil {
		return nil
	}
	out := new(CommonPhysicalStorageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonStoragePoolOptions) DeepCopyInto(out *CommonStoragePoolOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonStoragePoolOptions.
func (in *CommonStoragePoolOptions) DeepCopy() *CommonStoragePoolOptions {
	if in == nil {
		return nil
	}
	out := new(CommonStoragePoolOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorClientConfig) DeepCopyInto(out *LinstorClientConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorClientConfig.
func (in *LinstorClientConfig) DeepCopy() *LinstorClientConfig {
	if in == nil {
		return nil
	}
	out := new(LinstorClientConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SatelliteStatus) DeepCopyInto(out *SatelliteStatus) {
	*out = *in
	out.NodeStatus = in.NodeStatus
	if in.StoragePoolStatuses != nil {
		in, out := &in.StoragePoolStatuses, &out.StoragePoolStatuses
		*out = make([]*StoragePoolStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StoragePoolStatus)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SatelliteStatus.
func (in *SatelliteStatus) DeepCopy() *SatelliteStatus {
	if in == nil {
		return nil
	}
	out := new(SatelliteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolLVM) DeepCopyInto(out *StoragePoolLVM) {
	*out = *in
	out.CommonStoragePoolOptions = in.CommonStoragePoolOptions
	in.CommonPhysicalStorageOptions.DeepCopyInto(&out.CommonPhysicalStorageOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolLVM.
func (in *StoragePoolLVM) DeepCopy() *StoragePoolLVM {
	if in == nil {
		return nil
	}
	out := new(StoragePoolLVM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolLVMThin) DeepCopyInto(out *StoragePoolLVMThin) {
	*out = *in
	out.CommonStoragePoolOptions = in.CommonStoragePoolOptions
	in.CommonPhysicalStorageOptions.DeepCopyInto(&out.CommonPhysicalStorageOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolLVMThin.
func (in *StoragePoolLVMThin) DeepCopy() *StoragePoolLVMThin {
	if in == nil {
		return nil
	}
	out := new(StoragePoolLVMThin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolStatus) DeepCopyInto(out *StoragePoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolStatus.
func (in *StoragePoolStatus) DeepCopy() *StoragePoolStatus {
	if in == nil {
		return nil
	}
	out := new(StoragePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolZFS) DeepCopyInto(out *StoragePoolZFS) {
	*out = *in
	out.CommonStoragePoolOptions = in.CommonStoragePoolOptions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolZFS.
func (in *StoragePoolZFS) DeepCopy() *StoragePoolZFS {
	if in == nil {
		return nil
	}
	out := new(StoragePoolZFS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePools) DeepCopyInto(out *StoragePools) {
	*out = *in
	if in.LVMPools != nil {
		in, out := &in.LVMPools, &out.LVMPools
		*out = make([]*StoragePoolLVM, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StoragePoolLVM)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.LVMThinPools != nil {
		in, out := &in.LVMThinPools, &out.LVMThinPools
		*out = make([]*StoragePoolLVMThin, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StoragePoolLVMThin)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ZFSPools != nil {
		in, out := &in.ZFSPools, &out.ZFSPools
		*out = make([]*StoragePoolZFS, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(StoragePoolZFS)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePools.
func (in *StoragePools) DeepCopy() *StoragePools {
	if in == nil {
		return nil
	}
	out := new(StoragePools)
	in.DeepCopyInto(out)
	return out
}
