package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
//
// Default: 5s.
type TraefikServiceSpecMirroringMirrorsHealthCheckTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringMirrorsHealthCheckTimeout
type jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func TraefikServiceSpecMirroringMirrorsHealthCheckTimeout_FromNumber(value *float64) TraefikServiceSpecMirroringMirrorsHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsHealthCheckTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringMirrorsHealthCheckTimeout_FromString(value *string) TraefikServiceSpecMirroringMirrorsHealthCheckTimeout {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsHealthCheckTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

