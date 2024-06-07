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

// Headroom is the Schema for the Headrooms API
// +kubebuilder:object:root=true
// +kubebuilder:resource:path=headrooms,scope=Namespace,categories=karpenter
// +kubebuilder:subresource:status,scale
type Headroom struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   HeadroomSpec   `json:"spec"`
	Status HeadroomStatus `json:"status,omitempty"`
}

// HeadroomList contains a list of Headrooms
// +kubebuilder:object:root=true
type HeadroomList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Headroom `json:"items"`
}

type HeadroomSpec struct {
	// SatisfiedBy allows you to specify a selector against any pods that this Headroom should match against
	// Matching against pods with SatisfiedBy implies that the scheduling simulation will fill-up the headroom with pod
	// capacity as it schedules. e.g. Headroom for 500Mi memory and Pods that match that have 300Mi memory means the scheduling
	// simulation will only consider 200Mi memory for the headroom
	// +optional
	SatisfiedBy PodSelector `json:"satisfiedBy,omitempty"`
	// Template is a Pod Template to use for Headroom when scheduling
	// +optional
	Template v1.PodTemplateSpec `json:"template"`
	// Replicas is the number of pods to model from the template for Headroom during scheduling
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
}

type HeadroomStatus struct {
	// Conditions contains signals for health and readiness
	// It contains a status condition that indicates whether the capacity Headroom
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
