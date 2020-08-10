/*
Copyright 2020 Netflix, Inc.

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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type MachineType struct {
	ComputeResource `json:",inline"`
	Name            string            `json:"name"`
	Labels          map[string]string `json:"labels,omitempty"`
}

// +kubebuilder:object:root=true

type MachineTypeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MachineType `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

type MachineTypeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachineTypeConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MachineTypeConfig{}, &MachineTypeConfigList{})
}
