package autoscaling

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	autoscaling "github.com/maliciousbucket/plumage/imports/autoscalingk8sio"
)

const (
	memory = "memory"
	cpu    = "cpu"
)

type VerticalAutoscalerProps struct {
	Name          string
	Namespace     string
	ContainerName string
	ControlCpu    bool
	ControlMem    bool
	MinCpuMillis  int
	MaxCpuMillis  int
	MinMemMillis  int
	MaxMemMillis  int
}

func NewVerticalAutoscaler(scope constructs.Construct, id string, props *VerticalAutoscalerProps) autoscaling.VerticalPodAutoscaler {

	metadata := verticalAutoscalerMetadata(props.Namespace, props.Name)
	resources := verticalScalingResourcePolicies(props)
	updatePolicy := verticalScalingUpdatePolicies()
	return autoscaling.NewVerticalPodAutoscalerV1Beta2(scope, jsii.String(id), &autoscaling.VerticalPodAutoscalerV1Beta2Props{
		Metadata: metadata,
		Spec: &autoscaling.VerticalPodAutoscalerV1Beta2Spec{
			TargetRef: &autoscaling.VerticalPodAutoscalerV1Beta2SpecTargetRef{
				Kind:       jsii.String(DeploymentKind),
				Name:       jsii.String(props.Name),
				ApiVersion: jsii.String(AppsV1),
			},
			ResourcePolicy: resources,
			UpdatePolicy:   updatePolicy,
		},
	})
}

func verticalAutoscalerMetadata(ns string, appLabel string) *cdk8s.ApiObjectMetadata {
	labels := verticalAutoscalerLabels(appLabel)
	name := fmt.Sprintf("%s-vpa", appLabel)
	return &cdk8s.ApiObjectMetadata{

		Name:      jsii.String(name),
		Namespace: jsii.String(ns),
		Labels:    labels,
	}
}

func verticalAutoscalerLabels(appLabel string) *map[string]*string {
	labels := map[string]*string{
		AppLabel: &appLabel,
	}
	return &labels
}

func verticalScalingResourcePolicies(props *VerticalAutoscalerProps) *autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicy {
	containerName := "*"
	if props.ContainerName != "" {
		containerName = props.ContainerName
	}
	minAllowed := map[string]autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed{}
	maxAllowed := map[string]autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed{}
	if props.ControlCpu {
		if props.MinCpuMillis != 0 {
			millis := fmt.Sprintf("%dMi", props.MinCpuMillis)
			minCpu := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(millis))
			minAllowed[cpu] = minCpu
		}

		if props.MaxCpuMillis != 0 {
			millis := fmt.Sprintf("%dMi", props.MaxCpuMillis)
			maxCpu := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(millis))
			maxAllowed[cpu] = maxCpu
		}

	}

	if props.ControlMem {
		if props.MinMemMillis != 0 {
			millis := fmt.Sprintf("%dMi", props.MinMemMillis)
			minMem := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(millis))
			minAllowed[memory] = minMem
		}
		if props.MaxMemMillis != 0 {
			millis := fmt.Sprintf("%dMi", props.MaxMemMillis)
			maxMem := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(millis))
			maxAllowed[memory] = maxMem
		}
	}

	return &autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicy{
		ContainerPolicies: &[]*autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies{
			{
				ContainerName: jsii.String(containerName),
				MaxAllowed:    &maxAllowed,
				MinAllowed:    &minAllowed,
				Mode:          autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_AUTO,
			},
		},
	}
}

func verticalScalingUpdatePolicies() *autoscaling.VerticalPodAutoscalerV1Beta2SpecUpdatePolicy {
	return &autoscaling.VerticalPodAutoscalerV1Beta2SpecUpdatePolicy{
		UpdateMode: autoscaling.VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_AUTO,
	}
}
