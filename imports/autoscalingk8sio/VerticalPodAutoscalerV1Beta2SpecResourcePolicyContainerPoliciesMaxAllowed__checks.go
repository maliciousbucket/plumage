//go:build !no_runtime_type_checking

package autoscalingk8sio

import (
	"fmt"
)

func validateVerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateVerticalPodAutoscalerV1Beta2SpecResourcePolicyContainerPoliciesMaxAllowed_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}
