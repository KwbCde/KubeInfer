package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InferenceJobSpec defines the desired state of InferenceJob
type InferenceJobSpec struct {
	// +kubebuilder:validation:Required
	Model string `json:"model"`
	Input string `json:"input"`
	Image string `json:"image"`
}

// InferenceJobStatus defines the observed state of InferenceJob.
type InferenceJobStatus struct {
	Phase      string             `json:"phase,omitempty"`
	Result     string             `json:"result,omitempty"`
	PodName    string             `json:"podName,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// InferenceJob is the Schema for the inferencejobs API
type InferenceJob struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of InferenceJob
	// +required
	Spec InferenceJobSpec `json:"spec"`

	// status defines the observed state of InferenceJob
	// +optional
	Status InferenceJobStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// InferenceJobList contains a list of InferenceJob
type InferenceJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []InferenceJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InferenceJob{}, &InferenceJobList{})
}
