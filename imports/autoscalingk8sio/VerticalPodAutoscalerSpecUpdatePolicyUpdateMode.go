package autoscalingk8sio


// Controls when autoscaler applies changes to the pod resources.
//
// The default is 'Auto'.
type VerticalPodAutoscalerSpecUpdatePolicyUpdateMode string

const (
	// Off.
	VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_OFF VerticalPodAutoscalerSpecUpdatePolicyUpdateMode = "OFF"
	// Initial.
	VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_INITIAL VerticalPodAutoscalerSpecUpdatePolicyUpdateMode = "INITIAL"
	// Recreate.
	VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_RECREATE VerticalPodAutoscalerSpecUpdatePolicyUpdateMode = "RECREATE"
	// Auto.
	VerticalPodAutoscalerSpecUpdatePolicyUpdateMode_AUTO VerticalPodAutoscalerSpecUpdatePolicyUpdateMode = "AUTO"
)

