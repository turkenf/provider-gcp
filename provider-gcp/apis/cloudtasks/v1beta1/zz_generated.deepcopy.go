//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

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

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineRoutingOverrideObservation) DeepCopyInto(out *AppEngineRoutingOverrideObservation) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineRoutingOverrideObservation.
func (in *AppEngineRoutingOverrideObservation) DeepCopy() *AppEngineRoutingOverrideObservation {
	if in == nil {
		return nil
	}
	out := new(AppEngineRoutingOverrideObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppEngineRoutingOverrideParameters) DeepCopyInto(out *AppEngineRoutingOverrideParameters) {
	*out = *in
	if in.Instance != nil {
		in, out := &in.Instance, &out.Instance
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppEngineRoutingOverrideParameters.
func (in *AppEngineRoutingOverrideParameters) DeepCopy() *AppEngineRoutingOverrideParameters {
	if in == nil {
		return nil
	}
	out := new(AppEngineRoutingOverrideParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueObservation) DeepCopyInto(out *QueueObservation) {
	*out = *in
	if in.AppEngineRoutingOverride != nil {
		in, out := &in.AppEngineRoutingOverride, &out.AppEngineRoutingOverride
		*out = make([]AppEngineRoutingOverrideObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = make([]RateLimitsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueObservation.
func (in *QueueObservation) DeepCopy() *QueueObservation {
	if in == nil {
		return nil
	}
	out := new(QueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueParameters) DeepCopyInto(out *QueueParameters) {
	*out = *in
	if in.AppEngineRoutingOverride != nil {
		in, out := &in.AppEngineRoutingOverride, &out.AppEngineRoutingOverride
		*out = make([]AppEngineRoutingOverrideParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProjectRef != nil {
		in, out := &in.ProjectRef, &out.ProjectRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ProjectSelector != nil {
		in, out := &in.ProjectSelector, &out.ProjectSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RateLimits != nil {
		in, out := &in.RateLimits, &out.RateLimits
		*out = make([]RateLimitsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RetryConfig != nil {
		in, out := &in.RetryConfig, &out.RetryConfig
		*out = make([]RetryConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StackdriverLoggingConfig != nil {
		in, out := &in.StackdriverLoggingConfig, &out.StackdriverLoggingConfig
		*out = make([]StackdriverLoggingConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueParameters.
func (in *QueueParameters) DeepCopy() *QueueParameters {
	if in == nil {
		return nil
	}
	out := new(QueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitsObservation) DeepCopyInto(out *RateLimitsObservation) {
	*out = *in
	if in.MaxBurstSize != nil {
		in, out := &in.MaxBurstSize, &out.MaxBurstSize
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitsObservation.
func (in *RateLimitsObservation) DeepCopy() *RateLimitsObservation {
	if in == nil {
		return nil
	}
	out := new(RateLimitsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitsParameters) DeepCopyInto(out *RateLimitsParameters) {
	*out = *in
	if in.MaxConcurrentDispatches != nil {
		in, out := &in.MaxConcurrentDispatches, &out.MaxConcurrentDispatches
		*out = new(float64)
		**out = **in
	}
	if in.MaxDispatchesPerSecond != nil {
		in, out := &in.MaxDispatchesPerSecond, &out.MaxDispatchesPerSecond
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitsParameters.
func (in *RateLimitsParameters) DeepCopy() *RateLimitsParameters {
	if in == nil {
		return nil
	}
	out := new(RateLimitsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfigObservation) DeepCopyInto(out *RetryConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfigObservation.
func (in *RetryConfigObservation) DeepCopy() *RetryConfigObservation {
	if in == nil {
		return nil
	}
	out := new(RetryConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfigParameters) DeepCopyInto(out *RetryConfigParameters) {
	*out = *in
	if in.MaxAttempts != nil {
		in, out := &in.MaxAttempts, &out.MaxAttempts
		*out = new(float64)
		**out = **in
	}
	if in.MaxBackoff != nil {
		in, out := &in.MaxBackoff, &out.MaxBackoff
		*out = new(string)
		**out = **in
	}
	if in.MaxDoublings != nil {
		in, out := &in.MaxDoublings, &out.MaxDoublings
		*out = new(float64)
		**out = **in
	}
	if in.MaxRetryDuration != nil {
		in, out := &in.MaxRetryDuration, &out.MaxRetryDuration
		*out = new(string)
		**out = **in
	}
	if in.MinBackoff != nil {
		in, out := &in.MinBackoff, &out.MinBackoff
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfigParameters.
func (in *RetryConfigParameters) DeepCopy() *RetryConfigParameters {
	if in == nil {
		return nil
	}
	out := new(RetryConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackdriverLoggingConfigObservation) DeepCopyInto(out *StackdriverLoggingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackdriverLoggingConfigObservation.
func (in *StackdriverLoggingConfigObservation) DeepCopy() *StackdriverLoggingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(StackdriverLoggingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackdriverLoggingConfigParameters) DeepCopyInto(out *StackdriverLoggingConfigParameters) {
	*out = *in
	if in.SamplingRatio != nil {
		in, out := &in.SamplingRatio, &out.SamplingRatio
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackdriverLoggingConfigParameters.
func (in *StackdriverLoggingConfigParameters) DeepCopy() *StackdriverLoggingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(StackdriverLoggingConfigParameters)
	in.DeepCopyInto(out)
	return out
}
