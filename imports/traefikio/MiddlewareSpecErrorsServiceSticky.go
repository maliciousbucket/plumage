package traefikio

// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#sticky-sessions
type MiddlewareSpecErrorsServiceSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *MiddlewareSpecErrorsServiceStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}
