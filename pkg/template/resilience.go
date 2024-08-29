package template

type ResiliencePolicy struct {
}

type RateLimitPolicy struct {
}

type CircuitBreakerPolicy struct {
}

type RetryPolicy struct {
	Attempts          int `json:"attempts", yaml:"attempts"`
	InitialIntervalMS int `json:"initialIntervalMs", yaml:"initialIntervalMs"`
}
