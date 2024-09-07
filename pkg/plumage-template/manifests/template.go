package manifests

import (
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

type Template struct {
}

type ServiceTemplate struct {
	Name              string                           `yaml:"name"`
	Namespace         string                           `yaml:"namespace"`
	Host              string                           `yaml:"host"`
	Paths             []*ServicePath                   `yaml:"paths"`
	LoadBalancer      bool                             `yaml:"loadBalancer,omitempty"`
	Retry             *resilience.RetryConfig          `yaml:"retry,omitempty"`
	CircuitBreaker    *resilience.CircuitBreakerConfig `yaml:"circuitBreaker,omitempty"`
	RateLimit         *resilience.RateLimitConfig      `yaml:"rateLimit,omitempty"`
	MonitoringEnv     []string                         `yaml:"monitoringEnv,omitempty"`
	MonitoringAliases map[string]string                `yaml:"aliases,omitempty"`
}

type ServicePath struct {
	Host   string
	Prefix string
	Port   int
}

type GlobalTemplate struct {
	Host         string
	Namespace    string
	LoadBalancer bool
}
