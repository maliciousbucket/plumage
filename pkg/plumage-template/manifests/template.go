package manifests

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/pkg/config"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/autoscaling"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/ingress"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/middleware"
	"github.com/maliciousbucket/plumage/pkg/resilience"
)

type Template struct {
	Host string
}

type ServiceTemplate struct {
	Name           string         `yaml:"name"`
	Namespace      string         `yaml:"namespace"`
	Host           string         `yaml:"host"`
	Paths          []*ServicePath `yaml:"paths"`
	LoadBalancer   bool           `yaml:"loadBalancer,omitempty"`
	Middlewares    []string       `yaml:"middlewares,omitempty"`
	middlewareRefs []*string
	Retry          *resilience.RetryConfig          `yaml:"retry,omitempty"`
	CircuitBreaker *resilience.CircuitBreakerConfig `yaml:"circuitBreaker,omitempty"`
	RateLimit      *resilience.RateLimitConfig      `yaml:"rateLimit,omitempty"`
	Scaling        *plumagetemplate.ScalingConfig   `yaml:"scaling,omitempty"`
	MonitoringEnv  []string                         `yaml:"monitoringEnv,omitempty"`

	MonitoringAliases map[string]string `yaml:"aliases,omitempty"`
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

func NewService(scope constructs.Construct, id string, s *ServiceTemplate, target cdk8s.ApiObject, m *config.CollectorConfig) constructs.Construct {
	sc := constructs.NewConstruct(scope, jsii.String(id))
	if len(s.MonitoringEnv) > 0 || len(s.MonitoringAliases) > 0 {
		WithMonitoringEnv(s, target, m)(sc)
		fmt.Printf("Service: %s API: %s", s.Name, *target.Name())
	}

	if len(s.Middlewares) > 0 {
		if s.RateLimit != nil {
			WithRateLimit(s)(sc)
		}
		if s.CircuitBreaker != nil {
			WithCircuitBreaker(s)(sc)
		}
		if s.Retry != nil {
			WithRetry(s)(sc)
		}
	}

	WithIngressRoute(s, s.middlewareRefs)(sc)

	if s.Scaling != nil {
		WithAutoScaling(s)(sc)
	}
	return sc
}

type SynthFunc func(scope constructs.Construct) constructs.Construct

func (s *ServiceTemplate) IngressRouteProps() *ingress.RouteProps {
	var paths []*ingress.ServicePaths
	for _, path := range s.Paths {
		paths = append(paths, &ingress.ServicePaths{
			Path: path.Prefix,
			Port: path.Port,
		})
	}
	return &ingress.RouteProps{
		Name:      s.Name,
		Namespace: s.Namespace,
		Config: &ingress.RouteConfig{
			Host:               s.Host,
			Paths:              paths,
			EnableLoadBalancer: s.LoadBalancer,
		},
	}
}

func (s *ServiceTemplate) retryProps() *middleware.RetryProps {
	if s.Retry == nil {
		return nil
	}
	return &middleware.RetryProps{
		RetryAttempts: s.Retry.RetryAttempts(),
		IntervalMs:    s.Retry.IntervalMS(),
	}
}

func (s *ServiceTemplate) circuitBreakerProps() *middleware.CircuitBreakerProps {
	if s.CircuitBreaker == nil {
		return nil
	}
	return &middleware.CircuitBreakerProps{
		CircuitBreakerExpression: s.CircuitBreaker.CircuitBreakerExpression(),
		CheckPeriod:              s.CircuitBreaker.CheckPeriod(),
		FallbackDuration:         s.CircuitBreaker.FallbackDuration(),
		RecoveryDuration:         s.CircuitBreaker.RecoveryDuration(),
	}
}

func (s *ServiceTemplate) rateLimitProps() *middleware.RateLimitProps {
	if s.RateLimit == nil {
		return nil
	}
	return &middleware.RateLimitProps{
		AverageRequests: s.RateLimit.AverageRequests(),
		BurstRequests:   s.RateLimit.BurstRequests(),
		RatePeriod:      s.RateLimit.RatePeriod(),
		LimitStrategy:   s.RateLimit.LimitStrategy(),
	}
}

func (s *ServiceTemplate) autoscalingProps() *autoscaling.AutoScalerProps {
	if s.Scaling == nil {
		return nil
	}
	return &autoscaling.AutoScalerProps{
		Name:      s.Name,
		Namespace: s.Namespace,
		Scaling:   s.Scaling,
	}
}

func WithIngressRoute(s *ServiceTemplate, middlewareRefs []*string) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		props := s.IngressRouteProps()
		id := fmt.Sprintf("%s-ingressroute", s.Name)
		return ingress.NewServiceIngressRoute(scope, id, props, middlewareRefs)
	}
}

func WithRetry(s *ServiceTemplate) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		props := s.retryProps()
		id := fmt.Sprintf("%s-retry", s.Name)
		if props == nil {
			return nil
		}
		retry := middleware.NewRetryMiddleware(scope, id, s.Namespace, s.Name, props)
		s.middlewareRefs = append(s.middlewareRefs, retry.Name())
		return retry
	}
}

func WithCircuitBreaker(s *ServiceTemplate) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		props := s.circuitBreakerProps()
		if props == nil {
			return nil
		}
		id := fmt.Sprintf("%s-%s", s.Name, "circuitbreaker")
		circuitBreaker := middleware.NewCircuitBreakerMiddleware(scope, id, s.Namespace, s.Name, props)
		s.middlewareRefs = append(s.middlewareRefs, circuitBreaker.Name())
		return circuitBreaker
	}
}

func WithRateLimit(s *ServiceTemplate) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		props := s.rateLimitProps()
		if props == nil {
			return nil
		}
		id := fmt.Sprintf("%s-ratelimit", s.Name)
		rateLimit := middleware.NewRateLimitMiddleware(scope, id, s.Namespace, s.Name, props)
		s.middlewareRefs = append(s.middlewareRefs, rateLimit.Name())
		return rateLimit
	}
}

func WithAutoScaling(s *ServiceTemplate) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		props := s.autoscalingProps()
		if props == nil {
			return nil
		}
		id := fmt.Sprintf("%s-autoscaler", s.Name)
		return autoscaling.NewHorizontalAutoscaler(scope, id, props)
	}
}

func WithMonitoringEnv(s *ServiceTemplate, target cdk8s.ApiObject, m *config.CollectorConfig) SynthFunc {
	return func(scope constructs.Construct) constructs.Construct {
		return AddServiceEnvironmentVariables(scope, s, target, m)
	}
}
