package traefikio


// Spiffe defines the SPIFFE configuration.
type ServersTransportTcpSpecTlsSpiffe struct {
	// IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).
	Ids *[]*string `field:"optional" json:"ids" yaml:"ids"`
	// TrustDomain defines the allowed SPIFFE trust domain.
	TrustDomain *string `field:"optional" json:"trustDomain" yaml:"trustDomain"`
}

