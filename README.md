# Plumage

## Configuration

### User Config

Default Location:

#### Helm Charts Config

| Name                          | Default       | Description                                                      |
|-------------------------------|---------------|------------------------------------------------------------------|
| argoVersion                   | argo-cd-7.6.8 |                                                                  |
| argoValuesFile                | ""            | Optional values file for the ArgoCD Helm Chart                   |
| stateMetricsVersion           | 5.25.1        |                                                                  |
| stateMetricsValuesFile        | ""            | Optional values file for the Kube State Metrics Helm Chart       |
| metricsVersion                | 3.12.2        |                                                                  |
| metricsValuesFile             | ""            | Optional values file for the Kube Metrics Server Helm Chart      |
| promOperatorVersion           | 14.0.0        |                                                                  |
| promOperatorValuesFile        | ""            | Optional values file for the Prometheus Operator CRDs Helm Chart |
| certManagerVersion            | v1.15.3       |                                                                  |
| certManagerValuesFile         | ""            | Optional values file for the Cert Manager Helm Chart             |
| kubePrometheusStackVersion    | 65.2.0        |                                                                  |
| kubePrometheusStackValuesFile | ""            | Optional values file for the Kube Prometheus Stack Helm Chart    |
| k6Version                     | 3.9.0         |                                                                  |
| k6ValuesFile                  | ""            | Optional values file for the K6Operator Helm Chart               |
| charts                        | nil           |                                                                  |

### GitHub Config

Default Location: `<Config Directory>/github.yaml`

| Name         | Description                                                                           |
|--------------|---------------------------------------------------------------------------------------|
| sourceOwner  | The GitHub username of the owner of the source repository    .                        |
| sourceRepo   | The repository to which Plumage will commit generated manifests.                      |
| commitBranch | The branch that Plumage will commit to.                                               |
| baseBranch   | The base branch of the repository.                                                    |
| authorName   | GitHub username of the author of commits made by Plumage                              |
| authorEmail  | Email of the author of commits made by Plumage                                        |
| envFile      | Optional .env file for the GitHub Config values. Takes precedence over the yaml file. |
| targetDir    | Where Plumage will look for manifests to commit. Ideally the same as the output dir.  |

### Monitoring Config

Default Location: `<Config Directory>/monitoring.yaml`

| Name         | Default                                  | Description                                                                          |
|--------------|------------------------------------------|--------------------------------------------------------------------------------------|
| alloyAddress | alloy.galah-monitoring.svc.cluster.local | Address of the Alloy instance in the cluster.                                        |
| envFile      | ""                                       | Optional .env file for Collector Config values. Takes precedence over the yaml file. |
| collectors   | nil                                      | Collectors Config. See below.                                                        |

#### Collector Config

| Name                      | Default                                  | Description                                                                                                                                                                       |
|---------------------------|------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| alloyAddress              | alloy.galah-monitoring.svc.cluster.local | The address of the alloy instance in the cluster                                                                                                                                  |
| defaultProtocol           | grpc                                     | Default protocol to use if none specified for other configuration options. Can be either `grpc` or `http/protobuff`                                                               |
| lokiWriteEndpoint         | /loki/api/push                           | The write endpoint exposed by the Loki instance in the cluster.                                                                                                                   |
| lokiWritePort             | 8090                                     | The port that the Loki query frontend is listening on.                                                                                                                            |
| prometheusWriteEndpoint   | ""                                       | The prometheus write endpoint to send metrics to. By default alloy listens on `0.0.0.0:<prometheusWritePort>` for remote write metrics.                                           |
| prometheusWritePort       | 8080                                     |                                                                                                                                                                                   |
| zipkinEndpoint            | ""                                       | The zipkin endpoint to send zipkin traces to. By default alloy listens on `0.0.0.0:<zipkinPort>` for zipkin traces.                                                               |
| zipkinPort                | 9411                                     | The port of the Zipkin endpoint.                                                                                                                                                  |
| otlpExportProtocol        | grpc                                     | The protocol for services to use for sending Open Telemetry metrics, logs and traces. Can be either `grpc` or `http/protobuff`. Takes precedence over the `default` config value. |
| otlpMetricsExportProtocol | grpc                                     | The protocol for services to use for sending Open Telemetry traces. Can be either `grpc` or `http/protobuff`. Takes precedence over the `otlpExportProtocol` config value.        |
| otlpLogsExportProtocol    | grpc                                     | The protocol for services to use for sending Open Telemetry traces. Can be either `grpc` or `http/protobuff`. Takes precedence over the `otlpExportProtocol` config value.        |
| otlpTracesExportProtocol  | grpc                                     | The protocol for services to use for sending Open Telemetry traces. Can be either `grpc` or `http/protobuff`. Takes precedence over the `otlpExportProtocol` config value.        |
| otlpHTTPEndpoint          | alloy.galah-monitoring.svc.cluster.local | The http endpoint to send Open Telemetry metrics and traces to.                                                                                                                   |
| otlpHTTPPort              | 4318                                     | The http port of the otlp http endpoint.                                                                                                                                          |
| otlpGRPCEndpoint          | alloy.galah-monitoring.svc.cluster.local | The otlp grpc endpoint to send Open Telemetry metrics and traces to.                                                                                                              |
| otlpGRPCPort              | 4317                                     | The port of the otlp collector grpc endpoint.                                                                                                                                     |
| otlpMetricsEndpoint       | alloy.galah-monitoring.svc.cluster.local | The endpoint for services to send otlp metrics to.                                                                                                                                |
| otlpLogsEndpoint          | alloy.galah-monitoring.svc.cluster.local | The endpoint for services to send otlp logs to.                                                                                                                                   |
| otlpTracesEndpoint        | alloy.galah-monitoring.svc.cluster.local | The endpoint to send otlp traces to.                                                                                                                                              |
| otlpMetricsPath           | /v1/metrics                              | The metrics path of the `otlpHTTPEndpoint`.                                                                                                                                       |
| otlpLogsPath              | /v1/logs                                 | The metrics path of the `otlpHTTPEndpoint`.                                                                                                                                       |
| otlpTracesPath            | /v1/traces                               | The metrics path of the `otlpHTTPEndpoint`.                                                                                                                                       |
| otlpMetricsExporter       | otlp                                     | The exporter for services to use for exporting metrics. Can be either `otlp`, `prometheus`, or `none`.                                                                            |
| otlpLogsExporter          | otlp                                     | The exporter for services to use for exporting logs. Can be either `otlp` or `none`.                                                                                              |
| otlpTracesExporter        | otlp                                     | The exporter for services to use for exporting traces. Can be either `otlp`, `jaeger`, `zipkin` or `none`.                                                                        |

### Environment

The following environment variables must be set:

| Variable          | Type   | Description                                                                                                        |
|-------------------|--------|--------------------------------------------------------------------------------------------------------------------|
| GITHUB_TOKEN      | string | The GitHub token that is used for creating commits and authorising access to the ArgoCD repository (if private)    |
| GITHUB_USERNAME   | string | The GitHub username that is used for creating commits and authorising access to the ArgoCD repository (if private) |
| ARGOCD_ADDRESS    | string | The address of the ArgoCD service                                                                                  |
| ARGOCD_REPOSITORY | string | The GitOps repository where manifests will be committed and which ArgoCD will create applications from             |
| CONFIG_DIR        | string | The directory which holds the configuration files for Plumage. Default: system config dir                          |



