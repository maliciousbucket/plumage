package config

const (
	GRPCProtocol    Protocol = "grpc"
	HTTPProtocol    Protocol = "http/protobuff"
	OTLPValue                = "otlp"
	PrometheusValue          = "prometheus"
)

type Protocol string

type MonitoringConfig struct {
}

type CollectorConfig struct {
	LokiWriteEndpoint       string `json:"loki_write_endpoint" yaml:"lokiWriteEndpoint"`
	LokiWritePort           int    `json:"loki_write_port" yaml:"lokiWritePort"`
	PrometheusWriteEndpoint string `json:"prometheus_write_endpoint" yaml:"prometheusWriteEndpoint"`
	PrometheusWritePort     int    `json:"prometheus_write_port" yaml:"prometheusWritePort"`
	ZipkinEndpoint          string `json:"zipkin_endpoint" yaml:"zipkinEndpoint"`
	//TODO: Check if necessary
	ZipkinPort       int    `json:"zipkin_port" yaml:"zipkinPort"`
	OtlpHTTPEndpoint string `json:"otlp_http_endpoint" yaml:"otlpHTTPEndpoint"`
	OtlpHTTPPort     int    `json:"otlp_http_port" yaml:"otlpHTTPPort"`
	OtlpGRPCEndpoint string `json:"otlp_grpc_endpoint" yaml:"otlpGRPCEndpoint"`
	OtlpGRPCPort     int    `json:"otlp_grpc_port" yaml:"otlpGRPCPort"`
	OtlpMetricsPath  string `json:"otlp_metrics_path" yaml:"otlpMetricsPath"`
	OtlpLogsPath     string `json:"otlp_logs_path" yaml:"otlpLogsPath"`
	OtlpTracesPath   string `json:"otlp_traces_path" yaml:"otlpTracesPath"`
}
