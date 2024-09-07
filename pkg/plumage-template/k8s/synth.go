package k8s

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/ingress"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/middleware"
	"github.com/maliciousbucket/plumage/pkg/resilience"
	"github.com/maliciousbucket/plumage/pkg/types"
	"strings"
)

const (
	PromScrapePortAnnotation   = "prometheus.io/port"
	PromScrapePathAnnotation   = "prometheus.io/path"
	PrometheusScrapeAnnotation = "prometheus.io/scrape"
	AppLabel                   = "app.kubernetes.io/name"
	AppsV1                     = "apps/v1"
	ServiceKind                = "Service"
	DeploymentKind             = "Deployment"
)

type WebServiceProps struct {
	Name           string
	Namespace      string
	Image          string
	Commands       []string
	Args           []string
	StartupProbe   *types.CommandProbe
	HealthCheck    *plumagetemplate.HttpProbe
	Ports          []*plumagetemplate.ServicePort
	Resources      *plumagetemplate.ServiceResources
	Monitoring     *plumagetemplate.MonitoringConfig
	InitContainers []*plumagetemplate.InitContainer
	Scaling        *plumagetemplate.ScalingConfig
	Resilience     *resilience.ResTemplate
	Paths          []ingress.ServicePaths
	Env            map[string]string
	Middlewares    []string
	Ingress        *ingress.RouteConfig
}

type SynthOpts struct {
	Options []SynthFunc
}

func newWebService(scope constructs.Construct, id string, props *WebServiceProps) constructs.Construct {
	return nil
}

type SynthFunc func(scope constructs.Construct, p *WebServiceProps) constructs.Construct

func WithDeployment() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.deploymentProps()
		id := fmt.Sprintf("%s-%s", p.Name, "deployment")
		return NewDeployment(scope, id, p.Namespace, p.Name, props)
	}
}

func WithService() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.serviceProps()
		id := fmt.Sprintf("%s-%s", p.Name, "service")
		return NewService(scope, id, props)
	}
}

func WithAutoScaling() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.autoScalingProps()
		id := fmt.Sprintf("%s-%s", p.Name, "autoscaler")
		return NewHorizontalAutoscaler(scope, id, props)
	}
}

func WithDefaultAutoScaling() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		id := fmt.Sprintf("%s-%s", p.Name, "autoscaler")
		return DefaultAutoScaler(scope, id, p.Namespace, p.Name)
	}
}

func WithIngressRoute() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.ingressRouteProps()
		id := fmt.Sprintf("%s-%s", p.Name, "ingressroute")
		return ingress.NewIngressRoute(scope, id, props)
	}
}

func WithRetry() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.retryProps()
		id := fmt.Sprintf("%s-%s", p.Name, "retry")
		return middleware.NewRetryMiddleware(scope, id, p.Namespace, p.Name, props)
	}
}

func WithCircuitBreaker() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.circuitBreakerProps()
		id := fmt.Sprintf("%s-%s", p.Name, "circuitbreaker")
		return middleware.NewCircuitBreakerMiddleware(scope, id, p.Namespace, p.Name, props)
	}
}

func WithRateLimit() SynthFunc {
	return func(scope constructs.Construct, p *WebServiceProps) constructs.Construct {
		props := p.rateLimitProps()
		id := fmt.Sprintf("%s-%s", p.Name, "ratelimit")
		return middleware.NewRateLimitMiddleware(scope, id, p.Namespace, p.Name, props)
	}
}

func (p *WebServiceProps) deploymentProps() *DeploymentProps {
	minReplicas := 1
	if p.Scaling != nil {
		if p.Scaling.MinReplicas != 0 {
			minReplicas = int(p.Scaling.MinReplicas)
		}
	}
	return &DeploymentProps{
		Name:           p.Name,
		Image:          p.Image,
		Commands:       p.Commands,
		Args:           p.Args,
		StartupProbe:   p.StartupProbe,
		HealthCheck:    p.HealthCheck,
		Ports:          p.Ports,
		Resources:      p.Resources,
		Monitoring:     p.Monitoring,
		InitContainers: p.InitContainers,
		MinReplicas:    minReplicas,
		Env:            p.Env,
	}
}

func (p *WebServiceProps) serviceProps() *ServiceProps {
	return &ServiceProps{
		Name:       p.Name,
		Namespace:  p.Namespace,
		Ports:      p.Ports,
		Monitoring: p.Monitoring,
	}
}

func (p *WebServiceProps) autoScalingProps() *AutoScalerProps {
	return &AutoScalerProps{
		Name:      p.Name,
		Namespace: p.Namespace,
		Scaling:   p.Scaling,
	}
}

func (p *WebServiceProps) retryProps() *middleware.RetryProps {
	if p.Resilience == nil {
		return nil
	}

	return &middleware.RetryProps{
		RetryAttempts: p.Resilience.RetryPolicy.RetryAttempts(),
		IntervalMs:    p.Resilience.RetryPolicy.IntervalMS(),
	}
}

func (p *WebServiceProps) rateLimitProps() *middleware.RateLimitProps {
	if p.Resilience == nil {
		return nil
	}
	return &middleware.RateLimitProps{
		AverageRequests: p.Resilience.RateLimitPolicy.Average,
		BurstRequests:   p.Resilience.RateLimitPolicy.Burst,
		RatePeriod:      p.Resilience.RateLimitPolicy.Period,
		LimitStrategy:   p.Resilience.RateLimitPolicy.LimitStrategy(),
	}
}

func (p *WebServiceProps) circuitBreakerProps() *middleware.CircuitBreakerProps {
	if p.Resilience == nil {
		return nil
	}
	return &middleware.CircuitBreakerProps{
		CircuitBreakerExpression: p.Resilience.CircuitBreakerPolicy.CircuitBreakerExpression(),
		CheckPeriod:              p.Resilience.CircuitBreakerPolicy.CheckPeriod(),
		FallbackDuration:         p.Resilience.CircuitBreakerPolicy.FallbackDuration(),
		RecoveryDuration:         p.Resilience.CircuitBreakerPolicy.RecoveryDuration(),
	}
}

func (p *WebServiceProps) ingressRouteProps() *ingress.RouteProps {
	if p.Ingress == nil {
		return nil
	}
	return &ingress.RouteProps{
		Name:        p.Name,
		Namespace:   p.Namespace,
		Config:      p.Ingress,
		Middlewares: p.Middlewares,
		HealthCheck: p.HealthCheck,
	}
}

func InitContainersToK8s(containers []*plumagetemplate.InitContainer) *[]*k8s.Container {
	var k8sContainers []*k8s.Container
	for _, container := range containers {
		init := initContainerToK8s(container)
		k8sContainers = append(k8sContainers, init)
	}
	return &k8sContainers
}

func initContainerToK8s(container *plumagetemplate.InitContainer) *k8s.Container {
	commands := StringSliceToK8s(container.Commands)
	resources := ContainerResources(container.Resources)

	return &k8s.Container{
		Name:      jsii.String(container.Name),
		Image:     jsii.String(container.Image),
		Command:   commands,
		Resources: resources,
	}
}

func PortName(rpc bool, count int) string {
	var protocol string
	switch rpc {
	case false:
		protocol = "http"
	case true:
		protocol = "grpc"
	}
	return fmt.Sprintf("%s-%d", protocol, count)
}

func GetProtocol(prt string) string {
	protocol := strings.ToLower(prt)
	switch protocol {
	case "tcp":
		return "TCP"
	case "udp":
		return "UDP"
	default:
		return "TCP"
	}
}

func StringMapToEnv(m map[string]string) *[]*k8s.EnvVar {
	var env []*k8s.EnvVar
	for k, v := range m {
		env = append(env, &k8s.EnvVar{
			Name:  &k,
			Value: &v,
		})
	}
	return &env
}

func StringMapToK8s(m map[string]string) *map[string]*string {
	var k8sMap map[string]*string
	for k, v := range m {
		k8sMap[k] = &v
	}
	return &k8sMap
}

func StringSliceToK8s(sl []string) *[]*string {
	var k8sSlice []*string
	for _, v := range sl {
		k8sSlice = append(k8sSlice, &v)
	}
	return &k8sSlice
}

func loadMonitoringEnv(values, config, env map[string]string) map[string]string {
	for k, v := range values {
		if _, ok := env[k]; ok {
			env[k] = v
		}
	}
	//If a value in the config map is a key in the values map
	//Set  the key from the config map - to the value of the value map's value
	// In the env map
	//So that some keys cna be provided for non-standard otel etc env variables
	for key, configValue := range config {
		if v, ok := values[configValue]; ok {
			env[key] = v
		}
	}
	return env
}
