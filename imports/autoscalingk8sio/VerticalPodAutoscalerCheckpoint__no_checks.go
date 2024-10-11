//go:build no_runtime_type_checking

package autoscalingk8sio

// Building without runtime type checking enabled, so all the below just return nil

func validateVerticalPodAutoscalerCheckpoint_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateVerticalPodAutoscalerCheckpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVerticalPodAutoscalerCheckpoint_ManifestParameters(props *VerticalPodAutoscalerCheckpointProps) error {
	return nil
}

func validateVerticalPodAutoscalerCheckpoint_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewVerticalPodAutoscalerCheckpointParameters(scope constructs.Construct, id *string, props *VerticalPodAutoscalerCheckpointProps) error {
	return nil
}
