package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
//
// Default: 5s.
type TraefikServiceSpecMirroringHealthCheckTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringHealthCheckTimeout
type jsiiProxy_TraefikServiceSpecMirroringHealthCheckTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringHealthCheckTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func TraefikServiceSpecMirroringHealthCheckTimeout_FromNumber(value *float64) TraefikServiceSpecMirroringHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringHealthCheckTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringHealthCheckTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringHealthCheckTimeout_FromString(value *string) TraefikServiceSpecMirroringHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringHealthCheckTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringHealthCheckTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
