package traefikio


// ContentType holds the content-type middleware configuration.
//
// This middleware exists to enable the correct behavior until at least the default one can be changed in a future version.
type MiddlewareSpecContentType struct {
	// AutoDetect specifies whether to let the `Content-Type` header, if it has not been set by the backend, be automatically set to a value derived from the contents of the response.
	//
	// Deprecated: AutoDetect option is deprecated, Content-Type middleware is only meant to be used to enable the content-type detection, please remove any usage of this option.
	AutoDetect *bool `field:"optional" json:"autoDetect" yaml:"autoDetect"`
}

