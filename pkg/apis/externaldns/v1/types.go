package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSEndpoint is the CRD wrapper for Endpoint
// +k8s:openapi-gen=true
// +kubebuilder:resource:path=dnsendpoints
// +kubebuilder:subresource:status
type DNSEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSEndpointSpec   `json:"spec,omitempty"`
	Status DNSEndpointStatus `json:"status,omitempty"`
}

type DNSEndpointStatus struct {
	// The generation observed by by the external-dns controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

type DNSEndpointSpec struct {
	Endpoints []*Endpoint `json:"endpoints,omitempty"`
}

type Endpoint struct {
	// The hostname for the DNS record
	DNSName string `json:"dnsName,omitempty"`

	// The targets the DNS service points to
	Targets Targets `json:"targets,omitempty"`

	// RecordType type of record, e.g. CNAME, A, SRV, TXT, MX
	RecordType string `json:"recordType,omitempty"`

	// TTL for the record
	RecordTTL TTL `json:"recordTTL,omitempty"`

	// Labels stores labels defined for the Endpoint
	// +optional
	Labels Labels `json:"labels,omitempty"`

	// ProviderSpecific stores provider specific config
	// +optional
	ProviderSpecific ProviderSpecific `json:"providerSpecific,omitempty"`
}

type ProviderSpecific []ProviderSpecificProperty

type ProviderSpecificProperty struct {
	// Name of the property
	Name string `json:"name,omitempty"`
	// Value of the property
	Value string `json:"value,omitempty"`
}

type TTL int64

type Targets []string

type Labels map[string]string

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSEndpointList is a list of the DNSEndpoint resources.
type DNSEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DNSEndpoint `json:"items"`
}
