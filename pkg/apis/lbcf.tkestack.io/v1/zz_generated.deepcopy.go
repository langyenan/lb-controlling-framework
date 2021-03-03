// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bind) DeepCopyInto(out *Bind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bind.
func (in *Bind) DeepCopy() *Bind {
	if in == nil {
		return nil
	}
	out := new(Bind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindList) DeepCopyInto(out *BindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindList.
func (in *BindList) DeepCopy() *BindList {
	if in == nil {
		return nil
	}
	out := new(BindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindSpec) DeepCopyInto(out *BindSpec) {
	*out = *in
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]TargetLoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Pods.DeepCopyInto(&out.Pods)
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeregisterPolicy != nil {
		in, out := &in.DeregisterPolicy, &out.DeregisterPolicy
		*out = new(DeregPolicy)
		**out = **in
	}
	if in.DeregisterWebhook != nil {
		in, out := &in.DeregisterWebhook, &out.DeregisterWebhook
		*out = new(DeregisterWebhookSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnsurePolicy != nil {
		in, out := &in.EnsurePolicy, &out.EnsurePolicy
		*out = new(EnsurePolicyConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindSpec.
func (in *BindSpec) DeepCopy() *BindSpec {
	if in == nil {
		return nil
	}
	out := new(BindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindStatus) DeepCopyInto(out *BindStatus) {
	*out = *in
	if in.LoadBalancerStatuses != nil {
		in, out := &in.LoadBalancerStatuses, &out.LoadBalancerStatuses
		*out = make([]TargetLoadBalancerStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindStatus.
func (in *BindStatus) DeepCopy() *BindStatus {
	if in == nil {
		return nil
	}
	out := new(BindStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeregisterWebhookSpec) DeepCopyInto(out *DeregisterWebhookSpec) {
	*out = *in
	if in.FailurePolicy != nil {
		in, out := &in.FailurePolicy, &out.FailurePolicy
		*out = new(deregFailurePolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeregisterWebhookSpec.
func (in *DeregisterWebhookSpec) DeepCopy() *DeregisterWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(DeregisterWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Duration) DeepCopyInto(out *Duration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Duration.
func (in *Duration) DeepCopy() *Duration {
	if in == nil {
		return nil
	}
	out := new(Duration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnsurePolicyConfig) DeepCopyInto(out *EnsurePolicyConfig) {
	*out = *in
	if in.ResyncPeriodInSeconds != nil {
		in, out := &in.ResyncPeriodInSeconds, &out.ResyncPeriodInSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnsurePolicyConfig.
func (in *EnsurePolicyConfig) DeepCopy() *EnsurePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(EnsurePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodBackend) DeepCopyInto(out *PodBackend) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortSelector, len(*in))
		copy(*out, *in)
	}
	if in.ByLabel != nil {
		in, out := &in.ByLabel, &out.ByLabel
		*out = new(SelectPodByLabel)
		(*in).DeepCopyInto(*out)
	}
	if in.ByName != nil {
		in, out := &in.ByName, &out.ByName
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodBackend.
func (in *PodBackend) DeepCopy() *PodBackend {
	if in == nil {
		return nil
	}
	out := new(PodBackend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectPodByLabel) DeepCopyInto(out *SelectPodByLabel) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Except != nil {
		in, out := &in.Except, &out.Except
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectPodByLabel.
func (in *SelectPodByLabel) DeepCopy() *SelectPodByLabel {
	if in == nil {
		return nil
	}
	out := new(SelectPodByLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLoadBalancer) DeepCopyInto(out *TargetLoadBalancer) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLoadBalancer.
func (in *TargetLoadBalancer) DeepCopy() *TargetLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(TargetLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLoadBalancerCondition) DeepCopyInto(out *TargetLoadBalancerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLoadBalancerCondition.
func (in *TargetLoadBalancerCondition) DeepCopy() *TargetLoadBalancerCondition {
	if in == nil {
		return nil
	}
	out := new(TargetLoadBalancerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetLoadBalancerStatus) DeepCopyInto(out *TargetLoadBalancerStatus) {
	*out = *in
	if in.LBInfo != nil {
		in, out := &in.LBInfo, &out.LBInfo
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LastSyncedAttributes != nil {
		in, out := &in.LastSyncedAttributes, &out.LastSyncedAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeletionTimestamp != nil {
		in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
		*out = new(string)
		**out = **in
	}
	in.RetryAfter.DeepCopyInto(&out.RetryAfter)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TargetLoadBalancerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetLoadBalancerStatus.
func (in *TargetLoadBalancerStatus) DeepCopy() *TargetLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(TargetLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}
