package template

type ScalingTemplate struct {
	InitialReplicas int
	Resources       *ContainerResourcesTemplate
	Autoscaling     *AutoScalingTemplate
}

type AutoScalingTemplate struct {
	MinReplicas int                `yaml:"minReplicas"`
	MaxReplicas int                `yaml:"maxReplicas"`
	Policy      *AutoScalingPolicy `yaml:"policy"`
}

//Scale policy: Agressive, balanced, conservative
//Stabilisation: Aggressive, conservatice, none

type AutoScalingPolicy struct {
	ScaleOn         []string `json:"scaleOn"` //cpu, memory
	ResourcePolicy  string   `json:"resourcePolicy"`
	ScaleUpPolicy   string   `json:"scaleUpPolicy"`
	ScaleDownPolicy string   `json:"scaleDownPolicy"`
	Stabilisation   string   `json:"stabilisation"`
}

type ContainerResourcesTemplate struct {
	CpuRequestMillis int `yaml:"cpuRequestMillis,omitempty"`
	CpuLimitMillis   int `yaml:"cpuLimitMillis,omitempty"`
	MemoryRequestMb  int `yaml:"memoryRequestMb,omitempty"`
	MemoryLimitMb    int `yaml:"memoryLimitMb,omitempty"`
}

//func asc() {
//	rules := cdk8splus30.ScalingRules{
//		Policies: &[]*cdk8splus30.ScalingPolicy{
//			&cdk8splus30.ScalingPolicy{
//				Replicas: cdk8splus30.Replicas_Absolute(),
//				Duration: nil,
//			},
//			&cdk8splus30.ScalingPolicy{
//				Replicas: nil,
//				Duration: nil,
//			},
//		},
//		StabilizationWindow: nil,
//		LimitStrategy:            cdk8splus30.ScalingStrategy_MAX_CHANGE,
//	}
//
//	t := cdk8splus30.HorizontalPodAutoscalerProps{
//		Metadata:    nil,
//		MaxReplicas: nil,
//		Target:      nil,
//		Metrics: cdk8splus30.Metric_Pods(&cdk8splus30.MetricOptions{
//			Name:          nil,
//			Target:        nil,
//			LabelSelector: nil,
//		}),
//		MinReplicas: nil,
//		ScaleDown:   nil,
//		ScaleUp:     nil,
//	}
//	target := cdk8splus30.MetricTarget_Value()
//	cdk8splus30.Metric_ResourceMemory()
//	cdk8splus30.MetricObjectOptions{
//		Name:          nil,
//		Target:        nil,
//		LabelSelector: nil,
//		Object:        nil,
//	}
//}
