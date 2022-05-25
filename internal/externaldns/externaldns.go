package externaldns

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type TTL int64

type Targets []string

type ProviderSpecificProperty struct {
	Name  string
	Value string
}

type Labels map[string]string

type ProviderSpecific []ProviderSpecificProperty

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

type DNSEndpointSpec struct {
	Endpoints []*Endpoint `json:"endpoints,omitempty"`
}

type DNSEndpointStatus struct {
	// The generation observed by by the external-dns controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}

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
