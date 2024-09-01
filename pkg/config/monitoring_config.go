package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
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
	return MonitoringConfig{
		AlloyAddress: "alloy.svc.cluster.local",
		Collectors: &CollectorConfig{
			LokiWriteEndpoint:       "/loki/api/push",
			LokiWritePort:           8090,
			PrometheusWriteEndpoint: "",
			PrometheusWritePort:     8080,
			ZipkinEndpoint:          "",
			ZipkinPort:              9411,
			OtlpHTTPEndpoint:        "",
			OtlpHTTPPort:            4318,
			OtlpGRPCEndpoint:        "",
			OtlpGRPCPort:            4317,
			OtlpMetricsPath:         "/v1/metrics",
			OtlpLogsPath:            "/v1/logs",
			OtlpTracesPath:          "/v1/traces",
		},
		EnvFile: "",
	}
}
