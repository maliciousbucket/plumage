package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
//
// Default: 5s.
type TraefikServiceSpecWeightedServicesHealthCheckTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecWeightedServicesHealthCheckTimeout
type jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func TraefikServiceSpecWeightedServicesHealthCheckTimeout_FromNumber(value *float64) TraefikServiceSpecWeightedServicesHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesHealthCheckTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecWeightedServicesHealthCheckTimeout_FromString(value *string) TraefikServiceSpecWeightedServicesHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesHealthCheckTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
