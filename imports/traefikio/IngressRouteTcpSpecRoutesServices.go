package traefikio

// ServiceTCP defines an upstream TCP service to proxy traffic to.
type IngressRouteTcpSpecRoutesServices struct {
	// Name defines the name of the referenced Kubernetes Service.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port IngressRouteTcpSpecRoutesServicesPort `field:"required" json:"port" yaml:"port"`
	// Namespace defines the namespace of the referenced Kubernetes Service.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// NativeLB controls, when creating the load-balancer, whether the LB's children are directly the pods IPs or if the only child is the Kubernetes Service clusterIP.
	//
	// The Kubernetes Service itself does load-balance to the pods.
	// By default, NativeLB is false.
	NativeLb *bool `field:"optional" json:"nativeLb" yaml:"nativeLb"`
	// NodePortLB controls, when creating the load-balancer, whether the LB's children are directly the nodes internal IPs using the nodePort when the service type is NodePort.
	//
	// It allows services to be reachable when Traefik runs externally from the Kubernetes cluster but within the same network of the nodes.
	// By default, NodePortLB is false.
	NodePortLb *bool `field:"optional" json:"nodePortLb" yaml:"nodePortLb"`
	// ProxyProtocol defines the PROXY protocol configuration.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#proxy-protocol
	ProxyProtocol *IngressRouteTcpSpecRoutesServicesProxyProtocol `field:"optional" json:"proxyProtocol" yaml:"proxyProtocol"`
	// ServersTransport defines the name of ServersTransportTCP resource to use.
	//
	// It allows to configure the transport between Traefik and your servers.
	// Can only be used on a Kubernetes Service.
	ServersTransport *string `field:"optional" json:"serversTransport" yaml:"serversTransport"`
	// TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection.
	//
	// It is a duration in milliseconds, defaulting to 100.
	// A negative value means an infinite deadline (i.e. the reading capability is never closed).
	// Deprecated: TerminationDelay is not supported APIVersion traefik.io/v1, please use ServersTransport to configure the TerminationDelay instead.
	TerminationDelay *float64 `field:"optional" json:"terminationDelay" yaml:"terminationDelay"`
	// TLS determines whether to use TLS when dialing with the backend.
	Tls *bool `field:"optional" json:"tls" yaml:"tls"`
	// Weight defines the weight used when balancing requests between multiple Kubernetes Service.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}
