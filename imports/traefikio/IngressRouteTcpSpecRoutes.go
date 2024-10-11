package traefikio

// RouteTCP holds the TCP route configuration.
type IngressRouteTcpSpecRoutes struct {
	// Match defines the router's rule.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#rule_1
	Match *string `field:"required" json:"match" yaml:"match"`
	// Middlewares defines the list of references to MiddlewareTCP resources.
	Middlewares *[]*IngressRouteTcpSpecRoutesMiddlewares `field:"optional" json:"middlewares" yaml:"middlewares"`
	// Priority defines the router's priority.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#priority_1
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Services defines the list of TCP services.
	Services *[]*IngressRouteTcpSpecRoutesServices `field:"optional" json:"services" yaml:"services"`
	// Syntax defines the router's rule syntax.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#rulesyntax_1
	Syntax *string `field:"optional" json:"syntax" yaml:"syntax"`
}
