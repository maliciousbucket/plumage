package middleware

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

const (
	defaultRetryAttempts = 5
	defaultRetryDelay    = "100ms"
)

type RetryProps struct {
	RetryAttempts int
	IntervalMs    string
}

func NewRetryMiddleware(scope constructs.Construct, id string, ns string, appLabel string, props *RetryProps) traefikio.Middleware {
	if props == nil {
		return nil
	}

	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRetry)
	spec := &traefikio.MiddlewareSpecRetry{}

	if props.RetryAttempts != 0 {
		spec.Attempts = jsii.Number(props.RetryAttempts)
	}

	if props.IntervalMs != "" {
		spec.InitialInterval = traefikio.MiddlewareSpecRetryInitialInterval_FromString(jsii.String(props.IntervalMs))
	}

	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			Retry: spec,
		},
	})

}

func NewDefaultRetryMiddleware(scope constructs.Construct, id string, ns string, appLabel string) traefikio.Middleware {
	metadata := middlewareMetadata(ns, appLabel, MiddlewareTypeRetry)
	return traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: metadata,
		Spec: &traefikio.MiddlewareSpec{
			Retry: &traefikio.MiddlewareSpecRetry{
				Attempts:        jsii.Number(defaultRetryAttempts),
				InitialInterval: traefikio.MiddlewareSpecRetryInitialInterval_FromString(jsii.String(defaultRetryDelay)),
			},
		},
	})
}
