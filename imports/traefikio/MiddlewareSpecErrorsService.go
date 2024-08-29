package traefikio


// Service defines the reference to a Kubernetes Service that will serve the error page.
//
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/errorpages/#service
type MiddlewareSpecErrorsService struct {
	// Name defines the name of the referenced Kubernetes Service or TraefikService.
	//
	// The differentiation between the two is specified in the Kind field.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Healthcheck defines health checks for ExternalName services.
	HealthCheck *MiddlewareSpecErrorsServiceHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Kind defines the kind of the Service.
	Kind MiddlewareSpecErrorsServiceKind `field:"optional" json:"kind" yaml:"kind"`
	// Namespace defines the namespace of the referenced Kubernetes Service or TraefikService.
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
	// PassHostHeader defines whether the client Host header is forwarded to the upstream Kubernetes Service.
	//
	// By default, passHostHeader is true.
	PassHostHeader *bool `field:"optional" json:"passHostHeader" yaml:"passHostHeader"`
	// Port defines the port of a Kubernetes Service.
	//
	// This can be a reference to a named port.
	Port MiddlewareSpecErrorsServicePort `field:"optional" json:"port" yaml:"port"`
	// ResponseForwarding defines how Traefik forwards the response from the upstream Kubernetes Service to the client.
	ResponseForwarding *MiddlewareSpecErrorsServiceResponseForwarding `field:"optional" json:"responseForwarding" yaml:"responseForwarding"`
	// Scheme defines the scheme to use for the request to the upstream Kubernetes Service.
	//
	// It defaults to https when Kubernetes Service port is 443, http otherwise.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// ServersTransport defines the name of ServersTransport resource to use.
	//
	// It allows to configure the transport between Traefik and your servers.
	// Can only be used on a Kubernetes Service.
	ServersTransport *string `field:"optional" json:"serversTransport" yaml:"serversTransport"`
	// Sticky defines the sticky sessions configuration.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#sticky-sessions
	Sticky *MiddlewareSpecErrorsServiceSticky `field:"optional" json:"sticky" yaml:"sticky"`
	// Strategy defines the load balancing strategy between the servers.
	//
	// RoundRobin is the only supported value at the moment.
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
	// Weight defines the weight and should only be specified when Name references a TraefikService object (and to be precise, one that embeds a Weighted Round Robin).
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

