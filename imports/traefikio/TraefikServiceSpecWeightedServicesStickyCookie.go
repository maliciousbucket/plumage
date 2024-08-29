package traefikio


// Cookie defines the sticky cookie configuration.
type TraefikServiceSpecWeightedServicesStickyCookie struct {
	// HTTPOnly defines whether the cookie can be accessed by client-side APIs, such as JavaScript.
	HttpOnly *bool `field:"optional" json:"httpOnly" yaml:"httpOnly"`
	// MaxAge indicates the number of seconds until the cookie expires.
	//
	// When set to a negative number, the cookie expires immediately.
	// When set to zero, the cookie never expires.
	MaxAge *float64 `field:"optional" json:"maxAge" yaml:"maxAge"`
	// Name defines the Cookie name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// SameSite defines the same site policy.
	//
	// More info: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie/SameSite
	SameSite *string `field:"optional" json:"sameSite" yaml:"sameSite"`
	// Secure defines whether the cookie can only be transmitted over an encrypted connection (i.e. HTTPS).
	Secure *bool `field:"optional" json:"secure" yaml:"secure"`
}

