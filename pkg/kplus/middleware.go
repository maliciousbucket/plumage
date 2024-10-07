package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/middleware"
	"strings"
)

func NewCircuitBreaker(scope constructs.Construct, id string, ns string, service *ServiceTemplate) traefikio.Middleware {
	if service.CircuitBreaker == nil {
		return nil
	}

	props := &middleware.CircuitBreakerProps{
		CircuitBreakerExpression: service.CircuitBreaker.CircuitBreakerExpression(),
		CheckPeriod:              service.CircuitBreaker.CheckPeriod(),
		FallbackDuration:         service.CircuitBreaker.FallbackDuration(),
		RecoveryDuration:         service.CircuitBreaker.RecoveryDuration(),
	}

	circuitBreaker := middleware.NewCircuitBreakerMiddleware(scope, id, ns, service.Name, props)
	return circuitBreaker
}

func NewRetry(scope constructs.Construct, id string, ns string, service *ServiceTemplate) traefikio.Middleware {
	if service.Retry == nil {
		return nil
	}

	props := &middleware.RetryProps{
		RetryAttempts: service.Retry.Attempts,
		IntervalMs:    service.Retry.IntervalMS(),
	}

	retry := middleware.NewRetryMiddleware(scope, id, ns, service.Name, props)
	return retry
}

func NewRateLimit(scope constructs.Construct, id string, ns string, service *ServiceTemplate) traefikio.Middleware {
	if service.RateLimit == nil {
		return nil
	}

	props := &middleware.RateLimitProps{
		AverageRequests: service.RateLimit.AverageRequests(),
		BurstRequests:   service.RateLimit.BurstRequests(),
		RatePeriod:      service.RateLimit.RatePeriod(),
		LimitStrategy:   service.RateLimit.LimitStrategy(),
	}
	rateLimit := middleware.NewRateLimitMiddleware(scope, id, ns, service.Name, props)
	return rateLimit
}

func NewDefaultMiddlewares(scope constructs.Construct, ns, appLabel string, defaults []string) []string {
	if len(defaults) == 0 {
		return nil
	}

	result := []string{}
	for _, mw := range defaults {
		mwType := strings.ToLower(mw)
		switch mwType {
		case "retry":
			name := fmt.Sprintf("%s-retry", appLabel)
			middleware.NewDefaultRetryMiddleware(scope, name, ns, appLabel)
			result = append(result, name)
		case "ratelimit":
			name := fmt.Sprintf("%s-ratelimit", appLabel)
			middleware.NewDefaultRateLimitMiddleware(scope, name, ns, appLabel)
			result = append(result, name)
		case "circuitbreaker":
			name := fmt.Sprintf("%s-circuitbreaker", appLabel)
			middleware.NewDefaultCircuitBreakerMiddleware(scope, name, ns, appLabel)
			result = append(result, name)
		}
	}
	return result
}
