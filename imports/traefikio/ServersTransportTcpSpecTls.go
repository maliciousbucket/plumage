package traefikio


// TLS defines the TLS configuration.
type ServersTransportTcpSpecTls struct {
	// CertificatesSecrets defines a list of secret storing client certificates for mTLS.
	CertificatesSecrets *[]*string `field:"optional" json:"certificatesSecrets" yaml:"certificatesSecrets"`
	// InsecureSkipVerify disables TLS certificate verification.
	InsecureSkipVerify *bool `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.
	//
	// PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.
	PeerCertUri *string `field:"optional" json:"peerCertUri" yaml:"peerCertUri"`
	// RootCAsSecrets defines a list of CA secret used to validate self-signed certificates.
	RootCAsSecrets *[]*string `field:"optional" json:"rootCAsSecrets" yaml:"rootCAsSecrets"`
	// ServerName defines the server name used to contact the server.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// Spiffe defines the SPIFFE configuration.
	Spiffe *ServersTransportTcpSpecTlsSpiffe `field:"optional" json:"spiffe" yaml:"spiffe"`
}

