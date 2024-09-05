package k8s

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

const (
	MiddlewareTypeRetry          = "retry"
	MiddlewareTypeRateLimit      = "rate-limit"
	MiddlewareTypeCircuitBreaker = "circuit-breaker"
)

type RetryProps struct {
	RetryAttempts int
	IntervalMs    string
}

type RateLimitProps struct {
	AverageRequests int
	BurstRequests   int
	RatePeriod      string
	LimitStrategy   resilience.RateLimitStrategy
}

type CircuitBreakerProps struct {
	CircuitBreakerExpression string
	CheckPeriod              string
	FallbackDuration         string
	RecoveryDuration         string
}

func NewRateLimitMiddleware(scope constructs.Construct, id string, ns string, appLabel string, props *RateLimitProps) traefikio.Middleware {
	if props == nil {
		return nil
	}
	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRateLimit)
	var spec traefikio.MiddlewareSpecRateLimit
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
			RateLimit: &spec,
		},
	})
}

func NewCircuitBreakerMiddleware(scope constructs.Construct, id string, ns string, appLabel string, props *CircuitBreakerProps) traefikio.Middleware {
	if props == nil {
		return nil
	}

	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeCircuitBreaker)
	var spec traefikio.MiddlewareSpecCircuitBreaker
	if props.CircuitBreakerExpression != "" {
		spec.Expression = jsii.String(props.CircuitBreakerExpression)
	}

	if props.CheckPeriod != "" {
		spec.CheckPeriod = traefikio.MiddlewareSpecCircuitBreakerCheckPeriod_FromString(jsii.String(props.CheckPeriod))
	}

	if props.FallbackDuration != "" {
		spec.FallbackDuration = traefikio.MiddlewareSpecCircuitBreakerFallbackDuration_FromString(jsii.String(props.FallbackDuration))
	}

	if props.RecoveryDuration != "" {
		spec.RecoveryDuration = traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration_FromString(jsii.String(props.RecoveryDuration))
	}

	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			CircuitBreaker: &spec,
		},
	})
}

func NewRetryMiddleware(scope constructs.Construct, id string, ns string, appLabel string, props *RetryProps) traefikio.Middleware {
	if props == nil {
		return nil
	}

	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRetry)
	var spec traefikio.MiddlewareSpecRetry

	if props.RetryAttempts != 0 {
		spec.Attempts = jsii.Number(props.RetryAttempts)
	}

	if props.IntervalMs != "" {
		spec.InitialInterval = traefikio.MiddlewareSpecRetryInitialInterval_FromString(jsii.String(props.IntervalMs))
	}

	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			Retry: &spec,
		},
	})

}

func middlewareMetadata(ns string, appLabel string, mwwType string) *cdk8s.ApiObjectMetadata {
	name := fmt.Sprintf("%s-%s", appLabel, mwwType)
	labels := middlewareLabels()
	annotations := middlewareAnnotations()
	return &cdk8s.ApiObjectMetadata{
		Annotations: annotations,
		Labels:      labels,
		Name:        jsii.String(name),
		Namespace:   jsii.String(ns),
	}
}

func middlewareLabels() *map[string]*string {
	return nil
}

func middlewareAnnotations() *map[string]*string {
	return nil
}
