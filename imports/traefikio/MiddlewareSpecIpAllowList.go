package traefikio


// IPAllowList holds the IP allowlist middleware configuration.
//
// This middleware limits allowed requests based on the client IP.
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/ipallowlist/
type MiddlewareSpecIpAllowList struct {
	// IPStrategy holds the IP strategy configuration used by Traefik to determine the client IP.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/ipallowlist/#ipstrategy
	IpStrategy *MiddlewareSpecIpAllowListIpStrategy `field:"optional" json:"ipStrategy" yaml:"ipStrategy"`
	// RejectStatusCode defines the HTTP status code used for refused requests.
	//
	// If not set, the default is 403 (Forbidden).
	RejectStatusCode *float64 `field:"optional" json:"rejectStatusCode" yaml:"rejectStatusCode"`
	// SourceRange defines the set of allowed IPs (or ranges of allowed IPs by using CIDR notation).
	SourceRange *[]*string `field:"optional" json:"sourceRange" yaml:"sourceRange"`
}

