package traefikio

// Sticky defines the sticky sessions configuration.
//
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#sticky-sessions
type TraefikServiceSpecMirroringMirrorsSticky struct {
	// Cookie defines the sticky cookie configuration.
	Cookie *TraefikServiceSpecMirroringMirrorsStickyCookie `field:"optional" json:"cookie" yaml:"cookie"`
}
