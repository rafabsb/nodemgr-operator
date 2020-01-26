package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodemgrAnnotations define the annotations that must be a valid json.
type NodemgrAnnotations struct {
	// +optional
	LB []NodemgrLB `json:"lb,omitempty"`
}

// NodemgrLB defines a LB that points to the node.
type NodemgrLB struct {
	Name      string               `json:"name"`
	DNS       string               `json:"dns"`
	Instances []NodemgrLBInstances `json:"instances"`
}

// NodemgrLBInstances instances of a LB.
type NodemgrLBInstances struct {
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Site      string  `json:"site"`
	Network   string  `json:"network"`
	IP        string  `json:"ip"`
	LocalPort []int32 `json:"localport"`
}

// NodemgrTaint defines the Taint that a node has.
type NodemgrTaint struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
	Value  string `json:"value"`
}

// NodemgrSpec defines the desired state of Nodemgr
type NodemgrSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Labels map[string]string `json:"labels"`
	// +optional
	Taints []NodemgrTaint `json:"taints,omitempty"`
	// +optional
	Annotations NodemgrAnnotations `json:"annotations,omitempty"`
}

// NodemgrStatus defines the observed state of Nodemgr
type NodemgrStatus struct {
	Phase string `json:"phase"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Nodemgr is the Schema for the nodemgrs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=nodemgrs,scope=Cluster
// +kubebuilder:printcolumn:name="Status",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
type Nodemgr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodemgrSpec   `json:"spec,omitempty"`
	Status NodemgrStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodemgrList contains a list of Nodemgr
type NodemgrList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Nodemgr `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Nodemgr{}, &NodemgrList{})
}
