package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

const (
	MiddlewareKind = "Middleware"

	MiddlewareNameRetry          = MiddlewareName("retry")
	MiddlewareNameCircuitBreaker = MiddlewareName("circuitBreaker")
	MiddlewareNameRateLimiting   = MiddlewareName("rateLimit")
)

type MiddlewareName string

//func NewMiddleware(scope constructs.Construct, id string) cdk8s.Chart {
//
//	middleware := traefikio.MiddlewareProps{
//		Metadata: &cdk8s.ApiObjectMetadata{
//			Annotations:     nil,
//			Finalizers:      nil,
//			Labels:          nil,
//			Name:            nil,
//			Namespace:       nil,
//			OwnerReferences: nil,
//		},
//		Spec: &traefikio.MiddlewareSpec{
//			AddPrefix:         nil,
//			BasicAuth:         nil,
//			Buffering:         nil,
//			Chain:             nil,
//			CircuitBreaker:    nil,
//			Compress:          nil,
//			ContentType:       nil,
//			DigestAuth:        nil,
//			Errors:            nil,
//			ForwardAuth:       nil,
//			GrpcWeb:           nil,
//			Headers:           nil,
//			InFlightReq:       nil,
//			IpAllowList:       nil,
//			IpWhiteList:       nil,
//			PassTlsClientCert: nil,
//			Plugin:            nil,
//			RateLimit:         nil,
//			RedirectRegex:     nil,
//			RedirectScheme:    nil,
//			ReplacePath:       nil,
//			ReplacePathRegex:  nil,
//			Retry:             nil,
//			StripPrefix:       nil,
//			StripPrefixRegex:  nil,
//		},
//	}
//	return nil
//}

type MiddlewareProps struct {
}

func defaultMiddlewareMetadata(svcName string, variant MiddlewareName) *cdk8s.ApiObjectMetadata {
	nameSpace := TestbedNamespace
	name := fmt.Sprintf("%s-%s", svcName, variant)

	return &cdk8s.ApiObjectMetadata{
		Namespace:   &nameSpace,
		Annotations: defaultMiddlewareAnnotations(name),
		Labels:      defaultMiddlewareLabels(name),
		Name:        jsii.String(svcName),
	}
}

func defaultMiddlewareAnnotations(name string) *map[string]*string {
	annotations := map[string]string{
		KubernetesNameLabel: name,
	}
	return MapToAnnotations(annotations)
}

func defaultMiddlewareLabels(name string) *map[string]*string {
	labels := map[string]string{
		KubernetesNameLabel: name,
	}
	return MapToLabels(labels)
}

type RetrySpec interface {
	RetryAttempts() int
	IntervalMS() string
}

func retryMiddlewareSpec(spec RetrySpec) *traefikio.MiddlewareSpecRetry {

	var retrySpec traefikio.MiddlewareSpecRetry

	if spec.RetryAttempts() > 0 {
		retrySpec.Attempts = jsii.Number(spec.RetryAttempts())
	}

	if spec.IntervalMS() == "" {
		retrySpec.InitialInterval = traefikio.MiddlewareSpecRetryInitialInterval_FromString(jsii.String(spec.IntervalMS()))
	}
	return &retrySpec

}

func newRetryMiddlewareProps(svcName string, spec RetrySpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameRetry)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{Retry: retryMiddlewareSpec(spec)},
	}
}

func NewRetryMiddleware(scope constructs.Construct, svcName string, spec RetrySpec) traefikio.Middleware {
	props := newRetryMiddlewareProps(svcName, spec)
	name := fmt.Sprintf("%s-retry", svcName)
	middleware := traefikio.NewMiddleware(scope, jsii.String(name), props)
	return middleware
}

// RateLimitSpec TODO: Move strategy efinition to types
type RateLimitSpec interface {
	AverageRequests() int
	BurstRequests() int
	RatePeriod() string
	LimitStrategy() resilience.RateLimitStrategy
}

func rateLimitMiddlewareSpec(spec RateLimitSpec) *traefikio.MiddlewareSpecRateLimit {
	var middleWareSpec traefikio.MiddlewareSpecRateLimit
	if spec.AverageRequests() > 0 {
		middleWareSpec.Average = jsii.Number(spec.AverageRequests())
	}

	if spec.BurstRequests() > 0 {
		middleWareSpec.Burst = jsii.Number(spec.BurstRequests())
	}

	if spec.RatePeriod() != "" {
		period := spec.RatePeriod()
		middleWareSpec.Period = traefikio.MiddlewareSpecRateLimitPeriod_FromString(&period)
	}

	if spec.LimitStrategy() != nil {
		strategy := spec.LimitStrategy()
		switch s := strategy.(type) {
		case *resilience.IpDepthStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				IpStrategy: &traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy{
					Depth: jsii.Number(s.Depth),
				},
			}
		case *resilience.RequestHeaderStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHeaderName: jsii.String(s.HeaderName),
			}
		case *resilience.HostStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHost: jsii.Bool(s.Host),
			}
		}
	}
	return &middleWareSpec
}

func newRateLimitMiddlewareProps(svcName string, spec RateLimitSpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameRateLimiting)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{RateLimit: rateLimitMiddlewareSpec(spec)},
	}
}

func NewRateLimitMiddleware(scope constructs.Construct, svcName string, spec RateLimitSpec) traefikio.Middleware {
	props := newRateLimitMiddlewareProps(svcName, spec)
	name := fmt.Sprintf("%s-ratelimit", svcName)
	middleware := traefikio.NewMiddleware(scope, jsii.String(name), props)
	return middleware
}

type CircuitBreakerSpec interface {
	CircuitBreakerExpression() string
	CheckPeriod() string
	FallbackDuration() string
	RecoveryDuration() string
}

func circuitBreakerMiddlewareSpec(spec CircuitBreakerSpec) *traefikio.MiddlewareSpecCircuitBreaker {
	var circuitBreakerSpec traefikio.MiddlewareSpecCircuitBreaker

	if spec.CircuitBreakerExpression() != "" {
		circuitBreakerSpec.Expression = jsii.String(spec.CircuitBreakerExpression())
	}

	if spec.FallbackDuration() != "" {
		circuitBreakerSpec.FallbackDuration = traefikio.MiddlewareSpecCircuitBreakerFallbackDuration_FromString(jsii.String(spec.FallbackDuration()))
	}

	if spec.RecoveryDuration() != "" {
		circuitBreakerSpec.RecoveryDuration = traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration_FromString(jsii.String(spec.RecoveryDuration()))
	}

	if spec.CheckPeriod() != "" {
		circuitBreakerSpec.CheckPeriod = traefikio.MiddlewareSpecCircuitBreakerCheckPeriod_FromString(jsii.String(spec.CheckPeriod()))
	}
	return &circuitBreakerSpec

	//return &traefikio.MiddlewareSpecCircuitBreaker{
	//	Expression:       jsii.String(spec.CircuitBreakerExpression()),
	//	FallbackDuration: traefikio.MiddlewareSpecCircuitBreakerFallbackDuration_FromString(jsii.String(spec.FallbackDuration())),
	//	RecoveryDuration: traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration_FromString(jsii.String(spec.RecoveryDuration())),
	//	CheckPeriod:      traefikio.MiddlewareSpecCircuitBreakerCheckPeriod_FromString(jsii.String(spec.CheckPeriod())),
	//}
}

func newCircuitBreakerMiddleProps(svcName string, spec CircuitBreakerSpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameCircuitBreaker)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{CircuitBreaker: circuitBreakerMiddlewareSpec(spec)},
	}
}

func NewCircuitBreakerMiddleware(scope constructs.Construct, svcName string, spec CircuitBreakerSpec) traefikio.Middleware {

	props := newCircuitBreakerMiddleProps(svcName, spec)
	name := fmt.Sprintf("%s-circuitbreaker", svcName)

	middleware := traefikio.NewMiddleware(scope, jsii.String(name), props)

	return middleware
}

func EmptyMiddleware(scope constructs.Construct, svcName string, mwType MiddlewareName) traefikio.Middleware {
	props := emptyMiddlewareProps(svcName, mwType)
	name := fmt.Sprintf("%s-%s", svcName, mwType)
	return traefikio.NewMiddleware(scope, jsii.String(name), props)
}

func emptyMiddlewareProps(svcName string, mwType MiddlewareName) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, mwType)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
	}
}
