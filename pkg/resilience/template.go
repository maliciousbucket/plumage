package resilience

type ResTemplate struct {
	ScalingPolicy        ScalingConfig        `json:"scalingPolicy,omitempty"`
	RetryPolicy          RetryConfig          `yaml:"retry,omitempty"`
	RateLimitPolicy      RateLimitConfig      `yaml:"rateLimit,omitempty"`
	CircuitBreakerPolicy CircuitBreakerConfig `yaml:"circuitBreaker,omitempty"`
}

type ScalingConfig struct {
	MinReplicas             int
	MaxReplicas             int
	TargetReplicas          int
	TargetCpuUtilization    float64
	TargetCpuAmount         float64
	TargetMemoryUtilization float64
	TargetMemoryAmount      float64
}
