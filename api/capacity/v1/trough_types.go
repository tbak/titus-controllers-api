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

package v1

import (
	v1 "github.com/Netflix/titus-controllers-api/api/resourcepool/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TroughSpec defines the desired state of Trough
type TroughSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ResourcePool is the host of this trough
	ResourcePool ResourcePoolName `json:"resource_pool"`

	// Percent is the amount of trough capacity to export as allocatable
	Percent uint32 `json:"percent"`

	// SchedulerName is the name of the scheduler profile to use for scheduling pods into this trough
	SchedulerName string `json:"scheduler_name,omitempty"`
}

// TroughStatus is the capacity available as trough
type TroughStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Capacity is total trough capacity. Only part of this may be available as Allocatable depending on spec.percent
	Capacity *v1.ComputeResource `json:"capacity,omitempty"`

	// Allocatable is subset of the total trough capacity that is considered allocatable
	Allocatable *v1.ComputeResource `json:"allocatable,omitempty"`

	// Free is the available trough for immediate use. No new trough user pods should be admitted if this is zero.
	Free *v1.ComputeResource `json:"free,omitempty"`

	// LastUpdated is the time in RFC3339 format of last update to trough
	LastUpdated string `json:"last_updated,omitempty"`

	// LastEvaluated is the time in RFC3339 format of last trough availability evaluation
	LastEvaluated string `json:"last_evaluated,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=tr
//+kubebuilder:subresource:status
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Trough is the Schema for the Troughs API
type Trough struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TroughSpec   `json:"spec,omitempty"`
	Status TroughStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TroughList contains a list of Trough
type TroughList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Trough `json:"items"`
}
