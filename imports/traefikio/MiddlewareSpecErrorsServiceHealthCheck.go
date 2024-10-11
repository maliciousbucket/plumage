package traefikio

// Healthcheck defines health checks for ExternalName services.
type MiddlewareSpecErrorsServiceHealthCheck struct {
	// FollowRedirects defines whether redirects should be followed during the health check calls.
	//
	// Default: true.
	FollowRedirects *bool `field:"optional" json:"followRedirects" yaml:"followRedirects"`
	// Headers defines custom headers to be sent to the health check endpoint.
	Headers *map[string]*string `field:"optional" json:"headers" yaml:"headers"`
	// Hostname defines the value of hostname in the Host header of the health check request.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Interval defines the frequency of the health check calls.
	//
	// Default: 30s.
	Interval MiddlewareSpecErrorsServiceHealthCheckInterval `field:"optional" json:"interval" yaml:"interval"`
	// Method defines the healthcheck method.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Mode defines the health check mode.
	//
	// If defined to grpc, will use the gRPC health check protocol to probe the server.
	// Default: http.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Path defines the server URL path for the health check endpoint.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Port defines the server URL port for the health check endpoint.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Scheme replaces the server URL scheme for the health check endpoint.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// Status defines the expected HTTP status code of the response to the health check request.
	Status *float64 `field:"optional" json:"status" yaml:"status"`
	// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
	//
	// Default: 5s.
	Timeout MiddlewareSpecErrorsServiceHealthCheckTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}
