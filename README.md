# Plumage

## TODO

- [ ] Fix probe configuration
- [ ] Re-implement resource requests/limits
- [ ] Add examples
- [ ] Complete Documentation
- [ ] Argo Integration Tests
- [ ] Unit Tests
- [ ] Refactor Packages
- [ ] Refactor CLI
- [ ] Update Makefile
- [ ] Update argo to use repo from config

## Configuration

### User Config

Default Location: `<Config Directory>/config.yaml`

| Name      | Default                                       | Description                                                     |
|-----------|-----------------------------------------------|-----------------------------------------------------------------|
| outputDir | ~./plumage/dist                               | The output dir where Plumage will create synthesised manifests. |
| namespace | galah-testbed                                 | The namespace of where plumage will deploy services.            |
| compose   | **deprecated**                                | To be removed.                                                  |
| template  | See [Template Config](#template-config)       | Configuration values for the template.                          |
| charts    | See [Helm Charts Config](#helm-charts-config) | Configuration values for helm charts.                           |


#### Template Config

| Name         | Default               | Description                       |
|--------------|-----------------------|-----------------------------------|
| workingDir   | ~./plumage/templates  | Location of template files.       |
| templateFile | plumage-template.yaml | Name of the template file to use. |


##### Helm Charts Config

| Name                          | Default       | Description                                                       |
|-------------------------------|---------------|-------------------------------------------------------------------|
| argoVersion                   | argo-cd-7.6.8 |                                                                   |
| argoValuesFile                | ""            | Optional values file for the ArgoCD Helm Chart.                   |
| stateMetricsVersion           | 5.25.1        |                                                                   |
| stateMetricsValuesFile        | ""            | Optional values file for the Kube State Metrics Helm Chart.       |
| metricsVersion                | 3.12.2        |                                                                   |
| metricsValuesFile             | ""            | Optional values file for the Kube Metrics Server Helm Chart.      |
| promOperatorVersion           | 14.0.0        |                                                                   |
| promOperatorValuesFile        | ""            | Optional values file for the Prometheus Operator CRDs Helm Chart. |
| certManagerVersion            | v1.15.3       |                                                                   |
| certManagerValuesFile         | ""            | Optional values file for the Cert Manager Helm Chart.             |
| kubePrometheusStackVersion    | 65.2.0        |                                                                   |
| kubePrometheusStackValuesFile | ""            | Optional values file for the Kube Prometheus Stack Helm Chart.    |
| k6Version                     | 3.9.0         |                                                                   |
| k6ValuesFile                  | ""            | Optional values file for the K6Operator Helm Chart.               |
| charts                        | nil           |                                                                   |


#### Charts Config


| Name      | Default | Description                   |
|-----------|---------|-------------------------------|
| configDir |         |                               |
| fileName  | ""      |                               |
| charts    | nil     | List of `ChartConfig` objects |




### ChartConfig

| Name        | Default | Description                                                   |
|-------------|---------|---------------------------------------------------------------|
| repository  | ""      |                                                               |
| namespace   | default |                                                               |
| chartName   | ""      |                                                               |
| releaseName | ""      |                                                               |
| version     | ""      |                                                               |
| replace     | false   |                                                               |
| valuesFiles | ""      |                                                               |
| values      | ""      |                                                               |
| localFile   | ""      |                                                               |
| skipCRDs    | false   |                                                               |
| upgradeCRDs | false   |                                                               |
| labels      | nil     | Additional labels for the chart.                              |
| lint        | false   | Determines if the chart will be linted prior to installation. |


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


## Commands



```console
Usage:
  plumage [command]

Available Commands:
  app               Manage applications
  charts            Manage Helm Charts
  commit            Commit and Push synthesised manifests
  completion        Generate the autocompletion script for the specified shell
  config            configuration
  deploy            Deployment commands
  expose            Expose as service for testing
  help              Help about any command
  project           Manage ArgoCD projects
  service           Watch Kubernetes services
  set-argo-token    Set argo-cd token for interacting with the API
  synth             Synth Kubernetes manifests
  wait-related-pods Wait for related pods to be ready
  watch             Watch Kubernetes resources

Flags:
      --config string   config file (default is $HOME/.plumage.yaml)
  -h, --help            help for plumage

Use "plumage [command] --help" for more information about a command.

```

### Deployment

```console
Deployment commands

Usage:
  plumage deploy [flags]
  plumage deploy [command]

Available Commands:
  app           Deploy synthesised applications
  argo-auth     Add GitHub credentials to ArgoCD
  cluster       manage argo's clusters
  gateway       Deploy traefik gateway
  install-argo  Install Argo CD
  install-chart Install Helm Charts
  monitoring    Deploy monitoring infrastructure
  sync          Sync ArgoCD resources
  template      Deploy a template

Flags:
  -h, --help   help for deploy

Global Flags:
      --config string   config file (default is $HOME/.plumage.yaml)

Use "plumage deploy [command] --help" for more information about a command.

```

### Charts

```console
Manage Helm Charts

Usage:
  plumage charts [flags]
  plumage charts [command]

Available Commands:
  install     Install Helm Charts

Flags:
  -h, --help   help for charts

Global Flags:
      --config string   config file (default is $HOME/.plumage.yaml)

```

```console
Install Helm Charts

Usage:
  plumage charts install [flags]

Flags:
      --argo                    Install Argo
      --cert-manager            Install Certificate Manager
  -h, --help                    help for install
      --k6-operator             Install K6 Operator
      --kube-metrics-server     Install Metrics Server
      --kube-prometheus-stack   Install Prometheus Stack
      --kube-state-metrics      Install KubeState Metrics
      --prom-operator-crds      Install Prometheus Operator CRDs
  -r, --replace                 Replace existing Helm Charts
      --values-file string      Chart values.yaml file
  -v, --version string          Version

Global Flags:
      --config string   config file (default is $HOME/.plumage.yaml)

```

### Commit

```console
Commit and Push synthesised manifests

Usage:
  plumage commit [flags]
  plumage commit [command]

Available Commands:
  gateway     Commit traefik gateway manifests
  manifests   commit manifests

Flags:
  -e, --env-files string      comma seperated environment files
  -h, --help                  help for commit
  -m, --message string        commit message
  -t, --target-dir string     target directory
  -f, --target-files string   comma seperated target files

Global Flags:
      --config string   config file (default is $HOME/.plumage.yaml)

Use "plumage commit [command] --help" for more information about a command.
```

### Synth

```console
Synth Kubernetes manifests

Usage:
  plumage synth [flags]
  plumage synth [command]

Available Commands:
  gateway     Synth gateway manifests
  service     Synth manifests for a service
  template    Synth template manifests
  tests       Synth manifests for a tests

Flags:
  -h, --help   help for synth

Global Flags:
      --config string   config file (default is $HOME/.plumage.yaml)

Use "plumage synth [command] --help" for more information about a command.
```