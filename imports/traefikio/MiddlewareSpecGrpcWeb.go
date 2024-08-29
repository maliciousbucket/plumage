package traefikio


// GrpcWeb holds the gRPC web middleware configuration.
//
// This middleware converts a gRPC web request to an HTTP/2 gRPC request.
type MiddlewareSpecGrpcWeb struct {
	// AllowOrigins is a list of allowable origins.
	//
	// Can also be a wildcard origin "*".
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
}

