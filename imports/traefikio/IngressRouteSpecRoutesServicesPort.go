package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type IngressRouteSpecRoutesServicesPort interface {
	Value() interface{}
}

// The jsii proxy struct for IngressRouteSpecRoutesServicesPort
type jsiiProxy_IngressRouteSpecRoutesServicesPort struct {
	_ byte // padding
}

func (j *jsiiProxy_IngressRouteSpecRoutesServicesPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func IngressRouteSpecRoutesServicesPort_FromNumber(value *float64) IngressRouteSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func IngressRouteSpecRoutesServicesPort_FromString(value *string) IngressRouteSpecRoutesServicesPort {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesPort

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

