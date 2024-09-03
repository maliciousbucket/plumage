package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/template"
)

const (
	ClusterIpType   = "ClusterIP"
	StatefulSetKind = "StatefulSet"
	ServiceKind     = "Service"
)

func NewServiceChart() cdk8s.Chart {
	return nil
}

func NewService() constructs.Construct {
	return nil
}

func newServiceDeployment() constructs.Construct {
	return nil
}

func newServiceIngressRoute(scope constructs.Construct, id string) traefikio.IngressRoute {
	return nil

}

func newServiceMiddleware(scope constructs.Construct, s *template.ServiceTemplate) constructs.Construct {
	if s.Resilience.RetryPolicy != nil {
		NewRetryMiddleware(scope, s.Name, s.Resilience.RetryPolicy)
	} else {
		EmptyMiddleware(scope, s.Name, MiddlewareNameRetry)
	}

	if s.Resilience.CircuitBreakerPolicy != nil {
		NewCircuitBreakerMiddleware(scope, s.Name, s.Resilience.CircuitBreakerPolicy)
	} else {
		EmptyMiddleware(scope, s.Name, MiddlewareNameCircuitBreaker)
	}

	if s.Resilience.RateLimitPolicy != nil {
		NewRateLimitMiddleware(scope, s.Name, s.Resilience.RateLimitPolicy)
	}

	return scope
}
