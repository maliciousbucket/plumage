package kubernetes

const (
	PrometheusScrapeAnnotation = "prometheus.io/scrape"
	PrometheusPortAnnotation   = "prometheus.io/port"
	PrometheusPathAnnotation   = "prometheus.io/path"

	KubernetesNameLabel = "app.kubernetes.io/name"
)

func MapToAnnotations(m map[string]string) *map[string]*string {
	annotations := make(map[string]*string, len(m))

	for k, v := range m {
		annotations[k] = &v
	}
	return &annotations
}

//TODO: Redundant

func MapToLabels(m map[string]string) *map[string]*string {
	labels := make(map[string]*string, len(m))

	for k, v := range m {
		labels[k] = &v
	}
	return &labels
}
