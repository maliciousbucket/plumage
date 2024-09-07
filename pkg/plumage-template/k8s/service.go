package k8s

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/maliciousbucket/plumage/imports/k8s"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"strconv"
)

const (
	ServiceType = "ClusterIP"
)

type ServiceProps struct {
	Name       string
	Namespace  string
	Ports      []*plumagetemplate.ServicePort
	Monitoring *plumagetemplate.MonitoringConfig
}

func NewService(scope constructs.Construct, id string, props *ServiceProps) k8s.KubeService {
	metadata := serviceMetaData(props.Namespace, props.Name, props.Monitoring)
	labels := serviceLabels(props.Name)
	ports := ServicePorts(props)
	return k8s.NewKubeService(scope, jsii.String(id), &k8s.KubeServiceProps{
		Metadata: metadata,
		Spec: &k8s.ServiceSpec{
			Ports:    &ports,
			Selector: labels,
			Type:     jsii.String(ServiceType),
		},
	})
}

func serviceMetaData(nameSpace string, appLabel string, monitoring *plumagetemplate.MonitoringConfig) *k8s.ObjectMeta {
	labels := serviceLabels(appLabel)
	annotations := serviceAnnotations(appLabel, monitoring)
	return &k8s.ObjectMeta{
		Annotations: annotations,
		Labels:      labels,
		Name:        jsii.String(appLabel),
		Namespace:   jsii.String(nameSpace),
	}
}

func serviceLabels(appLabel string) *map[string]*string {
	labels := map[string]*string{
		AppLabel: &appLabel,
	}
	return &labels
}

func serviceAnnotations(appLabel string, monitoring *plumagetemplate.MonitoringConfig) *map[string]*string {
	annotations := map[string]*string{}

	if monitoring != nil {
		if monitoring.ScrapePort != 0 {
			port := strconv.FormatInt(int64(monitoring.ScrapePort), 10)
			annotations[PromScrapePortAnnotation] = &port
		}

		if monitoring.MetricsPath != "" {
			annotations[PromScrapePathAnnotation] = jsii.String(monitoring.MetricsPath)
		}

		annotations[PrometheusScrapeAnnotation] = jsii.String("true")
	}

	return &annotations
}

func ServicePorts(props *ServiceProps) []*k8s.ServicePort {
	var ports []*k8s.ServicePort
	i := 0
	for _, port := range props.Ports {
		name := PortName(port.Grpc, i)
		if props.Monitoring != nil {
			if props.Monitoring.ScrapePort != 0 {
				if port.ContainerPort == props.Monitoring.ScrapePort {
					name = "http-metrics"
				}
			}
		}
		protocol := GetProtocol(port.Protocol)
		ports = append(ports, &k8s.ServicePort{
			Port:        jsii.Number(port.HostPort),
			AppProtocol: nil,
			Name:        jsii.String(name),
			Protocol:    &protocol,
			TargetPort:  k8s.IntOrString_FromNumber(jsii.Number(port.ContainerPort)),
		})
		i++
	}
	return ports
}
