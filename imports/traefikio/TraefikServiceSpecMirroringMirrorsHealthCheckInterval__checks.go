//go:build !no_runtime_type_checking

package traefikio

import (
	"fmt"
)

func validateTraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateTraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}
