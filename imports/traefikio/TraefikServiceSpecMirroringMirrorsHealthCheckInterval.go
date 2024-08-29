package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Interval defines the frequency of the health check calls.
//
// Default: 30s.
type TraefikServiceSpecMirroringMirrorsHealthCheckInterval interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringMirrorsHealthCheckInterval
type jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func TraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromNumber(value *float64) TraefikServiceSpecMirroringMirrorsHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromString(value *string) TraefikServiceSpecMirroringMirrorsHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsHealthCheckInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

