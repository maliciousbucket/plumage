package traefikio


// ProxyProtocol defines the PROXY protocol configuration.
//
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#proxy-protocol
type IngressRouteTcpSpecRoutesServicesProxyProtocol struct {
	// Version defines the PROXY Protocol version to use.
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

