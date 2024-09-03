package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

const (
	defaultIngressName  = "traefik-ingress-controller"
	defaultTraefikImage = "traefik:v3.1"
)

func defaultIngressLabels() *map[string]*string {
	labels := map[string]*string{
		"app": jsii.String("traefik"),
	}
	return &labels
}

func defaultIngressServiceAccountProps(namespace string) kplus.ServiceAccountProps {
	metadata := defaultIngressServiceAccountMetadata(namespace)
	return kplus.ServiceAccountProps{
		Metadata: &metadata,
	}
}

func defaultIngressServiceAccountMetadata(namespace string) cdk8s.ApiObjectMetadata {
	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
	}
}

func NewIngressServiceAccount(scope constructs.Construct, namespace string) kplus.ServiceAccount {
	props := defaultIngressServiceAccountProps(namespace)
	account := kplus.NewServiceAccount(scope, jsii.String(defaultIngressName), &props)
	return account
}

func defaultIngressServiceMetadata(namespace string) cdk8s.ApiObjectMetadata {

	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
	}
}

func defaultIngressServiceProps(namespace string) *kplus.ServiceProps {
	metadata := defaultIngressServiceMetadata(namespace)

	props := &kplus.ServiceProps{
		Metadata: &metadata,

		Ports: &[]*kplus.ServicePort{
			&kplus.ServicePort{
				Name:       jsii.String("web"),
				Protocol:   "TCP",
				TargetPort: jsii.Number(8000),
				Port:       jsii.Number(8000),
			},
			&kplus.ServicePort{
				Name:       jsii.String("admin"),
				Protocol:   "TCP",
				TargetPort: jsii.Number(8080),
				Port:       jsii.Number(8080),
			},
		},
		//Selector: &podSelector,
	}
	return props
}

func DefaultIngressService(scope constructs.Construct, namespace string) kplus.Service {

	props := defaultIngressServiceProps(namespace)
	service := kplus.NewService(scope, jsii.String("traefik"), props)
	service.SelectLabel(jsii.String("app"), jsii.String("traefik"))
	return service
}

func defaultIngressDeploymentMetadata(namespace string) cdk8s.ApiObjectMetadata {
	labels := defaultIngressLabels()

	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
		Labels:    labels,
	}
}

func defaultIngressDeploymentContainers() *kplus.ContainerProps {
	args := &[]*string{
		jsii.String("--api.insecure"),
		jsii.String("--entryPoints.web.Address=:8000"),
		jsii.String("--providers.kubernetescrd"),
	}

	return &kplus.ContainerProps{
		Args:            args,
		ImagePullPolicy: kplus.ImagePullPolicy_IF_NOT_PRESENT,
		Name:            nil,
		Ports: &[]*kplus.ContainerPort{
			&kplus.ContainerPort{
				Number:   jsii.Number(8000),
				Name:     jsii.String("web"),
				Protocol: "TCP",
			},
			&kplus.ContainerPort{
				Number:   jsii.Number(8080),
				Name:     jsii.String("admin"),
				Protocol: "TCP",
			},
		},
		Image: jsii.String(defaultTraefikImage),
	}
}

func defaultIngressDeploymentProps(scope constructs.Construct, namespace string) *kplus.DeploymentProps {
	metadata := defaultIngressDeploymentMetadata(namespace)
	defaultContainer := defaultIngressDeploymentContainers()
	containers := &[]*kplus.ContainerProps{
		defaultContainer,
	}
	serviceAccount := NewIngressServiceAccount(scope, namespace)

	return &kplus.DeploymentProps{
		Metadata:                     &metadata,
		AutomountServiceAccountToken: nil,
		Containers:                   containers,
		ServiceAccount:               serviceAccount,
		Replicas:                     jsii.Number(1),
	}
}

func DefaultIngressDeployment(scope constructs.Construct, id string, namespace string) kplus.Deployment {

	props := defaultIngressDeploymentProps(scope, namespace)
	return kplus.NewDeployment(scope, jsii.String(id), props)
}

//func OtherIngress(scope constructs.Construct, namespace string) k8s.KubeService {
//	label := map[string]*string{"app": jsii.String("traefik")}
//
//	return k8s.NewKubeService(scope, jsii.String("better-service"), &k8s.KubeServiceProps{
//		Metadata: &k8s.ObjectMeta{
//			Name:      jsii.String(defaultIngressName),
//			Namespace: jsii.String(namespace),
//		},
//		Spec: &k8s.ServiceSpec{
//			LoadBalancerClass:        nil,
//			LoadBalancerIp:           nil,
//			LoadBalancerSourceRanges: nil,
//			Ports: &[]*k8s.ServicePort{
//				&k8s.ServicePort{
//					Name:       jsii.String("web"),
//					Protocol:   jsii.String("TCP"),
//					Port:       jsii.Number(8080),
//					TargetPort: k8s.IntOrString_FromNumber(jsii.Number(8080)),
//				},
//				&k8s.ServicePort{
//					Name:     jsii.String("admin"),
//					Protocol: jsii.String("TCP"),
//
//					Port:       jsii.Number(8000),
//					TargetPort: k8s.IntOrString_FromNumber(jsii.Number(8000)),
//				},
//			},
//			Selector: &label,
//		},
//	})
//}
