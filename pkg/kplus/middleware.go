package kplus

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/middleware"
)

func NewCircuitBreaker(scope constructs.Construct, id string, service *ServiceTemplate) traefikio.Middleware {
	if service.CircuitBreaker == nil {
		return nil
	}

	props := &middleware.CircuitBreakerProps{
		CircuitBreakerExpression: service.CircuitBreaker.CircuitBreakerExpression(),
		CheckPeriod:              service.CircuitBreaker.CheckPeriod(),
		FallbackDuration:         service.CircuitBreaker.FallbackDuration(),
		RecoveryDuration:         service.CircuitBreaker.RecoveryDuration(),
	}

	circuitBreaker := middleware.NewCircuitBreakerMiddleware(scope, id, service.Namespace, service.Name, props)
	return circuitBreaker
}

func NewRetry(scope constructs.Construct, id string, service *ServiceTemplate) traefikio.Middleware {
	if service.Retry == nil {
		return nil
	}

	props := &middleware.RetryProps{
		RetryAttempts: service.Retry.Attempts,
		IntervalMs:    service.Retry.IntervalMS(),
	}

	retry := middleware.NewRetryMiddleware(scope, id, service.Namespace, service.Name, props)
	return retry
}

func NewRateLimit(scope constructs.Construct, id string, service *ServiceTemplate) traefikio.Middleware {
	if service.RateLimit == nil {
		return nil
	}

	props := &middleware.RateLimitProps{
		AverageRequests: service.RateLimit.AverageRequests(),
		BurstRequests:   service.RateLimit.BurstRequests(),
		RatePeriod:      service.RateLimit.RatePeriod(),
		LimitStrategy:   service.RateLimit.LimitStrategy(),
	}
	rateLimit := middleware.NewRateLimitMiddleware(scope, id, service.Namespace, service.Name, props)
	return rateLimit
}
