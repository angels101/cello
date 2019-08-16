package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DNSpec defines the desired state of CA
// +k8s:openapi-gen=true
type DNSpec struct {
	CN string `json:"cn,omitempty"`
	C  string `json:"c"`
	ST string `json:"st"`
	L  string `json:"l,omitempty"`
	O  string `json:"o"`
	OU string `json:"ou"`
}

// CACerts defines the desired state of CA
// +k8s:openapi-gen=true
type CACerts struct {
	Cert    string `json:"cert,omitempty"`
	Key     string `json:"key,omitempty"`
	TLSCert string `json:"tlsCert,omitempty"`
	TLSKey  string `json:"tlsKey,omitempty"`
}

// CASpec defines the desired state of CA
// +k8s:openapi-gen=true
type CASpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Admin         string   `json:"admin"`
	AdminPassword string   `json:"adminPassword"`
	StorageSize   string   `json:"storageSize,omitempty"`
	StorageClass  string   `json:"storageClass,omitempty"`
	Image         string   `json:"image,omitempty"`
	DN            *DNSpec  `json:"dn,omitempty"`
	Certs         *CACerts `json:"certs,omitempty"`
	Hosts         []string `json:"hosts,omitempty"`
}

// CAStatus defines the observed state of CA
// +k8s:openapi-gen=true
type CAStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	AccessPoint string `json:"accessPoint"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CA is the Schema for the cas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type CA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CASpec   `json:"spec,omitempty"`
	Status CAStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CAList contains a list of CA
type CAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CA `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CA{}, &CAList{})
}
