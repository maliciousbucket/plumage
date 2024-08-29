package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
//
// Default: 5s.
type MiddlewareSpecErrorsServiceHealthCheckTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for MiddlewareSpecErrorsServiceHealthCheckTimeout
type jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func MiddlewareSpecErrorsServiceHealthCheckTimeout_FromNumber(value *float64) MiddlewareSpecErrorsServiceHealthCheckTimeout {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServiceHealthCheckTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServiceHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MiddlewareSpecErrorsServiceHealthCheckTimeout_FromString(value *string) MiddlewareSpecErrorsServiceHealthCheckTimeout {
	_init_.Initialize()

	if err := validateMiddlewareSpecErrorsServiceHealthCheckTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns MiddlewareSpecErrorsServiceHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

