package kplus

import (
	"fmt"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	autoscaling "github.com/maliciousbucket/plumage/imports/autoscalingk8sio"
)

const (
	ScalingTypeVertical   = ScalingType("vertical")
	ScalingTypeHorizontal = ScalingType("horizontal")
	ScalingTypeUnknown    = ScalingType("unknown")
)

type ScalingTemplate interface {
	ScalingType() ScalingType
}

type HorizontalScalingTemplate struct {
	Type        string                     `yaml:"type"`
	CPU         *HorizontalResourceScaling `yaml:"cpu"`
	Memory      *HorizontalResourceScaling `yaml:"memory"`
	MinReplicas int32                      `yaml:"minReplicas"`
	MaxReplicas int32                      `yaml:"maxReplicas"`
}

type HorizontalResourceScaling struct {
	Utilization int `yaml:"utilization"`
	Amount      int `yaml:"amount"`
}

type VerticalScalingTemplate struct {
	Type          string `yaml:"type"`
	Version       int    `yaml:"version"`
	MinCpuMillis  int32  `yaml:"minCpuMillis"`
	MaxCpuMillis  int32  `yaml:"maxCpuMillis"`
	MinMemoryMi   int32  `yaml:"minMemoryMb"`
	MaxMemoryMi   int32  `yaml:"maxMemoryMb"`
	ControlMem    bool   `yaml:"controlMemory"`
	ControlCpu    bool   `yaml:"controlCpu"`
	ControlLimits bool   `yaml:"controlLimits"`
	MinReplicas   int32  `yaml:"minReplicas"`
}

func (t *VerticalScalingTemplate) ScalingType() ScalingType {
	return ScalingTypeVertical
}

type ScalingType string

func (t *HorizontalScalingTemplate) ScalingType() ScalingType {
	return ScalingTypeHorizontal
}

func NewAutoScaler(scope constructs.Construct, deployment kplus.Deployment, template ScalingTemplate, id string) constructs.Construct {
	if template == nil {
		return nil
	}

	switch template.ScalingType() {
	case ScalingTypeVertical:
		if vertical, ok := template.(*VerticalScalingTemplate); ok {
			if vertical.Version == 0 || vertical.Version == 1 {
				return newVerticalScaler(scope, deployment, id, vertical)
			}
			return newVerticalScalerV2(scope, deployment, id, vertical)
		}
	case ScalingTypeHorizontal:
		if horizontal, ok := template.(*HorizontalScalingTemplate); ok {
			return addHorizontalAutoScaler(scope, deployment, id, horizontal)
		}
	default:
		return nil
	}
	return nil
}

func addHorizontalAutoScaler(scope constructs.Construct, deployment kplus.Deployment, id string, template *HorizontalScalingTemplate) kplus.HorizontalPodAutoscaler {
	name := fmt.Sprintf("%s-autoscaler", id)
	var minReplicas = jsii.Number(1)
	if deployment.Replicas() != nil && *deployment.Replicas() > 0 {
		minReplicas = deployment.Replicas()
	}
	if template.MinReplicas != 0 {
		minReplicas = jsii.Number(template.MinReplicas)
	}
	var maxReplicas = jsii.Number(3)
	if template.MaxReplicas != 0 && template.MaxReplicas >= template.MinReplicas {
		maxReplicas = jsii.Number(template.MaxReplicas)
	}
	metrics := []kplus.Metric{}

	if template.CPU != nil {
		cpuUtil := cpuUtilizationScalingMetric(template.CPU.Utilization)
		cpuAmt := cpuAmountScalingMetric(template.CPU.Amount)
		metrics = append(metrics, cpuUtil, cpuAmt)

	}

	if template.Memory != nil {
		memUtil := memoryUtilizationScalingMetric(template.Memory.Utilization)
		memAmt := memoryAmountScalingMetric(template.Memory.Amount)
		metrics = append(metrics, memUtil, memAmt)
	}

	scaler := kplus.NewHorizontalPodAutoscaler(scope, jsii.String(name), &kplus.HorizontalPodAutoscalerProps{
		Metadata:    &cdk8s.ApiObjectMetadata{Name: jsii.String(name)},
		MaxReplicas: maxReplicas,
		Target:      deployment,
		Metrics:     &metrics,
		MinReplicas: minReplicas,
		ScaleDown:   nil,
		ScaleUp:     nil,
	})
	return scaler
}

func cpuAmountScalingMetric(amount int) kplus.Metric {
	if amount == 0 {
		return nil
	}

	return kplus.Metric_ResourceCpu(kplus.MetricTarget_Value(jsii.Number(amount)))
}

func cpuUtilizationScalingMetric(utilization int) kplus.Metric {
	if utilization == 0 {
		return nil
	}

	if utilization >= 100 {
		utilization = 90
	}

	return kplus.Metric_ResourceCpu(kplus.MetricTarget_AverageUtilization(jsii.Number(utilization)))
}

func memoryAmountScalingMetric(amount int) kplus.Metric {
	if amount == 0 {
		return nil
	}
	return kplus.Metric_ResourceMemory(kplus.MetricTarget_Value(jsii.Number(amount)))
}

func memoryUtilizationScalingMetric(utilization int) kplus.Metric {
	if utilization == 0 {
		return nil
	}
	if utilization >= 100 {
		utilization = 90
	}
	return kplus.Metric_ResourceMemory(kplus.MetricTarget_AverageUtilization(jsii.Number(utilization)))
}

func newVerticalScalerV2(scope constructs.Construct, deployment kplus.Deployment, id string, template *VerticalScalingTemplate) autoscaling.VerticalPodAutoscalerV1Beta2 {
	name := fmt.Sprintf("%s-autoscaler", id)

	target := autoscaling.VerticalPodAutoscalerV1Beta2SpecTargetRef{
		Kind:       deployment.Kind(),
		Name:       deployment.Name(),
		ApiVersion: deployment.ApiVersion(),
	}

	polices := v2ScalingPolicies(template)

	scaler := autoscaling.NewVerticalPodAutoscalerV1Beta2(scope, jsii.String(name), &autoscaling.VerticalPodAutoscalerV1Beta2Props{
		Metadata: &cdk8s.ApiObjectMetadata{Name: jsii.String(name)},
		Spec: &autoscaling.VerticalPodAutoscalerV1Beta2Spec{
			TargetRef:      &target,
			ResourcePolicy: polices,
			UpdatePolicy: &autoscaling.VerticalPodAutoscalerV1Beta2SpecUpdatePolicy{
				UpdateMode: autoscaling.VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_AUTO,
			},
		},
	})

	scaler.AddDependency(deployment)

	return scaler

}

func v2ScalingPolicies(template *VerticalScalingTemplate) *autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicy {

	minAllowed := map[string]autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed{}
	maxAllowed := map[string]autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed{}

	if template.MinCpuMillis != 0 {
		amount := fmt.Sprintf("%dm", template.MinCpuMillis)
		minCpu := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(amount))
		minAllowed["cpu"] = minCpu
	}

	if template.MaxCpuMillis != 0 {
		amount := fmt.Sprintf("%dm", template.MinCpuMillis)
		maxCpu := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(amount))
		maxAllowed["cpu"] = maxCpu
	}

	if template.MinMemoryMi != 0 {
		amount := fmt.Sprintf("%dMi", template.MinCpuMillis)
		minMem := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(amount))
		minAllowed["memory"] = minMem
	}

	if template.MaxMemoryMi != 0 {
		amount := fmt.Sprintf("%dMi", template.MaxCpuMillis)
		maxMem := autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(amount))
		maxAllowed["memory"] = maxMem
	}

	return &autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicy{
		ContainerPolicies: &[]*autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPolicies{
			{
				MaxAllowed: &maxAllowed,
				MinAllowed: &minAllowed,
				Mode:       autoscaling.VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_AUTO,
			},
		},
	}
}

func newVerticalScaler(scope constructs.Construct, deployment kplus.Deployment, id string, template *VerticalScalingTemplate) autoscaling.VerticalPodAutoscaler {
	name := fmt.Sprintf("%s-autoscaler", id)

	target := autoscaling.VerticalPodAutoscalerSpecTargetRef{
		Kind:       deployment.Kind(),
		Name:       deployment.Name(),
		ApiVersion: deployment.ApiVersion(),
	}

	minReplicas := jsii.Number(1)
	if template.MinReplicas != 0 {
		minReplicas = jsii.Number(template.MinReplicas)
	}

	policy := verticalScalingPolicies(template)

	scaler := autoscaling.NewVerticalPodAutoscaler(scope, jsii.String(name), &autoscaling.VerticalPodAutoscalerProps{
		Metadata: &cdk8s.ApiObjectMetadata{Name: jsii.String(name)},
		Spec: &autoscaling.VerticalPodAutoscalerSpec{
			TargetRef:      &target,
			Recommenders:   &[]*autoscaling.VerticalPodAutoscalerSpecRecommenders{},
			ResourcePolicy: policy,
			UpdatePolicy:   &autoscaling.VerticalPodAutoscalerSpecUpdatePolicy{MinReplicas: minReplicas},
		},
	})

	scaler.AddDependency(deployment)
	return scaler
}

func verticalScalingPolicies(template *VerticalScalingTemplate) *autoscaling.VerticalPodAutoscalerSpecResourcePolicy {

	controlledResources := []*string{}
	if template.ControlCpu {
		controlledResources = append(controlledResources, jsii.String("ResourceCPU"))
	}

	if template.ControlMem {
		controlledResources = append(controlledResources, jsii.String("ResourceMemory"))
	}

	controlledValues := autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_ONLY
	if template.ControlLimits {
		controlledValues = autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_AND_LIMITS
	}

	minAllowed := map[string]autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed{}
	maxAllowed := map[string]autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed{}

	if template.MinCpuMillis != 0 {
		amount := fmt.Sprintf("%dm", template.MinCpuMillis)
		minCpu := autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(amount))
		minAllowed["cpu"] = minCpu
	}

	if template.MaxCpuMillis != 0 {
		amount := fmt.Sprintf("%dm", template.MaxCpuMillis)
		//autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_AND_LIMITS
		//autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesControlledValues_REQUESTS_ONLY
		maxCpu := autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(amount))
		maxAllowed["cpu"] = maxCpu
	}

	if template.MinMemoryMi != 0 {
		amount := fmt.Sprintf("%dMi", template.MinMemoryMi)
		minMem := autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMinAllowed_FromString(jsii.String(amount))
		minAllowed["memory"] = minMem
	}

	if template.MaxMemoryMi != 0 {
		amount := fmt.Sprintf("%dMi", template.MaxMemoryMi)
		maxMem := autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromString(jsii.String(amount))
		maxAllowed["memory"] = maxMem
	}

	return &autoscaling.VerticalPodAutoscalerSpecResourcePolicy{
		ContainerPolicies: &[]*autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPolicies{
			&autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPolicies{
				//ContainerName:       container.Name(),
				ControlledResources: &controlledResources,
				ControlledValues:    controlledValues,
				MaxAllowed:          &maxAllowed,
				MinAllowed:          &minAllowed,
				Mode:                autoscaling.VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode_AUTO,
			},
		},
	}
}

func AddDefaultScaling(scope constructs.Construct, deployment kplus.Deployment, id string, scaling DefaultAutoScaling) constructs.Construct {
	if scaling.DefaultScaling == nil {
		return nil
	}
	if scaling.DefaultScaling.ScalingType() == ScalingTypeHorizontal {

		if horizontal, ok := scaling.DefaultScaling.(*DefaultHorizontalScaling); ok {
			return addDefaultHorizontalScaling(scope, deployment, id, horizontal)
		}
		return nil
	}

	if scaling.DefaultScaling.ScalingType() == ScalingTypeVertical {
		if vertical, ok := scaling.DefaultScaling.(*DefaultVerticalScaling); ok {
			return addDefaultVerticalScaling(scope, deployment, id, vertical)
		}
		return nil
	}

	return nil
}

func addDefaultHorizontalScaling(scope constructs.Construct, deployment kplus.Deployment, id string, scaling *DefaultHorizontalScaling) kplus.HorizontalPodAutoscaler {
	template := &HorizontalScalingTemplate{
		Type: "horizontal",
		CPU: &HorizontalResourceScaling{
			Utilization: 80,
		},
		Memory: &HorizontalResourceScaling{
			Utilization: 80,
		},
		MinReplicas: int32(scaling.MinReplicas),
		MaxReplicas: int32(scaling.MaxReplicas),
	}

	return addHorizontalAutoScaler(scope, deployment, id, template)
}

func addDefaultVerticalScaling(scope constructs.Construct, deployment kplus.Deployment, id string, scaling *DefaultVerticalScaling) autoscaling.VerticalPodAutoscaler {
	minCpu := 100
	maxCpu := 300
	minMem := 75
	maxMem := 500

	if scaling.MinCpuMillis != 0 {
		minCpu = scaling.MinCpuMillis
	}

	if scaling.MaxCpuMillis != 0 {
		maxCpu = scaling.MaxCpuMillis
	}

	if scaling.MinMemoryMi != 0 {
		minMem = scaling.MinMemoryMi
	}
	if scaling.MaxMemoryMi != 0 {
		maxMem = scaling.MaxMemoryMi
	}

	template := &VerticalScalingTemplate{
		Type:          "vertical",
		Version:       2,
		MinCpuMillis:  int32(minCpu),
		MaxCpuMillis:  int32(maxCpu),
		MinMemoryMi:   int32(minMem),
		MaxMemoryMi:   int32(maxMem),
		ControlMem:    true,
		ControlCpu:    true,
		ControlLimits: scaling.ControlLimits,
		MinReplicas:   int32(*deployment.Replicas()),
	}

	return newVerticalScalerV2(scope, deployment, id, template)
}
