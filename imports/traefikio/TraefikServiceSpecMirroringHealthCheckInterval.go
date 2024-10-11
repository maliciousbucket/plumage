package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Interval defines the frequency of the health check calls.
//
// Default: 30s.
type TraefikServiceSpecMirroringHealthCheckInterval interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringHealthCheckInterval
type jsiiProxy_TraefikServiceSpecMirroringHealthCheckInterval struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringHealthCheckInterval) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func TraefikServiceSpecMirroringHealthCheckInterval_FromNumber(value *float64) TraefikServiceSpecMirroringHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringHealthCheckInterval_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringHealthCheckInterval",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringHealthCheckInterval_FromString(value *string) TraefikServiceSpecMirroringHealthCheckInterval {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringHealthCheckInterval_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringHealthCheckInterval

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringHealthCheckInterval",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
