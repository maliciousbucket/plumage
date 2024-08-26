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
