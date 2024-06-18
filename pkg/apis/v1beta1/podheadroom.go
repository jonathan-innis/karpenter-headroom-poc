/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"github.com/awslabs/operatorpkg/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodHeadroom is the Schema for the PodHeadrooms API
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=podheadrooms,scope=Namespace,categories=karpenter
// +kubebuilder:subresource:status,scale
type PodHeadroom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   PodHeadroomSpec   `json:"spec"`
	Status PodHeadroomStatus `json:"status,omitempty"`
}

// PodHeadroomList contains a list of PodHeadrooms
// +kubebuilder:object:root=true
type PodHeadroomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodHeadroom `json:"items"`
}

type PodHeadroomSpec struct {
	// SatisfiedBy allows you to specify a selector against any pods that this PodHeadroom should match against
	// Matching against pods with SatisfiedBy implies that the scheduling simulation will fill-up the pod headroom with pod
	// capacity as it schedules. e.g. PodHeadroom for 500Mi memory and Pods that match that have 300Mi memory means the scheduling
	// simulation will only consider 200Mi memory for the pod headroom
	// +optional
	SatisfiedBy PodSelector `json:"satisfiedBy,omitempty"`
	// Template is a Pod Template to use for PodHeadroom when scheduling
	// +optional
	Template v1.PodTemplateSpec `json:"template"`
	// Replicas is the number of pods to model from the template for PodHeadroom during scheduling
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
}

type PodHeadroomStatus struct {
	// Conditions contains signals for health and readiness
	// It contains a status condition that indicates whether the capacity PodHeadroom
	// is available on the cluster, allowing external components the ability to know
	// when all the capacity exists
	// +optional
	Conditions []status.Condition `json:"conditions,omitempty"`
}

type PodSelector struct {
	metav1.LabelSelector `json:",inline"`
	// MatchFields is a list of selector requirements by fields
	// +optional
	MatchFields metav1.LabelSelectorRequirement `json:"matchFields,omitempty"`
}
