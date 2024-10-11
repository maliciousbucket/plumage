package traefikio

// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#sticky-sessions
type TraefikServiceSpecMirroringSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *TraefikServiceSpecMirroringStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}
