package autoscalingk8sio

// Describes the rules on how changes are applied to the pods.
//
// If not specified, all fields in the `PodUpdatePolicy` are set to their default values.
type VerticalPodAutoscalerV1Beta2SpecUpdatePolicy struct {
	// Controls when autoscaler applies changes to the pod resources.
	//
	// The default is 'Auto'.
	UpdateMode VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode `field:"optional" json:"updateMode" yaml:"updateMode"`
}
