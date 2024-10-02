//go:build no_runtime_type_checking

package autoscalingk8sio

// Building without runtime type checking enabled, so all the below just return nil

func validateVerticalPodAutoscalerV1Beta2_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateVerticalPodAutoscalerV1Beta2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVerticalPodAutoscalerV1Beta2_ManifestParameters(props *VerticalPodAutoscalerV1Beta2Props) error {
	return nil
}

func validateVerticalPodAutoscalerV1Beta2_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewVerticalPodAutoscalerV1Beta2Parameters(scope constructs.Construct, id *string, props *VerticalPodAutoscalerV1Beta2Props) error {
	return nil
}

