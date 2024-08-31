package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/template"
)

const (
	MiddlewareKind = "Middleware"

	MiddlewareNameRetry          = MiddlewareName("Retry")
	MiddlewareNameCircuitBreaker = MiddlewareName("CircuitBreaker")
	MiddlewareNameRateLimiting   = MiddlewareName("RateLimit")
)

type MiddlewareName string

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

type MiddlewareProps struct {
}

func defaultMiddlewareMetadata(svcName string, variant MiddlewareName) *cdk8s.ApiObjectMetadata {
	nameSpace := TestbedNamespace
	name := fmt.Sprintf("%s-%s", svcName, variant)

	return &cdk8s.ApiObjectMetadata{
		Namespace:   &nameSpace,
		Annotations: defaultMiddlewareAnnotations(name),
		Labels:      defaultMiddlewareLabels(name),
		Name:        &name,
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
	IntervalMS() int
}

func RetryMiddlewareSpec(spec RetrySpec) *traefikio.MiddlewareSpecRetry {

	return &traefikio.MiddlewareSpecRetry{
		Attempts:        jsii.Number(spec.RetryAttempts()),
		InitialInterval: traefikio.MiddlewareSpecRetryInitialInterval_FromNumber(jsii.Number(spec.IntervalMS())),
	}
}

func NewRetryMiddlewareProps(svcName string, spec RetrySpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameRetry)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{Retry: RetryMiddlewareSpec(spec)},
	}
}

// RateLimitSpec TODO: Move strategy efinition to types
type RateLimitSpec interface {
	Average() int
	Burst() int
	Period() string
	Strategy() template.RateLimitStrategy
}

func RateLimitMiddlewareSpec(spec RateLimitSpec) *traefikio.MiddlewareSpecRateLimit {
	var middleWareSpec traefikio.MiddlewareSpecRateLimit
	if spec.Average() > 0 {
		middleWareSpec.Average = jsii.Number(spec.Average())
	}

	if spec.Burst() > 0 {
		middleWareSpec.Burst = jsii.Number(spec.Burst())
	}

	if spec.Period() != "" {
		period := spec.Period()
		middleWareSpec.Period = traefikio.MiddlewareSpecRateLimitPeriod_FromString(&period)
	}

	if spec.Strategy() != nil {
		strategy := spec.Strategy()
		switch s := strategy.(type) {
		case *template.IpDepthStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				IpStrategy: &traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy{
					Depth: jsii.Number(s.Depth),
				},
			}
		case *template.RequestHeaderStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHeaderName: jsii.String(s.HeaderName),
			}
		case *template.HostStrategy:
			middleWareSpec.SourceCriterion = &traefikio.MiddlewareSpecRateLimitSourceCriterion{
				RequestHost: jsii.Bool(s.Host),
			}
		}
	}
	return &middleWareSpec
}

func NewRateLimitMiddlewareProps(svcName string, spec RateLimitSpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameRateLimiting)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{RateLimit: RateLimitMiddlewareSpec(spec)},
	}
}

type CircuitBreakerSpec interface {
	CircuitBreakerExpression() string
	CheckPeriod() string
	FallbackDuration() string
	RecoveryDuration() string
}

func CircuitBreakerMiddlewareSpec(spec CircuitBreakerSpec) *traefikio.MiddlewareSpecCircuitBreaker {
	return &traefikio.MiddlewareSpecCircuitBreaker{
		Expression:       jsii.String(spec.CircuitBreakerExpression()),
		FallbackDuration: traefikio.MiddlewareSpecCircuitBreakerFallbackDuration_FromString(jsii.String(spec.FallbackDuration())),
		RecoveryDuration: traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration_FromString(jsii.String(spec.RecoveryDuration())),
		CheckPeriod:      traefikio.MiddlewareSpecCircuitBreakerCheckPeriod_FromString(jsii.String(spec.CheckPeriod())),
	}
}

func NewCircuitBreakerMiddleProps(svcName string, spec CircuitBreakerSpec) *traefikio.MiddlewareProps {
	metaData := defaultMiddlewareMetadata(svcName, MiddlewareNameCircuitBreaker)
	return &traefikio.MiddlewareProps{
		Metadata: metaData,
		Spec:     &traefikio.MiddlewareSpec{CircuitBreaker: CircuitBreakerMiddlewareSpec(spec)},
	}
}
