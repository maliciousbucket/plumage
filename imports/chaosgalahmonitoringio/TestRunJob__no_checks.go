//go:build no_runtime_type_checking

package chaosgalahmonitoringio

// Building without runtime type checking enabled, so all the below just return nil

func validateTestRunJob_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateTestRunJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTestRunJob_ManifestParameters(props *TestRunJobProps) error {
	return nil
}

func validateTestRunJob_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewTestRunJobParameters(scope constructs.Construct, id *string, props *TestRunJobProps) error {
	return nil
}

