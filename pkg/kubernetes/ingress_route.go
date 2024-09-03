package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
)

const (
	defaultRouteType  = "Rule"
	defaultEntryPoint = "web"
)

func newIngressRoute(scope constructs.Construct, id string) traefikio.IngressRoute {
	return nil
}

//func newProps(scope constructs.Construct, id string) traefikio.IngressRouteProps {
//	return nil
//}

//func newSpec(scope constructs.Construct, id string) traefikio.IngressRouteSpec {
//	return traefikio.IngressRouteSpec{
//		Routes: traefikio.IngressRouteSpecRoutes{
//			Kind:        "",
//			Match:       nil,
//			Middlewares: nil,
//			Priority:    nil,
//			Services:    nil,
//			Syntax:      nil,
//		},
//		EntryPoints: nil,
//		Tls:         nil,
//	}
//}

func defaultIngressRouteAnnotations(name string) *map[string]*string {
	return nil
}

func defaultIngressRouteLabels(name string) *map[string]*string {
	return nil
}

func defaultIngressRouteMetadata(svcName string) *cdk8s.ApiObjectMetadata {
	namespace := TestbedNamespace

	return &cdk8s.ApiObjectMetadata{
		Namespace:   jsii.String(namespace),
		Annotations: defaultIngressRouteAnnotations(svcName),
		Labels:      defaultIngressRouteLabels(svcName),
	}
}

func ingressRouteSpec() *traefikio.IngressRouteSpec {
	entryPoints := []*string{jsii.String(defaultEntryPoint)}

	return &traefikio.IngressRouteSpec{
		Routes:      nil,
		EntryPoints: &entryPoints,
		Tls:         nil,
	}
}

func newIngressRouteProps(svcName string) *traefikio.IngressRouteProps {
	metadata := defaultIngressRouteMetadata(svcName)
	return &traefikio.IngressRouteProps{
		Metadata: metadata,
		Spec:     nil,
	}
}

type RouteProps struct {
	ServiceName        string
	Namespace          string
	HostAddress        string
	PathPrefix         string
	Port               RoutePort
	HealthCheck        string
	EnableLoadBalancer bool
}

type RoutePort int

func (r *RoutePort) Value() interface{} {
	return int(*r)
}

func ToRoutePort(port int) *RoutePort {
	routePort := RoutePort(port)
	return &routePort
}

func newRoute(r *RouteProps) *traefikio.IngressRouteSpecRoutes {
	rule := fmt.Sprintf("Host('%s') && PathPrefix('%s')", r.HostAddress, r.PathPrefix)
	middlewares := ingressRouteMiddlewareReferences(r.ServiceName, r.Namespace)
	service := ingressRouteRouteService(r)
	services := []*traefikio.IngressRouteSpecRoutesServices{
		service,
	}

	return &traefikio.IngressRouteSpecRoutes{
		Kind:        defaultRouteType,
		Match:       &rule,
		Middlewares: &middlewares,
		Priority:    nil,
		Services:    &services,
		Syntax:      nil,
	}
}

func ingressRouteMiddlewareReferences(serviceName, nameSpace string) []*traefikio.IngressRouteSpecRoutesMiddlewares {
	retryName := fmt.Sprintf("%s-retry", serviceName)
	circuitBreakerName := fmt.Sprintf("%s-circuit-breaker", serviceName)
	rateLimitName := fmt.Sprintf("%s-rate-limit", serviceName)

	return []*traefikio.IngressRouteSpecRoutesMiddlewares{
		&traefikio.IngressRouteSpecRoutesMiddlewares{
			Name:      &retryName,
			Namespace: &nameSpace,
		},
		&traefikio.IngressRouteSpecRoutesMiddlewares{
			Name:      &circuitBreakerName,
			Namespace: &nameSpace,
		},
		&traefikio.IngressRouteSpecRoutesMiddlewares{
			Name:      &rateLimitName,
			Namespace: &nameSpace,
		},
	}
}

//TODO: Healthcheck and kinf

func ingressRouteRouteService(r *RouteProps) *traefikio.IngressRouteSpecRoutesServices {
	return &traefikio.IngressRouteSpecRoutesServices{
		Name:               &r.ServiceName,
		HealthCheck:        nil,
		Kind:               "",
		Namespace:          &r.Namespace,
		NativeLb:           &r.EnableLoadBalancer,
		NodePortLb:         nil,
		PassHostHeader:     nil,
		Port:               &r.Port,
		ResponseForwarding: nil,
		Scheme:             nil,
		ServersTransport:   nil,
		Sticky:             nil,
		Strategy:           nil,
		Weight:             nil,
	}
}
