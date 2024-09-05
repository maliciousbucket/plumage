package k8s

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"time"
)

const (
	RouteTypeRule     = traefikio.IngressRouteSpecRoutesKind_RULE
	defaultEntryPoint = "web"
)

type IngressRouteProps struct {
	Name        string
	Namespace   string
	Config      *IngressConfig
	Middlewares []string
	HealthCheck *plumagetemplate.HttpProbe
}

func NewIngressRoute(scope constructs.Construct, id string, props *IngressRouteProps) traefikio.IngressRoute {
	metadata := ingressRouteMetaData(props.Namespace, props.Name)
	serviceRoutes := routes(props)

	return traefikio.NewIngressRoute(scope, jsii.String(id), &traefikio.IngressRouteProps{
		Metadata: metadata,
		Spec: &traefikio.IngressRouteSpec{
			Routes: serviceRoutes,
			EntryPoints: &[]*string{
				jsii.String(defaultEntryPoint),
			},
			Tls: nil,
		},
	})
}

func routes(props *IngressRouteProps) *[]*traefikio.IngressRouteSpecRoutes {
	services := ingressRouteServices(props)
	middlewares := ingressRouteMiddlewareReferences(props.Namespace, props.Name, props.Middlewares)

	var ingressRoutes []*traefikio.IngressRouteSpecRoutes
	for _, path := range props.Config.Paths {
		routeSpecification := routeSpec(services, path.Path, middlewares, props.Config.Host)
		ingressRoutes = append(ingressRoutes, &routeSpecification)
	}

	return &ingressRoutes
}

func routeSpec(services []*traefikio.IngressRouteSpecRoutesServices, path string,
	middlewares []*traefikio.IngressRouteSpecRoutesMiddlewares, host string) traefikio.IngressRouteSpecRoutes {
	rule := fmt.Sprintf("Host('%s') && PathPrefix('%s')", host, path)
	return traefikio.IngressRouteSpecRoutes{
		Kind:        RouteTypeRule,
		Match:       jsii.String(rule),
		Middlewares: &middlewares,
		Priority:    nil,
		Services:    &services,
		Syntax:      nil,
	}
}

func ingressRouteServices(props *IngressRouteProps) []*traefikio.IngressRouteSpecRoutesServices {
	healthCheck := ingressRouteHealthCheck(props.HealthCheck)

	var services []*traefikio.IngressRouteSpecRoutesServices
	for _, path := range props.Config.Paths {
		services = append(services, &traefikio.IngressRouteSpecRoutesServices{
			Name:        jsii.String(props.Name),
			HealthCheck: healthCheck,
			Kind:        traefikio.IngressRouteSpecRoutesServicesKind_SERVICE,
			Namespace:   jsii.String(props.Namespace),
			NativeLb:    jsii.Bool(props.Config.EnableLoadBalancer),
			Port:        traefikio.IngressRouteSpecRoutesServicesPort_FromNumber(jsii.Number(path.Port)),
			Strategy:    nil,
			Weight:      nil,
		})
	}

	return services
}

func ingressRouteMiddlewareReferences(ns string, appLabel string, mw []string) []*traefikio.IngressRouteSpecRoutesMiddlewares {
	var middlewares []*traefikio.IngressRouteSpecRoutesMiddlewares
	for _, m := range mw {
		name := fmt.Sprintf("%s-%s", appLabel, m)
		middlewares = append(middlewares, &traefikio.IngressRouteSpecRoutesMiddlewares{
			Name:      jsii.String(name),
			Namespace: jsii.String(ns),
		})
	}
	return middlewares
}

func ingressRouteHealthCheck(check *plumagetemplate.HttpProbe) *traefikio.IngressRouteSpecRoutesServicesHealthCheck {
	if check == nil {
		return nil
	}

	if check.Path == "" || check.Port == 0 {
		return nil
	}

	var interval traefikio.IngressRouteSpecRoutesServicesHealthCheckInterval
	if check.PeriodSeconds != 0 {
		intervalPeriod := time.Second * time.Duration(check.PeriodSeconds)
		interval = traefikio.IngressRouteSpecRoutesServicesHealthCheckInterval_FromString(jsii.String(intervalPeriod.String()))

	}

	var timeout traefikio.IngressRouteSpecRoutesServicesHealthCheckTimeout
	if check.TimeoutSeconds != 0 {
		timeoutPeriod := time.Second * time.Duration(check.TimeoutSeconds)
		timeout = traefikio.IngressRouteSpecRoutesServicesHealthCheckTimeout_FromString(jsii.String(timeoutPeriod.String()))

	}
	return &traefikio.IngressRouteSpecRoutesServicesHealthCheck{
		Interval: interval,
		Method:   jsii.String("GET"),
		Mode:     jsii.String("http"),
		Path:     jsii.String(check.Path),
		Port:     jsii.Number(check.Port),
		Timeout:  timeout,
	}

}

func ingressRouteMetaData(ns string, appLabel string) *cdk8s.ApiObjectMetadata {
	name := fmt.Sprintf("%s-%s", appLabel, "ingressroute")
	labels := ingressRouteLabels()
	annotations := ingressRouteAnnotations()
	return &cdk8s.ApiObjectMetadata{
		Annotations: annotations,
		Labels:      labels,
		Name:        jsii.String(name),
		Namespace:   jsii.String(ns),
	}
}

func ingressRouteLabels() *map[string]*string {
	return nil
}

func ingressRouteAnnotations() *map[string]*string {
	return nil
}
