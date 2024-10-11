//go:build no_runtime_type_checking

package autoscalingk8sio

// Building without runtime type checking enabled, so all the below just return nil

func validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromNumberParameters(value *float64) error {
	return nil
}

func validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromStringParameters(value *string) error {
	return nil
}
