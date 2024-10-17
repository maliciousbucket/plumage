package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// Timeout defines the maximum duration Traefik will wait for a health check request before considering the server unhealthy.
//
// Default: 5s.
type IngressRouteSpecRoutesServicesHealthCheckTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for IngressRouteSpecRoutesServicesHealthCheckTimeout
type jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func IngressRouteSpecRoutesServicesHealthCheckTimeout_FromNumber(value *float64) IngressRouteSpecRoutesServicesHealthCheckTimeout {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesHealthCheckTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func IngressRouteSpecRoutesServicesHealthCheckTimeout_FromString(value *string) IngressRouteSpecRoutesServicesHealthCheckTimeout {
	_init_.Initialize()

	if err := validateIngressRouteSpecRoutesServicesHealthCheckTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns IngressRouteSpecRoutesServicesHealthCheckTimeout

	_jsii_.StaticInvoke(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

