package traefikio

// StripPrefix holds the strip prefix middleware configuration.
//
// This middleware removes the specified prefixes from the URL path.
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/stripprefix/
type MiddlewareSpecStripPrefix struct {
	// Deprecated: ForceSlash option is deprecated, please remove any usage of this option.
	//
	// ForceSlash ensures that the resulting stripped path is not the empty string, by replacing it with / when necessary.
	// Default: true.
	ForceSlash *bool `field:"optional" json:"forceSlash" yaml:"forceSlash"`
	// Prefixes defines the prefixes to strip from the request URL.
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}
