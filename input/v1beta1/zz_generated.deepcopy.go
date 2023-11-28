//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Delims) DeepCopyInto(out *Delims) {
	*out = *in
	if in.Left != nil {
		in, out := &in.Left, &out.Left
		*out = new(string)
		**out = **in
	}
	if in.Right != nil {
		in, out := &in.Right, &out.Right
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Delims.
func (in *Delims) DeepCopy() *Delims {
	if in == nil {
		return nil
	}
	out := new(Delims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoTemplate) DeepCopyInto(out *GoTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Delims != nil {
		in, out := &in.Delims, &out.Delims
		*out = new(Delims)
		(*in).DeepCopyInto(*out)
	}
	if in.Inline != nil {
		in, out := &in.Inline, &out.Inline
		*out = new(TemplateSourceInline)
		**out = **in
	}
	if in.FileSystem != nil {
		in, out := &in.FileSystem, &out.FileSystem
		*out = new(TemplateSourceFileSystem)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoTemplate.
func (in *GoTemplate) DeepCopy() *GoTemplate {
	if in == nil {
		return nil
	}
	out := new(GoTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GoTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSourceFileSystem) DeepCopyInto(out *TemplateSourceFileSystem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSourceFileSystem.
func (in *TemplateSourceFileSystem) DeepCopy() *TemplateSourceFileSystem {
	if in == nil {
		return nil
	}
	out := new(TemplateSourceFileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSourceInline) DeepCopyInto(out *TemplateSourceInline) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSourceInline.
func (in *TemplateSourceInline) DeepCopy() *TemplateSourceInline {
	if in == nil {
		return nil
	}
	out := new(TemplateSourceInline)
	in.DeepCopyInto(out)
	return out
}
