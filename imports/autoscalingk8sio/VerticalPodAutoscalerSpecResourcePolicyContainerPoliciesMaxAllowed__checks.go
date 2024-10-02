//go:build !no_runtime_type_checking

package autoscalingk8sio

import (
	"fmt"
)

func validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateVerticalPodAutoscalerSpecResourcePolicyContainerPoliciesMaxAllowed_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

