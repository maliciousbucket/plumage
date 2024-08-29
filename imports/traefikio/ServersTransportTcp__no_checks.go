//go:build no_runtime_type_checking

package traefikio

// Building without runtime type checking enabled, so all the below just return nil

func validateServersTransportTcp_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateServersTransportTcp_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServersTransportTcp_ManifestParameters(props *ServersTransportTcpProps) error {
	return nil
}

func validateServersTransportTcp_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewServersTransportTcpParameters(scope constructs.Construct, id *string, props *ServersTransportTcpProps) error {
	return nil
}

