package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Port defines the port of a Kubernetes Service.
//
// This can be a reference to a named port.
type TraefikServiceSpecMirroringMirrorsPort interface {
	Value() interface{}
}

// The jsii proxy struct for TraefikServiceSpecMirroringMirrorsPort
type jsiiProxy_TraefikServiceSpecMirroringMirrorsPort struct {
	_ byte // padding
}

func (j *jsiiProxy_TraefikServiceSpecMirroringMirrorsPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func TraefikServiceSpecMirroringMirrorsPort_FromNumber(value *float64) TraefikServiceSpecMirroringMirrorsPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsPort

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func TraefikServiceSpecMirroringMirrorsPort_FromString(value *string) TraefikServiceSpecMirroringMirrorsPort {
	_init_.Initialize()

	if err := validateTraefikServiceSpecMirroringMirrorsPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns TraefikServiceSpecMirroringMirrorsPort

	_jsii_.StaticInvoke(
		"traefikio.TraefikServiceSpecMirroringMirrorsPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
