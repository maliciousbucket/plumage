package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// DialTimeout is the amount of time to wait until a connection to a backend server can be established.
type ServersTransportTcpSpecDialTimeout interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportTcpSpecDialTimeout
type jsiiProxy_ServersTransportTcpSpecDialTimeout struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportTcpSpecDialTimeout) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ServersTransportTcpSpecDialTimeout_FromNumber(value *float64) ServersTransportTcpSpecDialTimeout {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecDialTimeout_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecDialTimeout

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecDialTimeout",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportTcpSpecDialTimeout_FromString(value *string) ServersTransportTcpSpecDialTimeout {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecDialTimeout_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecDialTimeout

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecDialTimeout",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

