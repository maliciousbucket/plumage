# Template

## Compose

## Env

## Monitoring

### Service Monitoring

- Aliases
  - Key
  - Env Value
- Metrics Port
- Metrics Endpoint
- Export or Scrape

- Loki / Prometheus
- HTTP / GRPC

### Default Service Monitoring Values

Environment

| Variable                            | Description | Default                                     |
|-------------------------------------|-------------|---------------------------------------------|
| OTEL_TRACES_EXPORTER                |             | otlp                                        |
| OTEL_METRICS_EXPORTER               |             | otlp                                        |
| OTEL_LOGS_EXPORTER                  |             | otlp                                        |
| OTEL_EXPORTER_OTLP_ENDPOINT         |             | {otlp grpc endpoint}:{otlp port}            |
| OTEL_EXPORTER_OTLP_TRACES_ENDPOINT  |             | {otlp grpc endpoint}:{otlp port}/v1/traces  |
| OTEL_EXPORTER_OTLP_METRICS_ENDPOINT |             | {otlp grpc endpoint}:{otlp port}/v1/metrics |
| OTEL_EXPORTER_OTLP_LOGS_ENDPOINT    |             | {otlp grpc endpoint}:{otlp port}/v1/logs    |
| OTEL_EXPORTER_OTLP_PROTOCOL         |             | grpc                                        |
| OTEL_EXPORTER_OTLP_TRACES_PROTOCOL  |             | grpc                                        |
| OTEL_EXPORTER_OTLP_METRICS_PROTOCOL |             | grpc                                        |
| OTEL_EXPORTER_OTLP_LOGS_PROTOCOL    |             | grpc                                        |

Service

| Variable        | Description | Default  |
|-----------------|-------------|----------|
| metricsPort     |             | 80       |
| metricsEndpoint |             | /metrics |
|                 |             |          |
|                 |             |          |

