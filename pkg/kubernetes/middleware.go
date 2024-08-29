package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

const (
	MiddlewareKind = "Middleware"
)

func NewMiddleware(scope constructs.Construct, id string) cdk8s.Chart {

	middleware := traefikio.MiddlewareProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Annotations:     nil,
			Finalizers:      nil,
			Labels:          nil,
			Name:            nil,
			Namespace:       nil,
			OwnerReferences: nil,
		},
		Spec: &traefikio.MiddlewareSpec{
			AddPrefix:         nil,
			BasicAuth:         nil,
			Buffering:         nil,
			Chain:             nil,
			CircuitBreaker:    nil,
			Compress:          nil,
			ContentType:       nil,
			DigestAuth:        nil,
			Errors:            nil,
			ForwardAuth:       nil,
			GrpcWeb:           nil,
			Headers:           nil,
			InFlightReq:       nil,
			IpAllowList:       nil,
			IpWhiteList:       nil,
			PassTlsClientCert: nil,
			Plugin:            nil,
			RateLimit:         nil,
			RedirectRegex:     nil,
			RedirectScheme:    nil,
			ReplacePath:       nil,
			ReplacePathRegex:  nil,
			Retry:             nil,
			StripPrefix:       nil,
			StripPrefixRegex:  nil,
		},
	}
	return nil
}

type CircuitBreakerSpec interface {
	CircuitBreakerExpression() string
	Validate() error
}

type MiddlewareProps struct {
}
