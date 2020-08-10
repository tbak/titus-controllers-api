package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ResourceShape struct {
	ComputeResource `json:",inline"`

	Labels map[string]string `json:"labels,omitempty"`
}

type ResourceDemandStatus struct {
	// Resource provisioning status (Pending, Rejected, Fulfilled)
	Status string `json:"status"`

	// Human readable message.
	StatusMessage string `json:"statusMessage"`

	// Time at which the report was made.
	ReportedAt uint64 `json:"reportedAt"`
}

type ResourcePoolSpec struct {
	// Resource pool name (for example 'critical', 'elastic', 'gpu.p4', etc).
	Name string `json:"name"`

	// Preferential machine shape (minimum size and resource ratio). A total amount of requested resource is
	// defined as resource_shape * resource_count.
	ResourceShape ResourceShape `json:"resourceShape"`

	// Number of resource shapes requested. A total amount of requested resource is
	// defined as resource_shape * resource_count.
	ResourceCount int64 `json:"resourceCount"`

	// Time at which the last resource request was made.
	RequestedAt int64 `json:"requestedAt"`

	// Resource demand fulfillment report.
	Status ResourceDemandStatus `json:"resourceDemandStatus"`
}

// +kubebuilder:object:root=true

type ResourcePoolConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ResourcePoolSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

type ResourcePoolConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePoolConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourcePoolConfig{}, &ResourcePoolConfigList{})
}
