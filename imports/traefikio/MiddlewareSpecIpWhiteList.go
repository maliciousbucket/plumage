package traefikio

// Deprecated: please use IPAllowList instead.
type MiddlewareSpecIpWhiteList struct {
	// IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/ipallowlist/#ipstrategy
	IpStrategy *MiddlewareSpecIpWhiteListIpStrategy `field:"optional" json:"ipStrategy" yaml:"ipStrategy"`
	// SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).
	//
	// Required.
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}
