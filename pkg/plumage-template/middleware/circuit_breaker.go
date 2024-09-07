package middleware

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

type CircuitBreakerProps struct {
	CircuitBreakerExpression string
	CheckPeriod              string
	FallbackDuration         string
	RecoveryDuration         string
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
