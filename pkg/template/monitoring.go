package template

type MonitoringTemplate struct {
}

type ScrapeConfigTemplate struct {
	MetricsPort int
	MetricsPath string
}

type MetricsExportTemplate struct {
	//Prom / OTLP
}

type LogsExportTemplate struct {
	// Loki / OTLP
}

type MetricsExportConfig interface {
	EndpointKey() string
	PortKey() string
	Protocol() string
	Exporter() string
}

type PrometheusExportConfig struct {
	PrometheusEndpointKey string
	PrometheusPortKey     string
	composite             bool
}

func (c *PrometheusExportConfig) EndpointKey() string {
	return c.PrometheusEndpointKey
}

func (c *PrometheusExportConfig) PortKey() string {
	return c.PrometheusPortKey
}

func (c *PrometheusExportConfig) Protocol() string {
	return "http/protobuf"
}

func (c *PrometheusExportConfig) Exporter() string {
	return "prometheus"
}

type OtlpMetricsExportConfig struct {
	OtlpEndpointKey string
	OtlpPortKey     string
	OtlpExporter    string
	OtlpExporterKey string
	ExportProtocol  string
}
