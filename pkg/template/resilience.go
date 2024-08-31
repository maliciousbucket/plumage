package template

import (
	"errors"
	"fmt"
)

type ResiliencePolicy struct {
	RetryPolicy          *RetryConfig          `yaml:"retryPolicy"`
	RateLimitPolicy      *RateLimitConfig      `yaml:"rateLimitPolicy"`
	CircuitBreakerPolicy *CircuitBreakerConfig `yaml:"circuitBreakerPolicy"`
}

func (r *ResiliencePolicy) Validate() error {
	var resErr error
	if r.RetryPolicy == nil {
		err := r.RetryPolicy.Validate()
		if err != nil {
			resErr = errors.Join(resErr, err)
		}
	}

	if r.RateLimitPolicy == nil {
		err := r.RateLimitPolicy.Validate()
		if err != nil {
			resErr = errors.Join(resErr, err)
		}
	}

	if r.CircuitBreakerPolicy == nil {
		err := r.CircuitBreakerPolicy.Validate()
		if err != nil {
			resErr = errors.Join(resErr, err)
		}
	}
	if resErr != nil {
		resErr = fmt.Errorf("invalid resilience policy: %w", resErr)
	}
	return resErr
}
