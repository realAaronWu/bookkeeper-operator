//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperCluster) DeepCopyInto(out *BookkeeperCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperCluster.
func (in *BookkeeperCluster) DeepCopy() *BookkeeperCluster {
	if in == nil {
		return nil
	}
	out := new(BookkeeperCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BookkeeperCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperClusterList) DeepCopyInto(out *BookkeeperClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BookkeeperCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperClusterList.
func (in *BookkeeperClusterList) DeepCopy() *BookkeeperClusterList {
	if in == nil {
		return nil
	}
	out := new(BookkeeperClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BookkeeperClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperClusterSpec) DeepCopyInto(out *BookkeeperClusterSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(BookkeeperImageSpec)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(BookkeeperStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoRecovery != nil {
		in, out := &in.AutoRecovery, &out.AutoRecovery
		*out = new(bool)
		**out = **in
	}
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = new(Probes)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.JVMOptions != nil {
		in, out := &in.JVMOptions, &out.JVMOptions
		*out = new(JVMOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.BlockOwnerDeletion != nil {
		in, out := &in.BlockOwnerDeletion, &out.BlockOwnerDeletion
		*out = new(bool)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RunAsPrivilegedUser != nil {
		in, out := &in.RunAsPrivilegedUser, &out.RunAsPrivilegedUser
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperClusterSpec.
func (in *BookkeeperClusterSpec) DeepCopy() *BookkeeperClusterSpec {
	if in == nil {
		return nil
	}
	out := new(BookkeeperClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperClusterStatus) DeepCopyInto(out *BookkeeperClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterCondition, len(*in))
		copy(*out, *in)
	}
	if in.VersionHistory != nil {
		in, out := &in.VersionHistory, &out.VersionHistory
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Members.DeepCopyInto(&out.Members)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperClusterStatus.
func (in *BookkeeperClusterStatus) DeepCopy() *BookkeeperClusterStatus {
	if in == nil {
		return nil
	}
	out := new(BookkeeperClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperImageSpec) DeepCopyInto(out *BookkeeperImageSpec) {
	*out = *in
	out.ImageSpec = in.ImageSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperImageSpec.
func (in *BookkeeperImageSpec) DeepCopy() *BookkeeperImageSpec {
	if in == nil {
		return nil
	}
	out := new(BookkeeperImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperStorageSpec) DeepCopyInto(out *BookkeeperStorageSpec) {
	*out = *in
	if in.LedgerVolumeClaimTemplate != nil {
		in, out := &in.LedgerVolumeClaimTemplate, &out.LedgerVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JournalVolumeClaimTemplate != nil {
		in, out := &in.JournalVolumeClaimTemplate, &out.JournalVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexVolumeClaimTemplate != nil {
		in, out := &in.IndexVolumeClaimTemplate, &out.IndexVolumeClaimTemplate
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperStorageSpec.
func (in *BookkeeperStorageSpec) DeepCopy() *BookkeeperStorageSpec {
	if in == nil {
		return nil
	}
	out := new(BookkeeperStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterCondition) DeepCopyInto(out *ClusterCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterCondition.
func (in *ClusterCondition) DeepCopy() *ClusterCondition {
	if in == nil {
		return nil
	}
	out := new(ClusterCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVMOptions) DeepCopyInto(out *JVMOptions) {
	*out = *in
	if in.MemoryOpts != nil {
		in, out := &in.MemoryOpts, &out.MemoryOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GcOpts != nil {
		in, out := &in.GcOpts, &out.GcOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.GcLoggingOpts != nil {
		in, out := &in.GcLoggingOpts, &out.GcLoggingOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExtraOpts != nil {
		in, out := &in.ExtraOpts, &out.ExtraOpts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVMOptions.
func (in *JVMOptions) DeepCopy() *JVMOptions {
	if in == nil {
		return nil
	}
	out := new(JVMOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MembersStatus) DeepCopyInto(out *MembersStatus) {
	*out = *in
	if in.Ready != nil {
		in, out := &in.Ready, &out.Ready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Unready != nil {
		in, out := &in.Unready, &out.Unready
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MembersStatus.
func (in *MembersStatus) DeepCopy() *MembersStatus {
	if in == nil {
		return nil
	}
	out := new(MembersStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probe) DeepCopyInto(out *Probe) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probe.
func (in *Probe) DeepCopy() *Probe {
	if in == nil {
		return nil
	}
	out := new(Probe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Probes) DeepCopyInto(out *Probes) {
	*out = *in
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(Probe)
		**out = **in
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(Probe)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Probes.
func (in *Probes) DeepCopy() *Probes {
	if in == nil {
		return nil
	}
	out := new(Probes)
	in.DeepCopyInto(out)
	return out
}
