package kubernetes

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/template"
	"strconv"
)

func defaultServiceDeploymentLabels(svcName string) *map[string]*string {
	labels := map[string]*string{
		KubernetesNameLabel: &svcName,
	}
	return &labels
}

func defaultServiceDeploymentAnnotations(svcName string) *map[string]*string {

	return nil
}

func defaultServiceDeploymentMetadata(svcName string, namespace string) cdk8s.ApiObjectMetadata {
	labels := defaultServiceDeploymentLabels(svcName)
	annotations := defaultServiceDeploymentAnnotations(svcName)

	return cdk8s.ApiObjectMetadata{
		Name:        jsii.String(svcName),
		Namespace:   jsii.String(namespace),
		Labels:      labels,
		Annotations: annotations,
	}
}

func newServiceDeploymentProps(sc template.ServiceConfig) *cdk8splus30.DeploymentProps {
	metadata := defaultServiceDeploymentMetadata(sc.Name, sc.Namespace)
	//containerProps, err := sc.Container.ToContainerProps()
	//if err != nil {
	//	return nil, err
	//}
	podSelect := true

	replicas := 2
	if sc.Service.Resilience.ScalingPolicy.MinReplicas != 0 {
		replicas = int(sc.Service.Resilience.ScalingPolicy.MinReplicas)
	}

	return &cdk8splus30.DeploymentProps{
		Metadata: &metadata,
		Select:   &podSelect,
		Replicas: jsii.Number(replicas),
	}

	//return &cdk8splus30.DeploymentProps{
	//	Metadata:                     &metadata,
	//	AutomountServiceAccountToken: nil,
	//	Containers: &[]*cdk8splus30.ContainerProps{
	//		containerProps,
	//	},
	//	Dns:                    nil,
	//	HostAliases:            nil,
	//	HostNetwork:            nil,
	//	InitContainers:         nil,
	//	Isolate:                nil,
	//	RestartPolicy:          "",
	//	SecurityContext:        nil,
	//	ServiceAccount:         nil,
	//	TerminationGracePeriod: nil,
	//	Volumes:                nil,
	//	PodMetadata:            nil,
	//	Select:                 &podSelect,
	//	Spread:                 nil,
	//	MinReady:               nil,
	//	ProgressDeadline:       nil,
	//	Replicas:               jsii.Number(replicas),
	//	Strategy:               nil,
	//}, nil

}

func NewServiceDeployment(scope constructs.Construct, sc template.ServiceConfig) (cdk8splus30.Deployment, error) {
	props := newServiceDeploymentProps(sc)
	name := fmt.Sprintf("%s-%s", sc.Name, "deployment")
	deployment := cdk8splus30.NewDeployment(scope, jsii.String(name), props)

	containerProps, err := sc.Container.ToContainerProps()
	if err != nil {
		return nil, err
	}
	container := deployment.AddContainer(containerProps)

	addPromAnnotations(deployment, sc.Service.Monitoring)

	_, err = AddVolumeConfigMaps(scope, container, sc.Namespace, sc.Container.Volumes)
	if err != nil {
		return nil, err
	}

	labelSelector := cdk8splus30.LabelSelector_Of(&cdk8splus30.LabelSelectorOptions{
		Expressions: nil,
		Labels: &map[string]*string{
			KubernetesNameLabel: &sc.Name,
		},
	})

	deployment.Select(labelSelector)

	return deployment, nil
}

func addPromAnnotations(deployment cdk8splus30.Deployment, m *template.MonitoringTemplate) {
	if m != nil {
		if m.ScrapeConfig != nil {
			path := "/metrics"

			if m.ScrapeConfig.MetricsPath != "" {
				path = m.ScrapeConfig.MetricsPath
			}

			deployment.Metadata().AddAnnotation(jsii.String(PrometheusScrapeAnnotation), jsii.String("true"))
			if m.ScrapeConfig.MetricsPort != 0 {
				deployment.Metadata().AddAnnotation(jsii.String(PrometheusPortAnnotation), jsii.String(strconv.FormatInt(int64(m.ScrapeConfig.MetricsPort), 10)))

			}
			deployment.Metadata().AddAnnotation(jsii.String(PrometheusPathAnnotation), jsii.String(path))
		}
	}
}
