package traefikio

// Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection.
//
// If not defined, the `default` TLSOption is used.
// More info: https://doc.traefik.io/traefik/v3.1/https/tls/#tls-options
type IngressRouteTcpSpecTlsOptions struct {
	// Name defines the name of the referenced Traefik resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced Traefik resource.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}
