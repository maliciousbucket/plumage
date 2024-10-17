package traefikio


// Route holds the HTTP route configuration.
type IngressRouteSpecRoutes struct {
	// Kind defines the kind of the route.
	//
	// Rule is the only supported kind.
	Kind IngressRouteSpecRoutesKind `field:"required" json:"kind" yaml:"kind"`
	// Match defines the router's rule.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#rule
	Match *string `field:"required" json:"match" yaml:"match"`
	// Middlewares defines the list of references to Middleware resources.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/providers/kubernetes-crd/#kind-middleware
	Middlewares *[]*IngressRouteSpecRoutesMiddlewares `field:"optional" json:"middlewares" yaml:"middlewares"`
	// Priority defines the router's priority.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#priority
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Services defines the list of Service.
	//
	// It can contain any combination of TraefikService and/or reference to a Kubernetes Service.
	Services *[]*IngressRouteSpecRoutesServices `field:"optional" json:"services" yaml:"services"`
	// Syntax defines the router's rule syntax.
	//
	// More info: https://doc.traefik.io/traefik/v3.1/routing/routers/#rulesyntax
	Syntax *string `field:"optional" json:"syntax" yaml:"syntax"`
}

