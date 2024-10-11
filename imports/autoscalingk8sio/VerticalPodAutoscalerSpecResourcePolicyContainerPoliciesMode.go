package autoscalingk8sio

// Whether autoscaler is enabled for the container.
//
// The default is "Auto".
type VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode string

const (
	// Auto.
	VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode_AUTO VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode = "AUTO"
	// Off.
	VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode_OFF VerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMode = "OFF"
)
