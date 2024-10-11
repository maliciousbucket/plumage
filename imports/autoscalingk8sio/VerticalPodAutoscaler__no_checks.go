//go:build no_runtime_type_checking

package autoscalingk8sio

// Building without runtime type checking enabled, so all the below just return nil

func validateVerticalPodAutoscaler_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateVerticalPodAutoscaler_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVerticalPodAutoscaler_ManifestParameters(props *VerticalPodAutoscalerProps) error {
	return nil
}

func validateVerticalPodAutoscaler_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewVerticalPodAutoscalerParameters(scope constructs.Construct, id *string, props *VerticalPodAutoscalerProps) error {
	return nil
}
