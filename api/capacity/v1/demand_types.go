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

// ResourceGroup is capacity as a pair of count and ComputeResource
type ResourceGroup struct {
	Count              uint32 `json:"count"`
	v1.ComputeResource `json:"resource"`
}

type ResourceGroups []ResourceGroup

type ResourcePoolName string

// DemandSpec defines the desired state of Demand
type DemandSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CapacityGroup string `json:"capacity_group,omitempty"`
	// Reservations are capacity reservations defined in CapacityGroup
	Reservations ResourceGroups `json:"reservations,omitempty"`
}

// DemandStatus defines the observed state of Demand
type DemandStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Bound is aggregate demand used by running pods
	Bound map[ResourcePoolName]v1.ComputeResource `json:"bound,omitempty"`
	// Unbound is demand waiting for capacity
	Unbound map[ResourcePoolName]ResourceGroups `json:"unbound,omitempty"`
	// ReservedUnallocated is demand from unused reservations
	ReservedUnallocated map[ResourcePoolName]ResourceGroups `json:"reserved_unallocated,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=cd
//+kubebuilder:subresource:status
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Demand is the Schema for the demands API
type Demand struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DemandSpec   `json:"spec,omitempty"`
	Status DemandStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DemandList contains a list of Demand
type DemandList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Demand `json:"items"`
}
