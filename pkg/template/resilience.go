package template

import (
	"errors"
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

type ResilienceTemplate struct {
	RetryPolicy          *resilience.RetryConfig          `yaml:"retry"`
	RateLimitPolicy      *resilience.RateLimitConfig      `yaml:"rateLimit"`
	CircuitBreakerPolicy *resilience.CircuitBreakerConfig `yaml:"circuitBreaker"`
	Scaling              *ScalingTemplate
}

func (r *ResilienceTemplate) Validate() error {
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
