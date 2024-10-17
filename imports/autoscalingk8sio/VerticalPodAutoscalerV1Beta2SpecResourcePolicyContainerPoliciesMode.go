package autoscalingk8sio


// Whether autoscaler is enabled for the container.
//
// The default is "Auto".
type VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode string

const (
	// Auto.
	VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_AUTO VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode = "AUTO"
	// Off.
	VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode_OFF VerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMode = "OFF"
)

