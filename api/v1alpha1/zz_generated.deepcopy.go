//go:build !ignore_autogenerated

/*
Copyright 2024.

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
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Additional) DeepCopyInto(out *Additional) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]corev1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ServicePorts != nil {
		in, out := &in.ServicePorts, &out.ServicePorts
		*out = make([]corev1.ServicePort, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Additional.
func (in *Additional) DeepCopy() *Additional {
	if in == nil {
		return nil
	}
	out := new(Additional)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockConfig) DeepCopyInto(out *BlockConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockConfig.
func (in *BlockConfig) DeepCopy() *BlockConfig {
	if in == nil {
		return nil
	}
	out := new(BlockConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonThanosFields) DeepCopyInto(out *CommonThanosFields) {
	*out = *in
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.ResourceRequirements != nil {
		in, out := &in.ResourceRequirements, &out.ResourceRequirements
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.LogFormat != nil {
		in, out := &in.LogFormat, &out.LogFormat
		*out = new(string)
		**out = **in
	}
	if in.EnableSelfMonitor != nil {
		in, out := &in.EnableSelfMonitor, &out.EnableSelfMonitor
		*out = new(bool)
		**out = **in
	}
	if in.ServiceMonitorConfig != nil {
		in, out := &in.ServiceMonitorConfig, &out.ServiceMonitorConfig
		*out = new(ServiceMonitorConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonThanosFields.
func (in *CommonThanosFields) DeepCopy() *CommonThanosFields {
	if in == nil {
		return nil
	}
	out := new(CommonThanosFields)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DownsamplingConfig) DeepCopyInto(out *DownsamplingConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DownsamplingConfig.
func (in *DownsamplingConfig) DeepCopy() *DownsamplingConfig {
	if in == nil {
		return nil
	}
	out := new(DownsamplingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ExternalLabelShardingConfig) DeepCopyInto(out *ExternalLabelShardingConfig) {
	{
		in := &in
		*out = make(ExternalLabelShardingConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLabelShardingConfig.
func (in ExternalLabelShardingConfig) DeepCopy() ExternalLabelShardingConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalLabelShardingConfig)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ExternalLabels) DeepCopyInto(out *ExternalLabels) {
	{
		in := &in
		*out = make(ExternalLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalLabels.
func (in ExternalLabels) DeepCopy() ExternalLabels {
	if in == nil {
		return nil
	}
	out := new(ExternalLabels)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupConfig) DeepCopyInto(out *GroupConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupConfig.
func (in *GroupConfig) DeepCopy() *GroupConfig {
	if in == nil {
		return nil
	}
	out := new(GroupConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngesterHashringSpec) DeepCopyInto(out *IngesterHashringSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalLabels != nil {
		in, out := &in.ExternalLabels, &out.ExternalLabels
		*out = make(ExternalLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TSDBConfig = in.TSDBConfig
	if in.ObjectStorageConfig != nil {
		in, out := &in.ObjectStorageConfig, &out.ObjectStorageConfig
		*out = new(ObjectStorageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngesterHashringSpec.
func (in *IngesterHashringSpec) DeepCopy() *IngesterHashringSpec {
	if in == nil {
		return nil
	}
	out := new(IngesterHashringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngesterSpec) DeepCopyInto(out *IngesterSpec) {
	*out = *in
	in.DefaultObjectStorageConfig.DeepCopyInto(&out.DefaultObjectStorageConfig)
	if in.Hashrings != nil {
		in, out := &in.Hashrings, &out.Hashrings
		*out = make([]IngesterHashringSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngesterSpec.
func (in *IngesterSpec) DeepCopy() *IngesterSpec {
	if in == nil {
		return nil
	}
	out := new(IngesterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageConfig) DeepCopyInto(out *ObjectStorageConfig) {
	*out = *in
	out.LocalObjectReference = in.LocalObjectReference
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageConfig.
func (in *ObjectStorageConfig) DeepCopy() *ObjectStorageConfig {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryFrontendSpec) DeepCopyInto(out *QueryFrontendSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.QueryLabelSelector != nil {
		in, out := &in.QueryLabelSelector, &out.QueryLabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.LogQueriesLongerThan != nil {
		in, out := &in.LogQueriesLongerThan, &out.LogQueriesLongerThan
		*out = new(Duration)
		**out = **in
	}
	if in.QueryRangeResponseCacheConfig != nil {
		in, out := &in.QueryRangeResponseCacheConfig, &out.QueryRangeResponseCacheConfig
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryRangeSplitInterval != nil {
		in, out := &in.QueryRangeSplitInterval, &out.QueryRangeSplitInterval
		*out = new(Duration)
		**out = **in
	}
	if in.LabelsSplitInterval != nil {
		in, out := &in.LabelsSplitInterval, &out.LabelsSplitInterval
		*out = new(Duration)
		**out = **in
	}
	if in.LabelsDefaultTimeRange != nil {
		in, out := &in.LabelsDefaultTimeRange, &out.LabelsDefaultTimeRange
		*out = new(Duration)
		**out = **in
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryFrontendSpec.
func (in *QueryFrontendSpec) DeepCopy() *QueryFrontendSpec {
	if in == nil {
		return nil
	}
	out := new(QueryFrontendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetentionResolutionConfig) DeepCopyInto(out *RetentionResolutionConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetentionResolutionConfig.
func (in *RetentionResolutionConfig) DeepCopy() *RetentionResolutionConfig {
	if in == nil {
		return nil
	}
	out := new(RetentionResolutionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouterSpec) DeepCopyInto(out *RouterSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExternalLabels != nil {
		in, out := &in.ExternalLabels, &out.ExternalLabels
		*out = make(ExternalLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouterSpec.
func (in *RouterSpec) DeepCopy() *RouterSpec {
	if in == nil {
		return nil
	}
	out := new(RouterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorConfig) DeepCopyInto(out *ServiceMonitorConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorConfig.
func (in *ServiceMonitorConfig) DeepCopy() *ServiceMonitorConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingConfig) DeepCopyInto(out *ShardingConfig) {
	*out = *in
	if in.ExternalLabelSharding != nil {
		in, out := &in.ExternalLabelSharding, &out.ExternalLabelSharding
		*out = make(ExternalLabelShardingConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingConfig.
func (in *ShardingConfig) DeepCopy() *ShardingConfig {
	if in == nil {
		return nil
	}
	out := new(ShardingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShardingStrategy) DeepCopyInto(out *ShardingStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShardingStrategy.
func (in *ShardingStrategy) DeepCopy() *ShardingStrategy {
	if in == nil {
		return nil
	}
	out := new(ShardingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TSDBConfig) DeepCopyInto(out *TSDBConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TSDBConfig.
func (in *TSDBConfig) DeepCopy() *TSDBConfig {
	if in == nil {
		return nil
	}
	out := new(TSDBConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosCompact) DeepCopyInto(out *ThanosCompact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosCompact.
func (in *ThanosCompact) DeepCopy() *ThanosCompact {
	if in == nil {
		return nil
	}
	out := new(ThanosCompact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosCompact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosCompactList) DeepCopyInto(out *ThanosCompactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosCompact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosCompactList.
func (in *ThanosCompactList) DeepCopy() *ThanosCompactList {
	if in == nil {
		return nil
	}
	out := new(ThanosCompactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosCompactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosCompactSpec) DeepCopyInto(out *ThanosCompactSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.ObjectStorageConfig != nil {
		in, out := &in.ObjectStorageConfig, &out.ObjectStorageConfig
		*out = new(ObjectStorageConfig)
		(*in).DeepCopyInto(*out)
	}
	out.RetentionConfig = in.RetentionConfig
	out.BlockConfig = in.BlockConfig
	out.GroupConfig = in.GroupConfig
	if in.ShardingConfig != nil {
		in, out := &in.ShardingConfig, &out.ShardingConfig
		*out = new(ShardingConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosCompactSpec.
func (in *ThanosCompactSpec) DeepCopy() *ThanosCompactSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosCompactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosCompactStatus) DeepCopyInto(out *ThanosCompactStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosCompactStatus.
func (in *ThanosCompactStatus) DeepCopy() *ThanosCompactStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosCompactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosQuery) DeepCopyInto(out *ThanosQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosQuery.
func (in *ThanosQuery) DeepCopy() *ThanosQuery {
	if in == nil {
		return nil
	}
	out := new(ThanosQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosQueryList) DeepCopyInto(out *ThanosQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosQueryList.
func (in *ThanosQueryList) DeepCopy() *ThanosQueryList {
	if in == nil {
		return nil
	}
	out := new(ThanosQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosQuerySpec) DeepCopyInto(out *ThanosQuerySpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.QuerierReplicaLabels != nil {
		in, out := &in.QuerierReplicaLabels, &out.QuerierReplicaLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StoreLabelSelector != nil {
		in, out := &in.StoreLabelSelector, &out.StoreLabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.QueryFrontend != nil {
		in, out := &in.QueryFrontend, &out.QueryFrontend
		*out = new(QueryFrontendSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosQuerySpec.
func (in *ThanosQuerySpec) DeepCopy() *ThanosQuerySpec {
	if in == nil {
		return nil
	}
	out := new(ThanosQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosQueryStatus) DeepCopyInto(out *ThanosQueryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosQueryStatus.
func (in *ThanosQueryStatus) DeepCopy() *ThanosQueryStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosReceive) DeepCopyInto(out *ThanosReceive) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosReceive.
func (in *ThanosReceive) DeepCopy() *ThanosReceive {
	if in == nil {
		return nil
	}
	out := new(ThanosReceive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosReceive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosReceiveList) DeepCopyInto(out *ThanosReceiveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosReceive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosReceiveList.
func (in *ThanosReceiveList) DeepCopy() *ThanosReceiveList {
	if in == nil {
		return nil
	}
	out := new(ThanosReceiveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosReceiveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosReceiveSpec) DeepCopyInto(out *ThanosReceiveSpec) {
	*out = *in
	in.Router.DeepCopyInto(&out.Router)
	in.Ingester.DeepCopyInto(&out.Ingester)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosReceiveSpec.
func (in *ThanosReceiveSpec) DeepCopy() *ThanosReceiveSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosReceiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosReceiveStatus) DeepCopyInto(out *ThanosReceiveStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosReceiveStatus.
func (in *ThanosReceiveStatus) DeepCopy() *ThanosReceiveStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosReceiveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosRuler) DeepCopyInto(out *ThanosRuler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosRuler.
func (in *ThanosRuler) DeepCopy() *ThanosRuler {
	if in == nil {
		return nil
	}
	out := new(ThanosRuler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosRuler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosRulerList) DeepCopyInto(out *ThanosRulerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosRuler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosRulerList.
func (in *ThanosRulerList) DeepCopy() *ThanosRulerList {
	if in == nil {
		return nil
	}
	out := new(ThanosRulerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosRulerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosRulerSpec) DeepCopyInto(out *ThanosRulerSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.QueryLabelSelector != nil {
		in, out := &in.QueryLabelSelector, &out.QueryLabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.ObjectStorageConfig.DeepCopyInto(&out.ObjectStorageConfig)
	if in.RuleConfigSelector != nil {
		in, out := &in.RuleConfigSelector, &out.RuleConfigSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalLabels != nil {
		in, out := &in.ExternalLabels, &out.ExternalLabels
		*out = make(ExternalLabels, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AlertLabelDrop != nil {
		in, out := &in.AlertLabelDrop, &out.AlertLabelDrop
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosRulerSpec.
func (in *ThanosRulerSpec) DeepCopy() *ThanosRulerSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosRulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosRulerStatus) DeepCopyInto(out *ThanosRulerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosRulerStatus.
func (in *ThanosRulerStatus) DeepCopy() *ThanosRulerStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosRulerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosService) DeepCopyInto(out *ThanosService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosService.
func (in *ThanosService) DeepCopy() *ThanosService {
	if in == nil {
		return nil
	}
	out := new(ThanosService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosServiceList) DeepCopyInto(out *ThanosServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosServiceList.
func (in *ThanosServiceList) DeepCopy() *ThanosServiceList {
	if in == nil {
		return nil
	}
	out := new(ThanosServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosServiceSpec) DeepCopyInto(out *ThanosServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosServiceSpec.
func (in *ThanosServiceSpec) DeepCopy() *ThanosServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosServiceStatus) DeepCopyInto(out *ThanosServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosServiceStatus.
func (in *ThanosServiceStatus) DeepCopy() *ThanosServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStore) DeepCopyInto(out *ThanosStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStore.
func (in *ThanosStore) DeepCopy() *ThanosStore {
	if in == nil {
		return nil
	}
	out := new(ThanosStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStoreList) DeepCopyInto(out *ThanosStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ThanosStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStoreList.
func (in *ThanosStoreList) DeepCopy() *ThanosStoreList {
	if in == nil {
		return nil
	}
	out := new(ThanosStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStoreSpec) DeepCopyInto(out *ThanosStoreSpec) {
	*out = *in
	in.CommonThanosFields.DeepCopyInto(&out.CommonThanosFields)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.ObjectStorageConfig.DeepCopyInto(&out.ObjectStorageConfig)
	if in.IndexCacheConfig != nil {
		in, out := &in.IndexCacheConfig, &out.IndexCacheConfig
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.CachingBucketConfig != nil {
		in, out := &in.CachingBucketConfig, &out.CachingBucketConfig
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	out.ShardingStrategy = in.ShardingStrategy
	if in.MinTime != nil {
		in, out := &in.MinTime, &out.MinTime
		*out = new(Duration)
		**out = **in
	}
	if in.MaxTime != nil {
		in, out := &in.MaxTime, &out.MaxTime
		*out = new(Duration)
		**out = **in
	}
	in.Additional.DeepCopyInto(&out.Additional)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStoreSpec.
func (in *ThanosStoreSpec) DeepCopy() *ThanosStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStoreStatus) DeepCopyInto(out *ThanosStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStoreStatus.
func (in *ThanosStoreStatus) DeepCopy() *ThanosStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosStoreStatus)
	in.DeepCopyInto(out)
	return out
}
