package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/template"
)

const (
	ClusterIpType   = "ClusterIP"
	StatefulSetKind = "StatefulSet"
	ServiceKind     = "Service"
)

func NewServiceChart(scope constructs.Construct, config template.ServiceConfig, hostAddress string) (constructs.Construct, error) {
	//chart := cdk8s.NewChart(scope, jsii.String(config.Name), nil)

	chart := constructs.NewConstruct(scope, jsii.String(config.Name))

	deployment, err := NewServiceDeployment(chart, config)
	if err != nil {
		return nil, err
	}
	var servicePorts []*kplus.ServicePort
	servicePorts = ServicePorts(&config)
	//if len(servicePorts) == 0 {
	//	servicePorts = append(servicePorts, &kplus.ServicePort{
	//		Name:       jsii.String("http"),
	//		NodePort:   nil,
	//		Protocol:   kplus.Protocol_TCP,
	//		TargetPort: jsii.Number(80),
	//		Port:       jsii.Number(80),
	//	})
	//}
	//fmt.Println(servicePorts)
	if len(servicePorts) != 0 {
		service := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
			Name:        jsii.String(config.Name),
			Ports:       &servicePorts,
			ServiceType: kplus.ServiceType_CLUSTER_IP,
		})
		service.SelectLabel(jsii.String("app"), jsii.String(config.Name))
	}

	if config.Service != nil {
		newServiceMiddleware(chart, config.Service)
	} else {
		EmptyServiceMiddleware(scope, config)
	}

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

	if config.Service != nil {
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
	}
	var ingressProps *traefikio.IngressRouteProps
	if len(props) > 0 {
		ingressProps = NewIngressRouteProps(config.Namespace, config.Name, props)
	} else {
		path := fmt.Sprintf("/%s", config.Name)
		prop := &RouteProps{
			ServiceName:        config.Name,
			Namespace:          config.Namespace,
			HostAddress:        hostAddress,
			PathPrefix:         path,
			Port:               80,
			HealthCheck:        "",
			EnableLoadBalancer: false,
		}
		props = append(props, prop)
		ingressProps = NewIngressRouteProps(config.Namespace, config.Name, props)
		//ingressProps = EmptyIngressRouteProps(config.Namespace, config.Name)
	}

	route := traefikio.NewIngressRoute(scope, jsii.String(config.Name), ingressProps)

	return route

}

func newServiceMiddleware(scope constructs.Construct, s *template.ServiceTemplate) constructs.Construct {

	if s.Resilience != nil {
		if s.Resilience.RetryPolicy.RetryAttempts() != 0 {
			NewRetryMiddleware(scope, s.Name, &s.Resilience.RetryPolicy)
		} else {
			EmptyMiddleware(scope, s.Name, MiddlewareNameRetry)
		}

		if s.Resilience.CircuitBreakerPolicy.CircuitBreakerExpression() != "" {
			NewCircuitBreakerMiddleware(scope, s.Name, &s.Resilience.CircuitBreakerPolicy)
		} else {
			EmptyMiddleware(scope, s.Name, MiddlewareNameCircuitBreaker)
		}

		if s.Resilience.RateLimitPolicy.LimitStrategy() != nil {
			NewRateLimitMiddleware(scope, s.Name, &s.Resilience.RateLimitPolicy)
		} else {
			EmptyMiddleware(scope, s.Name, MiddlewareNameRateLimiting)
		}
	}

	return scope
}

func EmptyServiceMiddleware(scope constructs.Construct, config template.ServiceConfig) constructs.Construct {
	EmptyMiddleware(scope, config.Name, MiddlewareNameRateLimiting)
	EmptyMiddleware(scope, config.Name, MiddlewareNameRetry)
	EmptyMiddleware(scope, config.Name, MiddlewareNameCircuitBreaker)
	return scope
}
