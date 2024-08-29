//go:build no_runtime_type_checking

package traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateMiddlewareSpecErrorsServiceHealthCheckInterval_FromNumberParameters(value *float64) error {
	return nil
}

func validateMiddlewareSpecErrorsServiceHealthCheckInterval_FromStringParameters(value *string) error {
	return nil
}

