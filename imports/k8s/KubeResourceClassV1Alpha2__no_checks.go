//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClassV1Alpha2_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceClassV1Alpha2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClassV1Alpha2_ManifestParameters(props *KubeResourceClassV1Alpha2Props) error {
	return nil
}

func validateKubeResourceClassV1Alpha2_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClassV1Alpha2Parameters(scope constructs.Construct, id *string, props *KubeResourceClassV1Alpha2Props) error {
	return nil
}
