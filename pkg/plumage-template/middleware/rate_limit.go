package middleware

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

const (
	defaultRateLimitAverage = 50
	defaultRateLimitBurst   = 80
	defaultRateLimitPeriod  = "1m"
	defaultRateLimitIpDepth = 2
)

type RateLimitProps struct {
	AverageRequests int
	BurstRequests   int
	RatePeriod      string
	LimitStrategy   resilience.RateLimitStrategy
}

func NewRateLimitMiddleware(scope constructs.Construct, id string, ns string, appLabel string, props *RateLimitProps) traefikio.Middleware {
	if props == nil {
		return nil
	}
	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRateLimit)
	spec := &traefikio.MiddlewareSpecRateLimit{}
	if props.AverageRequests != 0 {
		spec.Average = jsii.Number(props.AverageRequests)
	}

	if props.BurstRequests != 0 {
		spec.Burst = jsii.Number(props.BurstRequests)
	}

	if props.RatePeriod != "" {
		spec.Period = traefikio.MiddlewareSpecRateLimitPeriod_FromString(jsii.String(props.RatePeriod))
	}

	if props.LimitStrategy != nil {
		strategy := props.LimitStrategy
		switch s := strategy.(type) {
		case *resilience.IpDepthStrategy:
			spec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				IpStrategy: &traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy{
					Depth: jsii.Number(s.Depth),
				},
			}
		case *resilience.RequestHeaderStrategy:
			spec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHeaderName: jsii.String(s.HeaderName),
			}
		case *resilience.HostStrategy:
			spec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHost: jsii.Bool(s.Host),
			}
		}

	}
	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			RateLimit: spec,
		},
	})
}

func NewDefaultRateLimitMiddleware(scope constructs.Construct, id string, ns string, appLabel string) traefikio.Middleware {
	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRateLimit)
	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			RateLimit: &traefikio.MiddlewareSpecRateLimit{
				Average: jsii.Number(defaultRateLimitAverage),
				Burst:   jsii.Number(defaultRateLimitBurst),
				Period:  traefikio.MiddlewareSpecRateLimitPeriod_FromString(jsii.String(defaultRateLimitPeriod)),
				SourceCriterion: &traefikio.MiddlewareSpecRateLimitSourceCriterion{
					IpStrategy: &traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy{
						Depth: jsii.Number(defaultRateLimitIpDepth),
					},
				},
			},
		},
	})
}
