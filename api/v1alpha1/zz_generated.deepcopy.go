//go:build !ignore_autogenerated

/*
Copyright 2024 MatanMagen.

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
func (in *MatanApp) DeepCopyInto(out *MatanApp) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatanApp.
func (in *MatanApp) DeepCopy() *MatanApp {
	if in == nil {
		return nil
	}
	out := new(MatanApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MatanApp) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatanAppList) DeepCopyInto(out *MatanAppList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MatanApp, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatanAppList.
func (in *MatanAppList) DeepCopy() *MatanAppList {
	if in == nil {
		return nil
	}
	out := new(MatanAppList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MatanAppList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatanAppSpec) DeepCopyInto(out *MatanAppSpec) {
	*out = *in
	in.Secret.DeepCopyInto(&out.Secret)
	in.ConfigMap.DeepCopyInto(&out.ConfigMap)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatanAppSpec.
func (in *MatanAppSpec) DeepCopy() *MatanAppSpec {
	if in == nil {
		return nil
	}
	out := new(MatanAppSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MatanAppStatus) DeepCopyInto(out *MatanAppStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MatanAppStatus.
func (in *MatanAppStatus) DeepCopy() *MatanAppStatus {
	if in == nil {
		return nil
	}
	out := new(MatanAppStatus)
	in.DeepCopyInto(out)
	return out
}