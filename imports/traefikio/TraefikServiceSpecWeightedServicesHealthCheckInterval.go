package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Interval defines the frequency of the health check calls.
//
// Default: 30s.
type TraefikServiceSpecWeightedServicesHealthCheckInterval interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecWeightedServicesHealthCheckInterval
type jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func TraefikServiceSpecWeightedServicesHealthCheckInterval_FromNumber(value *float64) TraefikServiceSpecWeightedServicesHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesHealthCheckInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecWeightedServicesHealthCheckInterval_FromString(value *string) TraefikServiceSpecWeightedServicesHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecWeightedServicesHealthCheckInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecWeightedServicesHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

