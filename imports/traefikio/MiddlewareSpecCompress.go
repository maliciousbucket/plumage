package traefikio


// Compress holds the compress middleware configuration.
//
// This middleware compresses responses before sending them to the client, using gzip, brotli, or zstd compression.
// More info: https://doc.traefik.io/traefik/v3.1/middlewares/http/compress/
type MiddlewareSpecCompress struct {
	// DefaultEncoding specifies the default encoding if the `Accept-Encoding` header is not in the request or contains a wildcard (`*`).
	DefaultEncoding *string `field:"optional" json:"defaultEncoding" yaml:"defaultEncoding"`
	// Encodings defines the list of supported compression algorithms.
	Encodings *[]*string `field:"optional" json:"encodings" yaml:"encodings"`
	// ExcludedContentTypes defines the list of content types to compare the Content-Type header of the incoming requests and responses before compressing.
	//
	// `application/grpc` is always excluded.
	ExcludedContentTypes *[]*string `field:"optional" json:"excludedContentTypes" yaml:"excludedContentTypes"`
	// IncludedContentTypes defines the list of content types to compare the Content-Type header of the responses before compressing.
	IncludedContentTypes *[]*string `field:"optional" json:"includedContentTypes" yaml:"includedContentTypes"`
	// MinResponseBodyBytes defines the minimum amount of bytes a response body must have to be compressed.
	//
	// Default: 1024.
	MinResponseBodyBytes *float64 `field:"optional" json:"minResponseBodyBytes" yaml:"minResponseBodyBytes"`
}

