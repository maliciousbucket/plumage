package template

type MonitoringTemplate struct {
	ScrapeConfig        *ScrapeConfigTemplate  `yaml:"scrape"`
	MetricsExportConfig *MetricsExportTemplate `yaml:"metricsExport"`
	LogsExportConfig    *LogsExportTemplate    `yaml:"logsExport"`
	TracesExportConfig  *TracesExportTemplate  `yaml:"tracesExport"`
}

type ScrapeConfigTemplate struct {
	MetricsPort int    `yaml:"metricsPort"`
	MetricsPath string `yaml:"metricsPath"`
}

type MetricsExportTemplate struct {
	Prometheus *PrometheusExportConfig  `yaml:"prometheus,omitempty"`
	Otlp       *OtlpMetricsExportConfig `yaml:"otlp,omitempty"`
	//Prom / OTLP
}

type LogsExportTemplate struct {
	LogsProtocol    string `yaml:"logsProtocol,omitempty"`
	LogsEndpointKey string `yaml:"logsEndpointKey,omitempty"`
	LogsPortKey     string `yaml:"logsPortKey,omitempty"`
	// Loki / OTLP
}

type TracesExportTemplate struct {
	TracesExporter    string `yaml:"tracesExporter,omitempty"`
	TracesProtocol    string `yaml:"tracesProtocol,omitempty"`
	TracesEndpointKey string `yaml:"tracesEndpointKey,omitempty"`
	TracesPortKey     string `yaml:"tracesPortKey,omitempty"`
}

type MetricsExportConfig interface {
	EndpointKey() string
	PortKey() string
	Protocol() string
	Exporter() string
}

type PrometheusExportConfig struct {
	PrometheusEndpointKey string `yaml:"prometheusEndpointKey,omitempty"`
	PrometheusPortKey     string `yaml:"prometheusPortKey,omitempty"`
	Composite             bool   `yaml:"composite,omitempty"`
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
	OtlpEndpointKey string `yaml:"otlpEndpointKey,omitempty"`
	OtlpPortKey     string `yaml:"otlpPortKey,omitempty"`
	OtlpExporter    string `yaml:"otlpExporter,omitempty"`
	OtlpExporterKey string `yaml:"otlpExporterKey,omitempty"`
	ExportProtocol  string `yaml:"exportProtocol,omitempty"`
}

func (c *OtlpMetricsExportConfig) EndpointKey() string {
	return c.OtlpEndpointKey
}

func (c *OtlpMetricsExportConfig) PortKey() string {
	return c.OtlpPortKey
}

func (c *OtlpMetricsExportConfig) Protocol() string {
	return c.ExportProtocol
}

func (c *OtlpMetricsExportConfig) Exporter() string {
	return "otlp"
}
