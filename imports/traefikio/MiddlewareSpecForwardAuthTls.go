package traefikio

// TLS defines the configuration used to secure the connection to the authentication server.
type MiddlewareSpecForwardAuthTls struct {
	// Deprecated: TLS client authentication is a server side option (see https://github.com/golang/go/blob/740a490f71d026bb7d2d13cb8fa2d6d6e0572b70/src/crypto/tls/common.go#L634).
	CaOptional *bool `field:"optional" json:"caOptional" yaml:"caOptional"`
	// CASecret is the name of the referenced Kubernetes Secret containing the CA to validate the server certificate.
	//
	// The CA certificate is extracted from key `tls.ca` or `ca.crt`.
	CaSecret *string `field:"optional" json:"caSecret" yaml:"caSecret"`
	// CertSecret is the name of the referenced Kubernetes Secret containing the client certificate.
	//
	// The client certificate is extracted from the keys `tls.crt` and `tls.key`.
	CertSecret *string `field:"optional" json:"certSecret" yaml:"certSecret"`
	// InsecureSkipVerify defines whether the server certificates should be validated.
	InsecureSkipVerify *bool `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
}
