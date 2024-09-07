package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	GRPCProtocol    Protocol = "grpc"
	HTTPProtocol    Protocol = "http/protobuff"
	OTLPValue                = "otlp"
	PrometheusValue          = "prometheus"
)

type Protocol string

type MonitoringConfig struct {
	AlloyAddress string           `json:"alloyAddress"`
	Collectors   *CollectorConfig `yaml:"collectors"`
	EnvFile      string           `json:"envFile"`
}

type CollectorConfig struct {
	DefaultProtocol         string `yaml:"defaultProtocol"`
	LokiWriteEndpoint       string `json:"loki_write_endpoint" yaml:"lokiWriteEndpoint"`
	LokiWritePort           int    `json:"loki_write_port" yaml:"lokiWritePort"`
	PrometheusWriteEndpoint string `json:"prometheus_write_endpoint" yaml:"prometheusWriteEndpoint"`
	PrometheusWritePort     int    `json:"prometheus_write_port" yaml:"prometheusWritePort"`
	ZipkinEndpoint          string `json:"zipkin_endpoint" yaml:"zipkinEndpoint"`
	//TODO: Check if necessary
	ZipkinPort                int    `json:"zipkin_port" yaml:"zipkinPort"`
	OtlpExportProtocol        string `json:"otlp_export_protocol" yaml:"otlpExportProtocol"`
	OtlpMetricsExportProtocol string `json:"otlp_metrics_export_protocol" yaml:"otlpMetricsExportProtocol"`
	OtlpLogsExportProtocol    string `json:"otlp_logs_export_protocol" yaml:"otlpLogsExportProtocol"`
	OtlpTracesExportProtocol  string `json:"otlp_traces_export_protocol" yaml:"otlpTracesExportProtocol"`
	OtlpHTTPEndpoint          string `json:"otlp_http_endpoint" yaml:"otlpHTTPEndpoint"`
	OtlpHTTPPort              int    `json:"otlp_http_port" yaml:"otlpHTTPPort"`
	OtlpGRPCEndpoint          string `json:"otlp_grpc_endpoint" yaml:"otlpGRPCEndpoint"`
	OtlpGRPCPort              int    `json:"otlp_grpc_port" yaml:"otlpGRPCPort"`
	OtlpMetricsEndpoint       string `json:"otlp_metrics_endpoint" yaml:"otlpMetricsEndpoint"`
	OtlpLogsEndpoint          string `json:"otlp_logs_endpoint" yaml:"otlpLogsEndpoint"`
	OtlpTracesEndpoint        string `json:"otlp_traces_endpoint" yaml:"otlpTracesEndpoint"`

	OtlpMetricsPath     string `json:"otlp_metrics_path" yaml:"otlpMetricsPath"`
	OtlpLogsPath        string `json:"otlp_logs_path" yaml:"otlpLogsPath"`
	OtlpTracesPath      string `json:"otlp_traces_path" yaml:"otlpTracesPath"`
	OtlpMetricsExporter string `json:"otlp_metrics_exporter" yaml:"otlpMetricsExporter"`
	OtlpLogsExporter    string `json:"otlp_logs_exporter" yaml:"otlpLogsExporter"`
	OtlpTracesExporter  string `json:"otlp_traces_exporter" yaml:"otlpTracesExporter"`
}

//TODO: Add switch between http and grpc

func (c CollectorConfig) ToStringMap() map[string]string {
	useHttp := false
	switch strings.ToLower(c.DefaultProtocol) {
	case "http":
		useHttp = true
	case "grpc":
		useHttp = false
	default:
		useHttp = false
	}

	metricsEndpoint := c.OtlpMetricsEndpoint
	logsEndpoint := c.OtlpLogsEndpoint
	tracesEndpoint := c.OtlpTracesEndpoint
	otlpExporterProtocol := c.OtlpExportProtocol

	if useHttp {

		metricsEndpoint = fmt.Sprintf("%s:%d/%s", c.OtlpHTTPEndpoint, c.OtlpHTTPPort, c.OtlpMetricsEndpoint)
		logsEndpoint = fmt.Sprintf("%s:%d/%s", c.OtlpHTTPEndpoint, c.OtlpHTTPPort, c.OtlpLogsEndpoint)
		tracesEndpoint = fmt.Sprintf("%s:%d/%s", c.OtlpHTTPEndpoint, c.OtlpHTTPPort, c.OtlpTracesEndpoint)
		otlpExporterProtocol = "http/protobuff"
	}

	return map[string]string{
		"LOKI_ENDPOINT":       fmt.Sprintf("%s:%d", c.LokiWriteEndpoint, c.LokiWritePort),
		"PROMETHEUS_ENDPOINT": fmt.Sprintf("%s:%d", c.PrometheusWriteEndpoint, c.PrometheusWritePort),
		"ZIPKIN_ENDPOINT":     fmt.Sprintf("%s:%d", c.ZipkinEndpoint, c.ZipkinPort),

		"OTEL_EXPORTER_OTLP_PROTOCOL":         otlpExporterProtocol,
		"OTEL_EXPORTER_OTLP_METRICS_PROTOCOL": c.OtlpMetricsExportProtocol,
		"OTEL_EXPORTER_OTLP_LOGS_PROTOCOL":    c.OtlpLogsExportProtocol,
		"OTEL_EXPORTER_OTLP_TRACES_PROTOCOL":  c.OtlpTracesExportProtocol,

		"OTEL_EXPORTER_OTLP_ENDPOINT": fmt.Sprintf("%s:%d", c.OtlpGRPCEndpoint, c.OtlpGRPCPort),
		"OTEL_EXPORTER_HTTP_ENDPOINT": fmt.Sprintf("%s:%d", c.OtlpHTTPEndpoint, c.OtlpHTTPPort),

		"OTEL_METRICS_ENDPOINT": metricsEndpoint,
		"OTEL_LOGS_ENDPOINT":    logsEndpoint,
		"OTEL_TRACES_ENDPOINT":  tracesEndpoint,

		"OTEL_METRICS_EXPORTER": c.OtlpMetricsExporter,
		"OTEL_LOGS_EXPORTER":    c.OtlpLogsExporter,
		"OTEL_TRACES_EXPORTER":  c.OtlpTracesExporter,
	}
}

func NewMonitoringConfig(configDir string, envFile string) (*MonitoringConfig, error) {
	base := loadDefaultMonitoringConfig()

	config, err := loadMonitoringConfig(configDir, envFile, &base)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func loadMonitoringConfig(configDir string, envFile string, base *MonitoringConfig) (*MonitoringConfig, error) {
	fileName := filepath.Join(configDir, "monitoring.yaml")

	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(fileName)
			if err != nil {
				return nil, err
			}
			file.Close()
		} else {
			return nil, err
		}
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, base); err != nil {
		return nil, err
	}

	if envFile != "" {
		base.EnvFile = envFile
		err := loadMonitoringConfigFromEnv(envFile, base)
		if err != nil {
			return nil, err
		}
	}
	return base, nil

}

func loadMonitoringConfigFromEnv(envFile string, base *MonitoringConfig) error {
	if envFile == "" {
		return nil
	}

	if _, err := os.Stat(envFile); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("env file %s does not exist", envFile)
		}
	}
	err := godotenv.Load(envFile)
	if err != nil {
		return err
	}

	//var env map[string]string
	env, err := godotenv.Read(envFile)
	if err != nil {
		return err
	}

	if len(env) > 0 {
		if alloyAddress, ok := env["ALLOY_ADDRESS"]; ok {
			base.AlloyAddress = alloyAddress
		}

		if lokiEndpoint, ok := env["LOKI_ENDPOINT"]; ok {
			base.Collectors.LokiWriteEndpoint = lokiEndpoint
		}

		if lokiWritePort, ok := env["LOKI_WRITE_PORT"]; ok {
			port, err := strconv.Atoi(lokiWritePort)
			if err != nil {
				return fmt.Errorf("failed to convert LOKI_WRITE_PORT to int: %v", err)
			}
			base.Collectors.LokiWritePort = port
		}

		if prometheusEndpoint, ok := env["LOKI_PROMETHEUS_ENDPOINT"]; ok {
			base.Collectors.LokiWriteEndpoint = prometheusEndpoint
		}

		if prometheusWritePort, ok := env["PROMETHEUS_WRITE_PORT"]; ok {
			port, err := strconv.Atoi(prometheusWritePort)
			if err != nil {
				return fmt.Errorf("failed to convert PROMETHEUS_WRITE_PORT to int: %v", err)
			}
			base.Collectors.PrometheusWritePort = port
		}
		if zipkinEndpoint, ok := env["ZIPKIN_ENDPOINT"]; ok {
			base.Collectors.ZipkinEndpoint = zipkinEndpoint
		}

		if zipkinWritePort, ok := env["ZIPKIN_PORT"]; ok {
			port, err := strconv.Atoi(zipkinWritePort)
			if err != nil {
				return fmt.Errorf("failed to convert ZIPKIN_PORT to int: %v", err)
			}
			base.Collectors.ZipkinPort = port
		}

		if otlpHttpEndpoint, ok := env["OTLP_ENDPOINT"]; ok {
			base.Collectors.OtlpHTTPEndpoint = otlpHttpEndpoint
		}
		if otlpHTTPPort, ok := env["OTLP_HTTP_PORT"]; ok {
			port, err := strconv.Atoi(otlpHTTPPort)
			if err != nil {
				return fmt.Errorf("failed to convert OTLP_HTTP_PORT to int: %v", err)
			}
			base.Collectors.OtlpHTTPPort = port
		}

		if otlpGRPCEndpoint, ok := env["OTLP_GRPC_ENDPOINT"]; ok {
			base.Collectors.OtlpGRPCEndpoint = otlpGRPCEndpoint
		}
		if otlpGRPCPort, ok := env["OTLP_GRPC_PORT"]; ok {
			port, err := strconv.Atoi(otlpGRPCPort)
			if err != nil {
				return fmt.Errorf("failed to convert OTLP_GRPC_PORT to int: %v", err)
			}
			base.Collectors.OtlpGRPCPort = port
		}
		if otlpMetricsPath, ok := env["OTLP_METRICS_PATH"]; ok {
			base.Collectors.OtlpMetricsPath = otlpMetricsPath
		}
		if otlpLogsPath, ok := env["OTLP_LOGS_PATH"]; ok {
			base.Collectors.OtlpLogsPath = otlpLogsPath
		}
		if otlpTracesPath, ok := env["OTLP_TRACES_PATH"]; ok {
			base.Collectors.OtlpTracesPath = otlpTracesPath
		}
	}
	return nil

}

func loadDefaultMonitoringConfig() MonitoringConfig {
	alloyAddress := "alloy.svc.cluster.local"
	return MonitoringConfig{
		AlloyAddress: alloyAddress,
		Collectors: &CollectorConfig{
			DefaultProtocol:           "grpc",
			LokiWriteEndpoint:         "/loki/api/push",
			LokiWritePort:             8090,
			PrometheusWriteEndpoint:   "",
			PrometheusWritePort:       8080,
			ZipkinEndpoint:            "",
			ZipkinPort:                9411,
			OtlpExportProtocol:        string(GRPCProtocol),
			OtlpMetricsExportProtocol: string(GRPCProtocol),
			OtlpLogsExportProtocol:    string(GRPCProtocol),
			OtlpTracesExportProtocol:  string(GRPCProtocol),
			OtlpHTTPEndpoint:          "",
			OtlpHTTPPort:              4318,
			OtlpGRPCEndpoint:          "",
			OtlpGRPCPort:              4317,
			OtlpMetricsEndpoint:       fmt.Sprintf("%s:%s", alloyAddress, "4317"),
			OtlpLogsEndpoint:          fmt.Sprintf("%s:%s", alloyAddress, "4317"),
			OtlpTracesEndpoint:        fmt.Sprintf("%s:%s", alloyAddress, "4317"),
			OtlpMetricsPath:           "/v1/metrics",
			OtlpLogsPath:              "/v1/logs",
			OtlpTracesPath:            "/v1/traces",
			OtlpMetricsExporter:       OTLPValue,
			OtlpLogsExporter:          OTLPValue,
			OtlpTracesExporter:        OTLPValue,
		},
		EnvFile: "",
	}
}
