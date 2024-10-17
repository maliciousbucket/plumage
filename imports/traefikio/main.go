// traefikio
package traefikio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"traefikio.IngressRoute",
		reflect.TypeOf((*IngressRoute)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressRoute{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteProps",
		reflect.TypeOf((*IngressRouteProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpec",
		reflect.TypeOf((*IngressRouteSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutes",
		reflect.TypeOf((*IngressRouteSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikio.IngressRouteSpecRoutesKind",
		reflect.TypeOf((*IngressRouteSpecRoutesKind)(nil)).Elem(),
		map[string]interface{}{
			"RULE": IngressRouteSpecRoutesKind_RULE,
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesMiddlewares",
		reflect.TypeOf((*IngressRouteSpecRoutesMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesServices",
		reflect.TypeOf((*IngressRouteSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheck",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckInterval",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesHealthCheckInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckInterval{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteSpecRoutesServicesHealthCheckTimeout",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesHealthCheckTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteSpecRoutesServicesHealthCheckTimeout{}
		},
	)
	_jsii_.RegisterEnum(
		"traefikio.IngressRouteSpecRoutesServicesKind",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": IngressRouteSpecRoutesServicesKind_SERVICE,
			"TRAEFIK_SERVICE": IngressRouteSpecRoutesServicesKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteSpecRoutesServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesServicesResponseForwarding",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesServicesSticky",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecRoutesServicesStickyCookie",
		reflect.TypeOf((*IngressRouteSpecRoutesServicesStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecTls",
		reflect.TypeOf((*IngressRouteSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecTlsDomains",
		reflect.TypeOf((*IngressRouteSpecTlsDomains)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecTlsOptions",
		reflect.TypeOf((*IngressRouteSpecTlsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteSpecTlsStore",
		reflect.TypeOf((*IngressRouteSpecTlsStore)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteTcp",
		reflect.TypeOf((*IngressRouteTcp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressRouteTcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpProps",
		reflect.TypeOf((*IngressRouteTcpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpec",
		reflect.TypeOf((*IngressRouteTcpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecRoutes",
		reflect.TypeOf((*IngressRouteTcpSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecRoutesMiddlewares",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecRoutesServices",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteTcpSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteTcpSpecRoutesServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecRoutesServicesProxyProtocol",
		reflect.TypeOf((*IngressRouteTcpSpecRoutesServicesProxyProtocol)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecTls",
		reflect.TypeOf((*IngressRouteTcpSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecTlsDomains",
		reflect.TypeOf((*IngressRouteTcpSpecTlsDomains)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecTlsOptions",
		reflect.TypeOf((*IngressRouteTcpSpecTlsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteTcpSpecTlsStore",
		reflect.TypeOf((*IngressRouteTcpSpecTlsStore)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteUdp",
		reflect.TypeOf((*IngressRouteUdp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IngressRouteUdp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteUdpProps",
		reflect.TypeOf((*IngressRouteUdpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteUdpSpec",
		reflect.TypeOf((*IngressRouteUdpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteUdpSpecRoutes",
		reflect.TypeOf((*IngressRouteUdpSpecRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.IngressRouteUdpSpecRoutesServices",
		reflect.TypeOf((*IngressRouteUdpSpecRoutesServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.IngressRouteUdpSpecRoutesServicesPort",
		reflect.TypeOf((*IngressRouteUdpSpecRoutesServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_IngressRouteUdpSpecRoutesServicesPort{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.Middleware",
		reflect.TypeOf((*Middleware)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Middleware{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareProps",
		reflect.TypeOf((*MiddlewareProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpec",
		reflect.TypeOf((*MiddlewareSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecAddPrefix",
		reflect.TypeOf((*MiddlewareSpecAddPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecBasicAuth",
		reflect.TypeOf((*MiddlewareSpecBasicAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecBuffering",
		reflect.TypeOf((*MiddlewareSpecBuffering)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecChain",
		reflect.TypeOf((*MiddlewareSpecChain)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecChainMiddlewares",
		reflect.TypeOf((*MiddlewareSpecChainMiddlewares)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecCircuitBreaker",
		reflect.TypeOf((*MiddlewareSpecCircuitBreaker)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecCircuitBreakerCheckPeriod",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerCheckPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerCheckPeriod{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecCircuitBreakerFallbackDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerFallbackDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerFallbackDuration{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecCircuitBreakerRecoveryDuration",
		reflect.TypeOf((*MiddlewareSpecCircuitBreakerRecoveryDuration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecCircuitBreakerRecoveryDuration{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecCompress",
		reflect.TypeOf((*MiddlewareSpecCompress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecContentType",
		reflect.TypeOf((*MiddlewareSpecContentType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecDigestAuth",
		reflect.TypeOf((*MiddlewareSpecDigestAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrors",
		reflect.TypeOf((*MiddlewareSpecErrors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrorsService",
		reflect.TypeOf((*MiddlewareSpecErrorsService)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheck",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckInterval",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceHealthCheckInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckInterval{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecErrorsServiceHealthCheckTimeout",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceHealthCheckTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecErrorsServiceHealthCheckTimeout{}
		},
	)
	_jsii_.RegisterEnum(
		"traefikio.MiddlewareSpecErrorsServiceKind",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": MiddlewareSpecErrorsServiceKind_SERVICE,
			"TRAEFIK_SERVICE": MiddlewareSpecErrorsServiceKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecErrorsServicePort",
		reflect.TypeOf((*MiddlewareSpecErrorsServicePort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecErrorsServicePort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrorsServiceResponseForwarding",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrorsServiceSticky",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecErrorsServiceStickyCookie",
		reflect.TypeOf((*MiddlewareSpecErrorsServiceStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecForwardAuth",
		reflect.TypeOf((*MiddlewareSpecForwardAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecForwardAuthTls",
		reflect.TypeOf((*MiddlewareSpecForwardAuthTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecGrpcWeb",
		reflect.TypeOf((*MiddlewareSpecGrpcWeb)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecHeaders",
		reflect.TypeOf((*MiddlewareSpecHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecInFlightReq",
		reflect.TypeOf((*MiddlewareSpecInFlightReq)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecInFlightReqSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecInFlightReqSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecInFlightReqSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecIpAllowList",
		reflect.TypeOf((*MiddlewareSpecIpAllowList)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecIpAllowListIpStrategy",
		reflect.TypeOf((*MiddlewareSpecIpAllowListIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecIpWhiteList",
		reflect.TypeOf((*MiddlewareSpecIpWhiteList)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecIpWhiteListIpStrategy",
		reflect.TypeOf((*MiddlewareSpecIpWhiteListIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecPassTlsClientCert",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecPassTlsClientCertInfo",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecPassTlsClientCertInfoIssuer",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecPassTlsClientCertInfoSubject",
		reflect.TypeOf((*MiddlewareSpecPassTlsClientCertInfoSubject)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRateLimit",
		reflect.TypeOf((*MiddlewareSpecRateLimit)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecRateLimitPeriod",
		reflect.TypeOf((*MiddlewareSpecRateLimitPeriod)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRateLimitPeriod{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRateLimitSourceCriterion",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRateLimitSourceCriterionIpStrategy",
		reflect.TypeOf((*MiddlewareSpecRateLimitSourceCriterionIpStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRedirectRegex",
		reflect.TypeOf((*MiddlewareSpecRedirectRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRedirectScheme",
		reflect.TypeOf((*MiddlewareSpecRedirectScheme)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecReplacePath",
		reflect.TypeOf((*MiddlewareSpecReplacePath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecReplacePathRegex",
		reflect.TypeOf((*MiddlewareSpecReplacePathRegex)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecRetry",
		reflect.TypeOf((*MiddlewareSpecRetry)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareSpecRetryInitialInterval",
		reflect.TypeOf((*MiddlewareSpecRetryInitialInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_MiddlewareSpecRetryInitialInterval{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecStripPrefix",
		reflect.TypeOf((*MiddlewareSpecStripPrefix)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareSpecStripPrefixRegex",
		reflect.TypeOf((*MiddlewareSpecStripPrefixRegex)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.MiddlewareTcp",
		reflect.TypeOf((*MiddlewareTcp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MiddlewareTcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareTcpProps",
		reflect.TypeOf((*MiddlewareTcpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareTcpSpec",
		reflect.TypeOf((*MiddlewareTcpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareTcpSpecInFlightConn",
		reflect.TypeOf((*MiddlewareTcpSpecInFlightConn)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareTcpSpecIpAllowList",
		reflect.TypeOf((*MiddlewareTcpSpecIpAllowList)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.MiddlewareTcpSpecIpWhiteList",
		reflect.TypeOf((*MiddlewareTcpSpecIpWhiteList)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransport",
		reflect.TypeOf((*ServersTransport)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServersTransport{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportProps",
		reflect.TypeOf((*ServersTransportProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportSpec",
		reflect.TypeOf((*ServersTransportSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportSpecForwardingTimeouts",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportSpecForwardingTimeoutsDialTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsDialTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsDialTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportSpecForwardingTimeoutsIdleConnTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsIdleConnTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsIdleConnTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportSpecForwardingTimeoutsPingTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsPingTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsPingTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportSpecForwardingTimeoutsReadIdleTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsReadIdleTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsReadIdleTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout",
		reflect.TypeOf((*ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportSpecForwardingTimeoutsResponseHeaderTimeout{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportSpecSpiffe",
		reflect.TypeOf((*ServersTransportSpecSpiffe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportTcp",
		reflect.TypeOf((*ServersTransportTcp)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ServersTransportTcp{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportTcpProps",
		reflect.TypeOf((*ServersTransportTcpProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportTcpSpec",
		reflect.TypeOf((*ServersTransportTcpSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportTcpSpecDialKeepAlive",
		reflect.TypeOf((*ServersTransportTcpSpecDialKeepAlive)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportTcpSpecDialKeepAlive{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportTcpSpecDialTimeout",
		reflect.TypeOf((*ServersTransportTcpSpecDialTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportTcpSpecDialTimeout{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.ServersTransportTcpSpecTerminationDelay",
		reflect.TypeOf((*ServersTransportTcpSpecTerminationDelay)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ServersTransportTcpSpecTerminationDelay{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportTcpSpecTls",
		reflect.TypeOf((*ServersTransportTcpSpecTls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.ServersTransportTcpSpecTlsSpiffe",
		reflect.TypeOf((*ServersTransportTcpSpecTlsSpiffe)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TlsOption",
		reflect.TypeOf((*TlsOption)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TlsOption{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsOptionProps",
		reflect.TypeOf((*TlsOptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsOptionSpec",
		reflect.TypeOf((*TlsOptionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsOptionSpecClientAuth",
		reflect.TypeOf((*TlsOptionSpecClientAuth)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"traefikio.TlsOptionSpecClientAuthClientAuthType",
		reflect.TypeOf((*TlsOptionSpecClientAuthClientAuthType)(nil)).Elem(),
		map[string]interface{}{
			"NO_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_NO_CLIENT_CERT,
			"REQUEST_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUEST_CLIENT_CERT,
			"REQUIRE_ANY_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUIRE_ANY_CLIENT_CERT,
			"VERIFY_CLIENT_CERT_IF_GIVEN": TlsOptionSpecClientAuthClientAuthType_VERIFY_CLIENT_CERT_IF_GIVEN,
			"REQUIRE_AND_VERIFY_CLIENT_CERT": TlsOptionSpecClientAuthClientAuthType_REQUIRE_AND_VERIFY_CLIENT_CERT,
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TlsStore",
		reflect.TypeOf((*TlsStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TlsStore{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreProps",
		reflect.TypeOf((*TlsStoreProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreSpec",
		reflect.TypeOf((*TlsStoreSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreSpecCertificates",
		reflect.TypeOf((*TlsStoreSpecCertificates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreSpecDefaultCertificate",
		reflect.TypeOf((*TlsStoreSpecDefaultCertificate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreSpecDefaultGeneratedCert",
		reflect.TypeOf((*TlsStoreSpecDefaultGeneratedCert)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TlsStoreSpecDefaultGeneratedCertDomain",
		reflect.TypeOf((*TlsStoreSpecDefaultGeneratedCertDomain)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikService",
		reflect.TypeOf((*TraefikService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TraefikService{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceProps",
		reflect.TypeOf((*TraefikServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpec",
		reflect.TypeOf((*TraefikServiceSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroring",
		reflect.TypeOf((*TraefikServiceSpecMirroring)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringHealthCheck",
		reflect.TypeOf((*TraefikServiceSpecMirroringHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringHealthCheckInterval",
		reflect.TypeOf((*TraefikServiceSpecMirroringHealthCheckInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringHealthCheckInterval{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringHealthCheckTimeout",
		reflect.TypeOf((*TraefikServiceSpecMirroringHealthCheckTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringHealthCheckTimeout{}
		},
	)
	_jsii_.RegisterEnum(
		"traefikio.TraefikServiceSpecMirroringKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringMirrors",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheck",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckInterval",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsHealthCheckInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckInterval{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringMirrorsHealthCheckTimeout",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsHealthCheckTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringMirrorsHealthCheckTimeout{}
		},
	)
	_jsii_.RegisterEnum(
		"traefikio.TraefikServiceSpecMirroringMirrorsKind",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecMirroringMirrorsKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecMirroringMirrorsKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringMirrorsPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringMirrorsPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringMirrorsResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringMirrorsSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringMirrorsStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringMirrorsStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecMirroringPort",
		reflect.TypeOf((*TraefikServiceSpecMirroringPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecMirroringPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecMirroringResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringSticky",
		reflect.TypeOf((*TraefikServiceSpecMirroringSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecMirroringStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecMirroringStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeighted",
		reflect.TypeOf((*TraefikServiceSpecWeighted)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedServices",
		reflect.TypeOf((*TraefikServiceSpecWeightedServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheck",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckInterval",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesHealthCheckInterval)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckInterval{}
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecWeightedServicesHealthCheckTimeout",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesHealthCheckTimeout)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecWeightedServicesHealthCheckTimeout{}
		},
	)
	_jsii_.RegisterEnum(
		"traefikio.TraefikServiceSpecWeightedServicesKind",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesKind)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE": TraefikServiceSpecWeightedServicesKind_SERVICE,
			"TRAEFIK_SERVICE": TraefikServiceSpecWeightedServicesKind_TRAEFIK_SERVICE,
		},
	)
	_jsii_.RegisterClass(
		"traefikio.TraefikServiceSpecWeightedServicesPort",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_TraefikServiceSpecWeightedServicesPort{}
		},
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedServicesResponseForwarding",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesResponseForwarding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedServicesSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedServicesStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedServicesStickyCookie)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedSticky",
		reflect.TypeOf((*TraefikServiceSpecWeightedSticky)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"traefikio.TraefikServiceSpecWeightedStickyCookie",
		reflect.TypeOf((*TraefikServiceSpecWeightedStickyCookie)(nil)).Elem(),
	)
}
