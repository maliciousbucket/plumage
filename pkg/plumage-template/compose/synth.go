package compose

import (
	"fmt"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/autoscaling"
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
	Ports          []plumagetemplate.ServicePort
	Resources      *plumagetemplate.ServiceResources
	Monitoring     *plumagetemplate.MonitoringConfig
	InitContainers []*plumagetemplate.InitContainer
	Scaling        *plumagetemplate.ScalingConfig
	Resilience     *resilience.ResTemplate
	Paths          []ingress.ServicePaths
	Env            map[string]string
	FileMounts     []map[string]string
	DirMounts      []map[string]string
	Middlewares    []string
	Ingress        *ingress.RouteConfig
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

func (p *WebServiceProps) autoScalingProps() *autoscaling.HorizontalAutoScalerProps {
	return &autoscaling.HorizontalAutoScalerProps{
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

func StringSliceToK8s(sl []string) *[]*string {
	var k8sSlice []*string
	for _, v := range sl {
		k8sSlice = append(k8sSlice, &v)
	}
	return &k8sSlice
}
