package traefikio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/maliciousbucket/plumage/imports/traefikio/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/imports/traefikio/internal"
)

// ServersTransportTCP is the CRD implementation of a TCPServersTransport.
//
// If no tcpServersTransport is specified, a default one named default@internal will be used.
// The default@internal tcpServersTransport can be configured in the static configuration.
// More info: https://doc.traefik.io/traefik/v3.1/routing/services/#serverstransport_3
type ServersTransportTcp interface {
	cdk8s.ApiObject
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The object's API version (e.g. `authorization.k8s.io/v1`).
	ApiVersion() *string
	// The chart in which this object is defined.
	Chart() cdk8s.Chart
	// The object kind.
	Kind() *string
	// Metadata associated with this API object.
	Metadata() cdk8s.ApiObjectMetadataDefinition
	// The name of the API object.
	//
	// If a name is specified in `metadata.name` this will be the name returned.
	// Otherwise, a name will be generated by calling
	// `Chart.of(this).generatedObjectName(this)`, which by default uses the
	// construct path to generate a DNS-compatible name for the resource.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// Create a dependency between this ApiObject and other constructs.
	//
	// These can be other ApiObjects, Charts, or custom.
	AddDependency(dependencies ...constructs.IConstruct)
	// Applies a set of RFC-6902 JSON-Patch operations to the manifest synthesized for this API object.
	//
	// Example:
	//     kubePod.addJsonPatch(JsonPatch.replace('/spec/enableServiceLinks', true));
	//
	AddJsonPatch(ops ...cdk8s.JsonPatch)
	// Renders the object to Kubernetes JSON.
	ToJson() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ServersTransportTcp
type jsiiProxy_ServersTransportTcp struct {
	internal.Type__cdk8sApiObject
}

func (j *jsiiProxy_ServersTransportTcp) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) Chart() cdk8s.Chart {
	var returns cdk8s.Chart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) Metadata() cdk8s.ApiObjectMetadataDefinition {
	var returns cdk8s.ApiObjectMetadataDefinition
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServersTransportTcp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

// Defines a "ServersTransportTCP" API object.
func NewServersTransportTcp(scope constructs.Construct, id *string, props *ServersTransportTcpProps) ServersTransportTcp {
	_init_.Initialize()

	if err := validateNewServersTransportTcpParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServersTransportTcp{}

	_jsii_.Create(
		"traefikio.ServersTransportTcp",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Defines a "ServersTransportTCP" API object.
func NewServersTransportTcp_Override(s ServersTransportTcp, scope constructs.Construct, id *string, props *ServersTransportTcpProps) {
	_init_.Initialize()

	_jsii_.Create(
		"traefikio.ServersTransportTcp",
		[]interface{}{scope, id, props},
		s,
	)
}

// Return whether the given object is an `ApiObject`.
//
// We do attribute detection since we can't reliably use 'instanceof'.
func ServersTransportTcp_IsApiObject(o interface{}) *bool {
	_init_.Initialize()

	if err := validateServersTransportTcp_IsApiObjectParameters(o); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcp",
		"isApiObject",
		[]interface{}{o},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ServersTransportTcp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateServersTransportTcp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Renders a Kubernetes manifest for "ServersTransportTCP".
//
// This can be used to inline resource manifests inside other objects (e.g. as templates).
func ServersTransportTcp_Manifest(props *ServersTransportTcpProps) interface{} {
	_init_.Initialize()

	if err := validateServersTransportTcp_ManifestParameters(props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcp",
		"manifest",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Returns the `ApiObject` named `Resource` which is a child of the given construct.
//
// If `c` is an `ApiObject`, it is returned directly. Throws an
// exception if the construct does not have a child named `Default` _or_ if
// this child is not an `ApiObject`.
func ServersTransportTcp_Of(c constructs.IConstruct) cdk8s.ApiObject {
	_init_.Initialize()

	if err := validateServersTransportTcp_OfParameters(c); err != nil {
		panic(err)
	}
	var returns cdk8s.ApiObject

	_jsii_.StaticInvoke(
		"traefikio.ServersTransportTcp",
		"of",
		[]interface{}{c},
		&returns,
	)

	return returns
}

func ServersTransportTcp_GVK() *cdk8s.GroupVersionKind {
	_init_.Initialize()
	var returns *cdk8s.GroupVersionKind
	_jsii_.StaticGet(
		"traefikio.ServersTransportTcp",
		"GVK",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_ServersTransportTcp) AddDependency(dependencies ...constructs.IConstruct) {
	args := []interface{}{}
	for _, a := range dependencies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addDependency",
		args,
	)
}

func (s *jsiiProxy_ServersTransportTcp) AddJsonPatch(ops ...cdk8s.JsonPatch) {
	args := []interface{}{}
	for _, a := range ops {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addJsonPatch",
		args,
	)
}

func (s *jsiiProxy_ServersTransportTcp) ToJson() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServersTransportTcp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
