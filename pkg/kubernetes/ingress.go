package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio"
	"github.com/maliciousbucket/plumage/pkg/imports"
)

const (
	defaultIngressName   = "traefik-ingress-controller"
	defaultTraefikImage  = "traefik:v3.1"
	entrypointAnnotation = "traefik.ingress.kubernetes.io/router-entrypoints"
	traefikCrdUrl        = "https://raw.githubusercontent.com/traefik/traefik/refs/heads/master/docs/content/reference/dynamic-configuration/kubernetes-crd-definition-v1.yml"
)

func defaultIngressServiceAccountMetadata(namespace string) cdk8s.ApiObjectMetadata {
	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
	}
}

func NewTraefikIngress(scope constructs.Construct, id string, ns string) constructs.Construct {
	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		DisableResourceNameHashes: jsii.Bool(true),
		Namespace:                 jsii.String(ns),
	})

	roleImp := cdk8s.NewInclude(chart, jsii.String("traefik-import-role"), &cdk8s.IncludeProps{
		Url: jsii.String("dist/include/traefik-cluster-role.yaml"),
	})

	roleObj := *roleImp.ApiObjects()
	patch := cdk8s.JsonPatch_Add(jsii.String("/metadata/namespace"), jsii.String("galah-test-bed"))
	roleObj[0].AddJsonPatch(patch)

	acc := kplus.NewServiceAccount(chart, jsii.String("yo,"), &kplus.ServiceAccountProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name:      jsii.String("traefik-ingress-controller"),
			Namespace: jsii.String(ns),
		},
	})

	role := kplus.ClusterRole_FromClusterRoleName(chart, jsii.String("traefik-cluster-role"), roleObj[0].Name())
	binding := kplus.NewClusterRoleBinding(chart, jsii.String("traefik-role-binding"), &kplus.ClusterRoleBindingProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Annotations:     nil,
			Finalizers:      nil,
			Labels:          nil,
			Name:            jsii.String("traefik-role-binding"),
			Namespace:       nil,
			OwnerReferences: nil,
		},
		Role: role,
	})

	binding.AddSubjects(acc)
	binding.AddSubjects()

	args := traefikIngressArgs()
	labels := traefikIngressLabels()
	annotations := traefikIngressAnnotations(ns)

	deployment := kplus.NewDeployment(chart, jsii.String(id), &kplus.DeploymentProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Namespace:   jsii.String(ns),
			Name:        jsii.String("traefik"),
			Labels:      &labels,
			Annotations: &annotations,
		},
		AutomountServiceAccountToken: jsii.Bool(true),
		SecurityContext: &kplus.PodSecurityContextProps{
			EnsureNonRoot: jsii.Bool(false),
		},
		ServiceAccount: acc,
		PodMetadata: &cdk8s.ApiObjectMetadata{
			Labels: &labels,
		},
		Replicas: jsii.Number(1),
		Strategy: nil,
	})
	deployment.AddContainer(&kplus.ContainerProps{
		Args:            &args,
		ImagePullPolicy: kplus.ImagePullPolicy_IF_NOT_PRESENT,
		Name:            jsii.String("traefik"),
		Ports: &[]*kplus.ContainerPort{
			{

				Name:   jsii.String("web"),
				Number: jsii.Number(8000),
			},
			{

				Name:   jsii.String("admin"),
				Number: jsii.Number(8080),
			},
			{

				Name:   jsii.String("websecure"),
				Number: jsii.Number(4443),
			},
		},
		SecurityContext: &kplus.ContainerSecurityContextProps{
			EnsureNonRoot: jsii.Bool(false),
		},
		Image: jsii.String(defaultTraefikImage),
	})
	selector := kplus.LabelSelector_Of(&kplus.LabelSelectorOptions{
		Labels: &labels,
	})
	deployment.Select(selector)

	web := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
		Name: jsii.String("traefik-web-service"),
		Ports: &[]*kplus.ServicePort{
			{
				Protocol:   kplus.Protocol_TCP,
				TargetPort: jsii.Number(8000),
				Port:       jsii.Number(8000),
			},
		},
		ServiceType: kplus.ServiceType_LOAD_BALANCER,
	})

	web.SelectLabel(jsii.String("app"), jsii.String("traefik"))

	db := deployment.ExposeViaService(&kplus.DeploymentExposeViaServiceOptions{
		Name: jsii.String("traefik-dashboard-service"),
		Ports: &[]*kplus.ServicePort{
			{
				Protocol:   kplus.Protocol_TCP,
				TargetPort: jsii.Number(8080),
				Port:       jsii.Number(8080),
				Name:       jsii.String("admin"),
			},
			{
				Protocol:   kplus.Protocol_TCP,
				TargetPort: jsii.Number(4443),
				Port:       jsii.Number(4443),
				Name:       jsii.String("web-secure"),
			},
		},
		ServiceType: kplus.ServiceType_LOAD_BALANCER,
	})

	db.SelectLabel(jsii.String("app"), jsii.String("traefik"))

	ig := kplus.NewIngress(chart, jsii.String("traefik-ingress"), &kplus.IngressProps{
		Metadata:  nil,
		ClassName: jsii.String("traefik"),
		Rules:     nil,
		Tls:       nil,
	})

	ig.Metadata().AddLabel(jsii.String("name"), jsii.String("traefik-web-ingress"))
	ig.Metadata().AddAnnotation(jsii.String("traefik.ingress.kubernetes.io/router.entrypoints"), jsii.String("web"))

	be := kplus.IngressBackend_FromService(db, &kplus.ServiceIngressBackendOptions{Port: jsii.Number(8080)})

	ig.AddDefaultBackend(be)
	web.ExposeViaIngress(jsii.String("/testbed"), &kplus.ExposeServiceViaIngressOptions{
		Ingress:  ig,
		PathType: kplus.HttpIngressPathType_PREFIX,
	})
	traefikRemovePrefixMiddleware(chart, "strip-dashboard")
	stripTestBedMiddleware(chart, "strip-testbed")
	newDashbaordRoute(chart, "dashboard-route", db)
	grafanaRoute(chart, ns)

	imports.NewInclude(chart, "traefik-crds", traefikCrdUrl)

	return chart

}

func traefikIngressAnnotations(ns string) map[string]*string {
	annotations := map[string]*string{
		//entrypointAnnotation: jsii.String("web"),
	}
	return annotations
}

func traefikIngressLabels() map[string]*string {
	labels := map[string]*string{
		"app.kubernetes.io/name": jsii.String("traefik"),
		"app":                    jsii.String("traefik"),
	}
	return labels
}

func traefikIngressArgs() []*string {
	args := []*string{
		jsii.String("--api.insecure"),
		jsii.String("--accesslog"),
		jsii.String("--entryPoints.web.Address=:8000"),
		jsii.String("--providers.kubernetescrd"),

		jsii.String("--entryPoints.websecure.Address=:4443"),
		jsii.String("--entryPoints.traefik.Address=:8080"),
		jsii.String("--providers.kubernetescrd.namespaces=default,galah-testbed,galah-monitoring,argocd"),
		jsii.String("--providers.kubernetescrd.allowCrossNamespace=true"),
		jsii.String("--api.insecure=true"),
		jsii.String("--metrics.prometheus=true"),
		jsii.String("--metrics.prometheus.addEntryPointsLabels=true"),
	}
	return args
}

func newDashbaordRoute(scope constructs.Construct, id string, service kplus.Service) traefikio.IngressRoute {
	return traefikio.NewIngressRoute(scope, jsii.String(id), &traefikio.IngressRouteProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("dashboard-route"),
		},
		Spec: &traefikio.IngressRouteSpec{
			EntryPoints: &[]*string{
				jsii.String("traefik"),
			},
			Routes: &[]*traefikio.IngressRouteSpecRoutes{
				{
					Kind:  traefikio.IngressRouteSpecRoutesKind_RULE,
					Match: jsii.String("PathPrefix(`/dashboard`)"),
					Middlewares: &[]*traefikio.IngressRouteSpecRoutesMiddlewares{
						{
							Name: jsii.String("strip-dashboard"),
						},
					},
					Services: &[]*traefikio.IngressRouteSpecRoutesServices{
						{
							Name: service.Name(),
							Port: traefikio.IngressRouteSpecRoutesServicesPort_FromNumber(jsii.Number(8080)),
						},
					},
				},
			},
		},
	})
}

func traefikRemovePrefixMiddleware(scope constructs.Construct, id string) traefikio.Middleware {
	middleware := traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("strip-dashboard"),
		},
		Spec: &traefikio.MiddlewareSpec{
			StripPrefix: &traefikio.MiddlewareSpecStripPrefix{
				Prefixes: &[]*string{jsii.String("/dashboard")},
			},
		},
	})
	return middleware
}

func stripTestBedMiddleware(scope constructs.Construct, id string) traefikio.Middleware {
	middleware := traefikio.NewMiddleware(scope, jsii.String(id), &traefikio.MiddlewareProps{
		Metadata: &cdk8s.ApiObjectMetadata{
			Name: jsii.String("strip-testbed"),
		},
		Spec: &traefikio.MiddlewareSpec{
			StripPrefix: &traefikio.MiddlewareSpecStripPrefix{
				Prefixes: &[]*string{jsii.String("/testbed")},
			},
		},
	})
	return middleware
}
