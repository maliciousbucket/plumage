package traefikio


// IPWhiteList defines the IPWhiteList middleware configuration.
//
// This middleware accepts/refuses connections based on the client IP.
// Deprecated: please use IPAllowList instead.
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/tcp/ipwhitelist/
type MiddlewareTcpSpecIpWhiteList struct {
	// SourceRange defines the allowed IPs (or ranges of allowed IPs by using CIDR notation).
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}

