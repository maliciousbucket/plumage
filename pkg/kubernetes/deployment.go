package kubernetes

import (
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/types"
)

func defaultServiceDeploymentLabels(svcName string) *map[string]*string {
	return nil
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

func newServiceDeploymentProps(c *types.ContainerSpec, svcName string, namespace string) (*cdk8splus30.DeploymentProps, error) {
	metadata := defaultServiceDeploymentMetadata(svcName, namespace)
	containerProps, err := c.ToContainerProps()
	if err != nil {
		return nil, err
	}

	return cdk8splus30.DeploymentProps{
		Metadata:                     metadata,
		AutomountServiceAccountToken: nil,
		Containers: &[]*cdk8splus30.ContainerProps{
			containerProps,
		},
		Dns:                    nil,
		HostAliases:            nil,
		HostNetwork:            nil,
		InitContainers:         nil,
		Isolate:                nil,
		RestartPolicy:          "",
		SecurityContext:        nil,
		ServiceAccount:         nil,
		TerminationGracePeriod: nil,
		Volumes:                nil,
		PodMetadata:            nil,
		Select:                 nil,
		Spread:                 nil,
		MinReady:               nil,
		ProgressDeadline:       nil,
		Replicas:               nil,
		Strategy:               nil,
	}, nil

}
