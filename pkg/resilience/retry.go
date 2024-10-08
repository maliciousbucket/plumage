package resilience

import "fmt"

//TODO: update / change policy if time permits

type RetryConfig struct {
	Attempts          int `json:"attempts,omitempty" yaml:"attempts,omitempty"`
	InitialIntervalMS int `json:"initialIntervalMs,omitempty" yaml:"initialIntervalMs,omitempty"`
}

func (r *RetryConfig) Validate() error {
	if r.Attempts < 1 {
		return fmt.Errorf("invalid retry attempts: %d, must be greater than 1", r.Attempts)
	}
	if r.InitialIntervalMS < 0 {
		return fmt.Errorf("invalid intervalMS, must be greater than zero: %d", r.InitialIntervalMS)
	}
	return nil
}

func (r *RetryConfig) RetryAttempts() int {
	return r.Attempts
}

func (r *RetryConfig) IntervalMS() string {
	return fmt.Sprintf("%dms", r.InitialIntervalMS)
}
