package autoscalingk8sio


// Controls when autoscaler applies changes to the pod resources.
//
// The default is 'Auto'.
type VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode string

const (
	// Off.
	VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_OFF VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode = "OFF"
	// Initial.
	VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_INITIAL VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode = "INITIAL"
	// Recreate.
	VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_RECREATE VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode = "RECREATE"
	// Auto.
	VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode_AUTO VerticalPodAutoscalerV1Beta2SpecUpdatePolicyUpdateMode = "AUTO"
)

