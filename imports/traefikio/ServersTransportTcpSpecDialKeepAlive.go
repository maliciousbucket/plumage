package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"
)

// DialKeepAlive is the interval between keep-alive probes for an active network connection.
//
// If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.
type ServersTransportTcpSpecDialKeepAlive interface {
	Value() interface{}
}

// The jsii proxy struct for ServersTransportTcpSpecDialKeepAlive
type jsiiProxy_ServersTransportTcpSpecDialKeepAlive struct {
	_ byte // padding
}

func (j *jsiiProxy_ServersTransportTcpSpecDialKeepAlive) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func ServersTransportTcpSpecDialKeepAlive_FromNumber(value *float64) ServersTransportTcpSpecDialKeepAlive {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecDialKeepAlive_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecDialKeepAlive

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecDialKeepAlive",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ServersTransportTcpSpecDialKeepAlive_FromString(value *string) ServersTransportTcpSpecDialKeepAlive {
	_init_.Initialize()

	if err := validateServersTransportTcpSpecDialKeepAlive_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ServersTransportTcpSpecDialKeepAlive

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcpSpecDialKeepAlive",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}
