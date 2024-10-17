package traefikio


// IPAllowList defines the IPAllowList middleware configuration.
//
// This middleware accepts/refuses connections based on the client IP.
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/tcp/ipallowlist/
type MiddlewareTcpSpecIpAllowList struct {
	// SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}

