// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeMetadata) DeepCopyInto(out *NodeMetadata) {
	*out = *in
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
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(v1.Node)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeMetadata.
func (in *NodeMetadata) DeepCopy() *NodeMetadata {
	if in == nil {
		return nil
	}
	out := new(NodeMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodemgr) DeepCopyInto(out *Nodemgr) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodemgr.
func (in *Nodemgr) DeepCopy() *Nodemgr {
	if in == nil {
		return nil
	}
	out := new(Nodemgr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Nodemgr) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrAnnotations) DeepCopyInto(out *NodemgrAnnotations) {
	*out = *in
	if in.LB != nil {
		in, out := &in.LB, &out.LB
		*out = make([]NodemgrLB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrAnnotations.
func (in *NodemgrAnnotations) DeepCopy() *NodemgrAnnotations {
	if in == nil {
		return nil
	}
	out := new(NodemgrAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrLB) DeepCopyInto(out *NodemgrLB) {
	*out = *in
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = make([]NodemgrLBInstances, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrLB.
func (in *NodemgrLB) DeepCopy() *NodemgrLB {
	if in == nil {
		return nil
	}
	out := new(NodemgrLB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrLBInstances) DeepCopyInto(out *NodemgrLBInstances) {
	*out = *in
	if in.LocalPort != nil {
		in, out := &in.LocalPort, &out.LocalPort
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrLBInstances.
func (in *NodemgrLBInstances) DeepCopy() *NodemgrLBInstances {
	if in == nil {
		return nil
	}
	out := new(NodemgrLBInstances)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrList) DeepCopyInto(out *NodemgrList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Nodemgr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrList.
func (in *NodemgrList) DeepCopy() *NodemgrList {
	if in == nil {
		return nil
	}
	out := new(NodemgrList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodemgrList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrMetadata) DeepCopyInto(out *NodemgrMetadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Annotations.DeepCopyInto(&out.Annotations)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]NodemgrTaint, len(*in))
		copy(*out, *in)
	}
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(Nodemgr)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrMetadata.
func (in *NodemgrMetadata) DeepCopy() *NodemgrMetadata {
	if in == nil {
		return nil
	}
	out := new(NodemgrMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrSpec) DeepCopyInto(out *NodemgrSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]NodemgrTaint, len(*in))
		copy(*out, *in)
	}
	in.Annotations.DeepCopyInto(&out.Annotations)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrSpec.
func (in *NodemgrSpec) DeepCopy() *NodemgrSpec {
	if in == nil {
		return nil
	}
	out := new(NodemgrSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrStatus) DeepCopyInto(out *NodemgrStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrStatus.
func (in *NodemgrStatus) DeepCopy() *NodemgrStatus {
	if in == nil {
		return nil
	}
	out := new(NodemgrStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodemgrTaint) DeepCopyInto(out *NodemgrTaint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodemgrTaint.
func (in *NodemgrTaint) DeepCopy() *NodemgrTaint {
	if in == nil {
		return nil
	}
	out := new(NodemgrTaint)
	in.DeepCopyInto(out)
	return out
}
