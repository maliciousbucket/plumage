package kubernetes

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
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

func defaultIngressServiceAccountProps(namespace string) cdk8splus30.ServiceAccountProps {
	metadata := defaultIngressServiceAccountMetadata(namespace)
	return cdk8splus30.ServiceAccountProps{
		Metadata: &metadata,
	}
}

func defaultIngressServiceAccountMetadata(namespace string) cdk8s.ApiObjectMetadata {
	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
	}
}

func NewIngressServiceAccount(scope constructs.Construct, namespace string) cdk8splus30.ServiceAccount {
	props := defaultIngressServiceAccountProps(namespace)
	account := cdk8splus30.NewServiceAccount(scope, jsii.String(defaultIngressName), &props)
	return account
}

//func newTraefikService() *cdk8splus30.Service {
//	return cdk8splus30.NewService(cdk8splus30.ServiceProps{
//		Metadata:                 nil,
//		ClusterIP:                nil,
//		ExternalIPs:              nil,
//		ExternalName:             nil,
//		LoadBalancerSourceRanges: nil,
//		Ports:                    nil,
//		PublishNotReadyAddresses: nil,
//		Selector:                 nil,
//		Type:                     "",
//	})
//}

func defaultIngressServiceMetadata(namespace string) cdk8s.ApiObjectMetadata {

	return cdk8s.ApiObjectMetadata{
		Name:      jsii.String(defaultIngressName),
		Namespace: jsii.String(namespace),
	}
}

func defaultIngressServiceProps(namespace string) *cdk8splus30.ServiceProps {
	metadata := defaultIngressServiceMetadata(namespace)

	props := &cdk8splus30.ServiceProps{
		Metadata: &metadata,

		Ports: &[]*cdk8splus30.ServicePort{
			&cdk8splus30.ServicePort{
				Name:       jsii.String("web"),
				Protocol:   "TCP",
				TargetPort: jsii.Number(8000),
				Port:       jsii.Number(8000),
			},
			&cdk8splus30.ServicePort{
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

func DefaultIngressService(scope constructs.Construct, namespace string) cdk8splus30.Service {

	props := defaultIngressServiceProps(namespace)
	service := cdk8splus30.NewService(scope, jsii.String("traefik"), props)
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

func defaultIngressDeploymentContainers() *cdk8splus30.ContainerProps {
	args := &[]*string{
		jsii.String("--api.insecure"),
		jsii.String("--entryPoints.web.Address=:8000"),
		jsii.String("--providers.kubernetescrd"),
	}

	return &cdk8splus30.ContainerProps{
		Args:            args,
		ImagePullPolicy: cdk8splus30.ImagePullPolicy_IF_NOT_PRESENT,
		Name:            nil,
		Ports: &[]*cdk8splus30.ContainerPort{
			&cdk8splus30.ContainerPort{
				Number:   jsii.Number(8000),
				Name:     jsii.String("web"),
				Protocol: "TCP",
			},
			&cdk8splus30.ContainerPort{
				Number:   jsii.Number(8080),
				Name:     jsii.String("admin"),
				Protocol: "TCP",
			},
		},
		Image: jsii.String(defaultTraefikImage),
	}
}

func defaultIngressDeploymentProps(scope constructs.Construct, namespace string) *cdk8splus30.DeploymentProps {
	metadata := defaultIngressDeploymentMetadata(namespace)
	defaultContainer := defaultIngressDeploymentContainers()
	containers := &[]*cdk8splus30.ContainerProps{
		defaultContainer,
	}

	serviceAccount := NewIngressServiceAccount(scope, namespace)

	return &cdk8splus30.DeploymentProps{
		Metadata:                     &metadata,
		AutomountServiceAccountToken: nil,
		Containers:                   containers,
		ServiceAccount:               serviceAccount,
		Replicas:                     jsii.Number(1),
	}
}

func DefaultIngressDeployment(scope constructs.Construct, id string, namespace string) cdk8splus30.Deployment {

	props := defaultIngressDeploymentProps(scope, namespace)
	return cdk8splus30.NewDeployment(scope, jsii.String(id), props)
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
