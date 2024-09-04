package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/template"
)

const (
	ClusterIpType   = "ClusterIP"
	StatefulSetKind = "StatefulSet"
	ServiceKind     = "Service"
)

func NewServiceChart(scope constructs.Construct, config template.ServiceConfig, hostAddress string) (cdk8s.Chart, error) {
	chart := cdk8s.NewChart(scope, jsii.String(config.Name), nil)

	deployment, err := NewServiceDeployment(chart, config)
	if err != nil {
		return nil, err
	}

	servicePorts := ServicePorts(&config)

	service := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
		Name:        jsii.String(config.Name),
		Ports:       &servicePorts,
		ServiceType: kplus.ServiceType_CLUSTER_IP,
	})
	service.SelectLabel(jsii.String("app"), jsii.String(config.Name))

	newServiceMiddleware(chart, config.Service)

	newServiceIngressRoute(chart, hostAddress, config)

	return chart, nil
}

//func NewServiceProps(config template.ServiceConfig) *kplus.ServiceProps {
//	metadata := DefaultServiceMetadata(config.Namespace, config.Name)
//	ports := ServicePorts(&config)
//	return &kplus.ServiceProps{
//		Metadata: &metadata,
//		Ports:    &ports,
//		Type:     kplus.ServiceType_CLUSTER_IP,
//	}
//
//	return nil
//}

func newServiceDeployment(scope constructs.Construct, config template.ServiceConfig) (kplus.Deployment, error) {
	deployment, err := NewServiceDeployment(scope, config)
	if err != nil {
		return nil, err
	}

	return deployment, nil
}

func newServiceIngressRoute(scope constructs.Construct, hostAddress string, config template.ServiceConfig) traefikio.IngressRoute {
	var props []*RouteProps
	for _, path := range config.Service.Paths {
		routeProps := &RouteProps{
			ServiceName:        config.Name,
			Namespace:          config.Namespace,
			HostAddress:        hostAddress,
			PathPrefix:         path.Path,
			Port:               *ToRoutePort(path.Port),
			HealthCheck:        "",
			EnableLoadBalancer: false,
		}
		props = append(props, routeProps)
	}

	ingressProps := NewIngressRouteProps(config.Namespace, config.Name, props)

	route := traefikio.NewIngressRoute(scope, jsii.String(config.Name), ingressProps)

	return route

}

func newServiceMiddleware(scope constructs.Construct, s *template.ServiceTemplate) constructs.Construct {
	if s.Resilience.RetryPolicy != nil {
		NewRetryMiddleware(scope, s.Name, s.Resilience.RetryPolicy)
	} else {
		EmptyMiddleware(scope, s.Name, MiddlewareNameRetry)
	}

	if s.Resilience.CircuitBreakerPolicy != nil {
		NewCircuitBreakerMiddleware(scope, s.Name, s.Resilience.CircuitBreakerPolicy)
	} else {
		EmptyMiddleware(scope, s.Name, MiddlewareNameCircuitBreaker)
	}

	if s.Resilience.RateLimitPolicy != nil {
		NewRateLimitMiddleware(scope, s.Name, s.Resilience.RateLimitPolicy)
	}

	return scope
}
