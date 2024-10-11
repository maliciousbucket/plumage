package traefikio

// Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection.
//
// If not defined, the `default` TLSOption is used.
// More info: https://doc.traefik.io/traefik/v3.1/https/tls/#tls-options
type IngressRouteSpecTlsOptions struct {
	// Name defines the name of the referenced TLSOption.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/providers/kubernetes-crd/#kind-tlsoption
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace defines the namespace of the referenced TLSOption.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/providers/kubernetes-crd/#kind-tlsoption
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}
