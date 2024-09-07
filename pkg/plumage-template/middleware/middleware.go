package middleware

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

const (
	MiddlewareTypeRetry          = "retry"
	MiddlewareTypeRateLimit      = "ratelimit"
	MiddlewareTypeCircuitBreaker = "circuitbreaker"
)

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
