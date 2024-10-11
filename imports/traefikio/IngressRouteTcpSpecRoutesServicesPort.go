package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type IngressRouteTcpSpecRoutesServicesPort interface {
	Value() interface{}
}

// The jsii proxy struct for IngressRouteTcpSpecRoutesServicesPort
type jsiiProxy_IngressRouteTcpSpecRoutesServicesPort struct {
	_ byte // padding
}

func (j *jsiiProxy_IngressRouteTcpSpecRoutesServicesPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func IngressRouteTcpSpecRoutesServicesPort_FromNumber(value *float64) IngressRouteTcpSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteTcpSpecRoutesServicesPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteTcpSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteTcpSpecRoutesServicesPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func IngressRouteTcpSpecRoutesServicesPort_FromString(value *string) IngressRouteTcpSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteTcpSpecRoutesServicesPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteTcpSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteTcpSpecRoutesServicesPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
