package template

type ResiliencePolicy struct {
}

type RateLimitPolicy struct {
}

type CircuitBreakerPolicy struct {
}

type RetryPolicy struct {
	Attempts          int `json:"attempts,omitempty" yaml:"attempts"`
	InitialIntervalMS int `json:"initialIntervalMs,omitempty" yaml:"initialIntervalMs"`
}
