package traefikio

// MiddlewareTCPSpec defines the desired state of a MiddlewareTCP.
type MiddlewareTcpSpec struct {
	// InFlightConn defines the InFlightConn middleware configuration.
	InFlightConn *MiddlewareTcpSpecInFlightConn `field:"optional" json:"inFlightConn" yaml:"inFlightConn"`
	// IPAllowList defines the IPAllowList middleware configuration.
	//
	// This middleware accepts/refuses connections based on the client IP.
	// More info: https://doc.traefik.io/traefik/v3.1/middlewares/tcp/ipallowlist/
	IpAllowList *MiddlewareTcpSpecIpAllowList `field:"optional" json:"ipAllowList" yaml:"ipAllowList"`
	// IPWhiteList defines the IPWhiteList middleware configuration.
	//
	// This middleware accepts/refuses connections based on the client IP.
	// Deprecated: please use IPAllowList instead.
	// More info: https://doc.traefik.io/traefik/v3.1/middlewares/tcp/ipwhitelist/
	IpWhiteList *MiddlewareTcpSpecIpWhiteList `field:"optional" json:"ipWhiteList" yaml:"ipWhiteList"`
}
