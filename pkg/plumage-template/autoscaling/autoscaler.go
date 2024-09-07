package autoscaling

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
)

const (
	DeploymentKind = "Deployment"
	AppsV1         = "apps/v1"
	AppLabel       = "app"
)

type AutoScalerProps struct {
	Name      string
	Namespace string
	Scaling   *plumagetemplate.ScalingConfig
}

func NewHorizontalAutoscaler(scope constructs.Construct, id string, props *AutoScalerProps) k8s.KubeHorizontalPodAutoscalerV2 {
	metadata := autoSclaerMeta(props.Namespace, props.Name)
	metrics := autoScalerMetrics(props.Scaling)
	minReplicas := 1
	if props.Scaling.MinReplicas != 0 {
		minReplicas = props.Scaling.MinReplicas
	}
	maxReplicas := 2
	if props.Scaling.MaxReplicas != 0 {
		maxReplicas = props.Scaling.MaxReplicas
	}
	return k8s.NewKubeHorizontalPodAutoscalerV2(scope, jsii.String(id), &k8s.KubeHorizontalPodAutoscalerV2Props{
		Metadata: metadata,
		Spec: &k8s.HorizontalPodAutoscalerSpecV2{
			MaxReplicas: jsii.Number(maxReplicas),
			ScaleTargetRef: &k8s.CrossVersionObjectReferenceV2{
				Kind:       jsii.String(DeploymentKind),
				Name:       jsii.String(props.Name),
				ApiVersion: jsii.String(AppsV1),
			},
			//TODO Behaviour
			Behavior:    nil,
			Metrics:     metrics,
			MinReplicas: jsii.Number(minReplicas),
		},
	})
}

func autoSclaerMeta(ns string, appLabel string) *k8s.ObjectMeta {
	labels := autoScalerLabels(appLabel)
	return &k8s.ObjectMeta{
		Name:      jsii.String(appLabel),
		Namespace: jsii.String(ns),
		Labels:    labels,
	}
}

func autoScalerLabels(appLabel string) *map[string]*string {
	labels := map[string]*string{
		AppLabel: jsii.String(appLabel),
	}
	return &labels
}

func autoScalerMetrics(config *plumagetemplate.ScalingConfig) *[]*k8s.MetricSpecV2 {
	if config == nil {
		return nil
	}
	var metrics []*k8s.MetricSpecV2

	if config.TargetCpuPercent != 0 {
		cpuPercent := jsii.Number(config.TargetCpuPercent)
		cpuUtilization := k8s.MetricSpecV2{
			Type: jsii.String("Resource"),
			Resource: &k8s.ResourceMetricSourceV2{
				Name: jsii.String("cpu"),
				Target: &k8s.MetricTargetV2{
					Type:               jsii.String("Utilization"),
					AverageUtilization: cpuPercent,
				},
			},
		}
		metrics = append(metrics, &cpuUtilization)
	}

	if config.TargetCpuAmount != 0 {
		cpuQuantity := k8s.Quantity_FromNumber(jsii.Number(config.TargetCpuAmount))
		cpuAmount := k8s.MetricSpecV2{
			Type: jsii.String("Resource"),
			Resource: &k8s.ResourceMetricSourceV2{
				Name: jsii.String("cpu"),
				Target: &k8s.MetricTargetV2{
					Type:         jsii.String("AverageValue"),
					AverageValue: cpuQuantity,
				},
			},
		}
		metrics = append(metrics, &cpuAmount)
	}

	if config.TargetMemoryPercent != 0 {
		memoryPercent := jsii.Number(config.TargetMemoryPercent)
		memoryUtilization := k8s.MetricSpecV2{
			Type: jsii.String("Resource"),
			Resource: &k8s.ResourceMetricSourceV2{
				Name: jsii.String("memory"),
				Target: &k8s.MetricTargetV2{
					Type:               jsii.String("Utilization"),
					AverageUtilization: memoryPercent,
				},
			},
		}
		metrics = append(metrics, &memoryUtilization)
	}

	if config.TargetMemoryAmount != 0 {
		memoryQuantity := k8s.Quantity_FromNumber(jsii.Number(config.TargetMemoryAmount))
		memoryAmount := k8s.MetricSpecV2{
			Type: jsii.String("Resource"),
			Resource: &k8s.ResourceMetricSourceV2{
				Name: jsii.String("memory"),
				Target: &k8s.MetricTargetV2{
					Type:         jsii.String("AverageValue"),
					AverageValue: memoryQuantity,
				},
			},
		}
		metrics = append(metrics, &memoryAmount)
	}
	return &metrics

}

func DefaultAutoScaler(scope constructs.Construct, id string, ns string, appLabel string) k8s.KubeHorizontalPodAutoscalerV2 {
	metada := autoSclaerMeta(ns, appLabel)

	return k8s.NewKubeHorizontalPodAutoscalerV2(scope, jsii.String(id), &k8s.KubeHorizontalPodAutoscalerV2Props{
		Metadata: metada,
		Spec: &k8s.HorizontalPodAutoscalerSpecV2{
			MaxReplicas: jsii.Number(3),
			ScaleTargetRef: &k8s.CrossVersionObjectReferenceV2{
				Kind:       jsii.String(DeploymentKind),
				Name:       jsii.String(appLabel),
				ApiVersion: jsii.String("apps/v1"),
			},
			Behavior: nil,
			Metrics: &[]*k8s.MetricSpecV2{
				&k8s.MetricSpecV2{
					Type: jsii.String("Resource"),
					Resource: &k8s.ResourceMetricSourceV2{
						Name: jsii.String("cpu"),
						Target: &k8s.MetricTargetV2{
							Type:               jsii.String("Utilization"),
							AverageUtilization: jsii.Number(70),
						},
					},
				},
				&k8s.MetricSpecV2{
					Type: jsii.String("Resource"),
					Resource: &k8s.ResourceMetricSourceV2{
						Name: jsii.String("memory"),
						Target: &k8s.MetricTargetV2{
							Type:               jsii.String("Utilization"),
							AverageUtilization: jsii.Number(70),
						},
					},
				},
			},
			MinReplicas: jsii.Number(1),
		},
	})
}
