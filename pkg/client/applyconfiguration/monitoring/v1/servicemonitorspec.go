// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceMonitorSpecApplyConfiguration represents an declarative configuration of the ServiceMonitorSpec type for use
// with apply.
type ServiceMonitorSpecApplyConfiguration struct {
	JobLabel              *string                              `json:"jobLabel,omitempty"`
	TargetLabels          []string                             `json:"targetLabels,omitempty"`
	PodTargetLabels       []string                             `json:"podTargetLabels,omitempty"`
	Endpoints             []EndpointApplyConfiguration         `json:"endpoints,omitempty"`
	Selector              *metav1.LabelSelector                `json:"selector,omitempty"`
	NamespaceSelector     *NamespaceSelectorApplyConfiguration `json:"namespaceSelector,omitempty"`
	SampleLimit           *uint64                              `json:"sampleLimit,omitempty"`
	TargetLimit           *uint64                              `json:"targetLimit,omitempty"`
	LabelLimit            *uint64                              `json:"labelLimit,omitempty"`
	LabelNameLengthLimit  *uint64                              `json:"labelNameLengthLimit,omitempty"`
	LabelValueLengthLimit *uint64                              `json:"labelValueLengthLimit,omitempty"`
	KeepDroppedTargets    *uint64                              `json:"keepDroppedTargets,omitempty"`
	AttachMetadata        *AttachMetadataApplyConfiguration    `json:"attachMetadata,omitempty"`
}

// ServiceMonitorSpecApplyConfiguration constructs an declarative configuration of the ServiceMonitorSpec type for use with
// apply.
func ServiceMonitorSpec() *ServiceMonitorSpecApplyConfiguration {
	return &ServiceMonitorSpecApplyConfiguration{}
}

// WithJobLabel sets the JobLabel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the JobLabel field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithJobLabel(value string) *ServiceMonitorSpecApplyConfiguration {
	b.JobLabel = &value
	return b
}

// WithTargetLabels adds the given value to the TargetLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetLabels field.
func (b *ServiceMonitorSpecApplyConfiguration) WithTargetLabels(values ...string) *ServiceMonitorSpecApplyConfiguration {
	for i := range values {
		b.TargetLabels = append(b.TargetLabels, values[i])
	}
	return b
}

// WithPodTargetLabels adds the given value to the PodTargetLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodTargetLabels field.
func (b *ServiceMonitorSpecApplyConfiguration) WithPodTargetLabels(values ...string) *ServiceMonitorSpecApplyConfiguration {
	for i := range values {
		b.PodTargetLabels = append(b.PodTargetLabels, values[i])
	}
	return b
}

// WithEndpoints adds the given value to the Endpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Endpoints field.
func (b *ServiceMonitorSpecApplyConfiguration) WithEndpoints(values ...*EndpointApplyConfiguration) *ServiceMonitorSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithEndpoints")
		}
		b.Endpoints = append(b.Endpoints, *values[i])
	}
	return b
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithSelector(value metav1.LabelSelector) *ServiceMonitorSpecApplyConfiguration {
	b.Selector = &value
	return b
}

// WithNamespaceSelector sets the NamespaceSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NamespaceSelector field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithNamespaceSelector(value *NamespaceSelectorApplyConfiguration) *ServiceMonitorSpecApplyConfiguration {
	b.NamespaceSelector = value
	return b
}

// WithSampleLimit sets the SampleLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampleLimit field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithSampleLimit(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.SampleLimit = &value
	return b
}

// WithTargetLimit sets the TargetLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetLimit field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithTargetLimit(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.TargetLimit = &value
	return b
}

// WithLabelLimit sets the LabelLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelLimit field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithLabelLimit(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.LabelLimit = &value
	return b
}

// WithLabelNameLengthLimit sets the LabelNameLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelNameLengthLimit field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithLabelNameLengthLimit(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.LabelNameLengthLimit = &value
	return b
}

// WithLabelValueLengthLimit sets the LabelValueLengthLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LabelValueLengthLimit field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithLabelValueLengthLimit(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.LabelValueLengthLimit = &value
	return b
}

// WithKeepDroppedTargets sets the KeepDroppedTargets field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeepDroppedTargets field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithKeepDroppedTargets(value uint64) *ServiceMonitorSpecApplyConfiguration {
	b.KeepDroppedTargets = &value
	return b
}

// WithAttachMetadata sets the AttachMetadata field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AttachMetadata field is set to the value of the last call.
func (b *ServiceMonitorSpecApplyConfiguration) WithAttachMetadata(value *AttachMetadataApplyConfiguration) *ServiceMonitorSpecApplyConfiguration {
	b.AttachMetadata = value
	return b
}
