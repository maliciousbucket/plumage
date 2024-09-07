package k8s

import (
	"fmt"
	"github.com/maliciousbucket/plumage/pkg/config"
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/plumage-template/ingress"
	"github.com/maliciousbucket/plumage/pkg/types"
	"strconv"
)

func ServiceConfigOptions(template *plumagetemplate.PlumageTemplate, config *plumagetemplate.ServiceConfig, props *WebServiceProps) (*WebServiceProps, *SynthOpts, error) {
	var options []ServiceConfigFunc
	synthOptions := SynthOpts{
		Options: []SynthFunc{},
	}

	if config.Service.SynthOptions.Service {
		synthOptions.Options = append(synthOptions.Options, WithService())
	}

	if config.Service.SynthOptions.Deployment {
		synthOptions.Options = append(synthOptions.Options, WithDeployment())
	}

	if config.Service.SynthOptions.IngressRoute {
		options = append(options, WithServiceIngressRoute(config))
		synthOptions.Options = append(synthOptions.Options, WithIngressRoute())
	}

	if config.Service.SynthOptions.AutoScaling {
		options = append(options, WithDeploymentAutoScaler(config, &synthOptions))
	}

	if len(config.Service.Middleware) > 0 {
		options = append(options, WithMiddlewares(config, &synthOptions))
	}

	if config.Service.SynthOptions.ComposeVolumes {
		options = append(options, WithComposeVolumes(config, props))
	}

	if config.Service.SynthOptions.ComposeImage || config.Service.Image == "" {
		options = append(options, WithComposeImage(config))

	}

	if config.Service.SynthOptions.InitContainers {
		options = append(options, WithInitContainers(template, config))

	}

	if config.Service.SynthOptions.ComposePorts {
		options = append(options, WithComposePorts(config))
	}

	if config.Service.Monitoring != nil {
		if len(config.Service.Monitoring.MonitoringEnv) > 0 {
			options = append(options, WithComposeEnv(config))
		}
	}

	if config.Compose.HealthCheck != nil {
		options = append(options, WithComposeStartupCheck(config))
	}

	return nil, nil, nil
}

func loadBaseProps(template *plumagetemplate.PlumageTemplate, config *plumagetemplate.ServiceConfig, ns string) *WebServiceProps {
	if config.Service.SynthOptions == nil {
		config.Service.SynthOptions = &plumagetemplate.SynthOptions{
			Deployment:   true,
			IngressRoute: false,
			AutoScaling:  false,
			Service:      true,
			ComposeImage: false,
		}
	}

	return &WebServiceProps{
		Name:      config.Service.Name,
		Namespace: ns,
		Image:     config.Service.Image,
		Commands:  config.Service.Commands,
		Args:      config.Service.Args,
		//StartupProbe: config.Service.
		//Health Check
		Ports: config.Service.Ports,
		//InitContainers: initContainers,
		Resources:  config.Service.Resources,
		Monitoring: config.Service.Monitoring,
		//Resilience
		Scaling:     config.Service.Scaling,
		Env:         config.Service.Env,
		Middlewares: config.Service.Middleware,
	}
}

type ServiceConfigFunc func(p *WebServiceProps) error

func WithComposeVolumes(config *plumagetemplate.ServiceConfig, p *WebServiceProps) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		return nil
	}
}

func WithInitContainers(template *plumagetemplate.PlumageTemplate, config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		if len(config.Service.InitContainerNames) == 0 {
			return fmt.Errorf("no init containers specified")
		}
		var initContainers []*plumagetemplate.InitContainer

		if config.Service.InitContainerNames != nil && len(config.Service.InitContainerNames) > 0 {
			for _, container := range config.Service.InitContainerNames {
				init := template.GetInitContainer(container)
				if init == nil {
					return fmt.Errorf("init container %s not found", container)
				}
				initContainers = append(initContainers, init)
			}
		}
		return nil

	}
}

func WithServiceIngressRoute(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		ingressCfg := &ingress.RouteConfig{}
		loadIngressConfig(config.Service.Host, config.Service.Paths, config.Service.LoadBalancer, ingressCfg)
		if ingressCfg == nil {
			return fmt.Errorf("no ingress configuration found")
		}
		p.Ingress = ingressCfg
		return nil
	}
}

func WithDeploymentAutoScaler(config *plumagetemplate.ServiceConfig, opts *SynthOpts) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		if config.Service.Scaling != nil {
			p.Scaling = config.Service.Scaling
			opts.Options = append(opts.Options, WithAutoScaling())
		} else {
			opts.Options = append(opts.Options, WithDefaultAutoScaling())
		}
		return nil
	}
}

func WithMiddlewares(config *plumagetemplate.ServiceConfig, opts *SynthOpts) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		if len(config.Service.Middleware) == 0 {
			return fmt.Errorf("no middleware specified")
		}
		for _, middleware := range config.Service.Middleware {
			switch middleware {
			case "retry":
				opts.Options = append(opts.Options, WithRetry())
			case "circuitbreaker":
				opts.Options = append(opts.Options, WithCircuitBreaker())
			case "ratelimit":
				opts.Options = append(opts.Options, WithRateLimit())

			default:
				return fmt.Errorf("unknown middleware %s", middleware)

			}

		}
		return nil
	}
}

func WithComposeImage(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		composeImage := config.Compose.Image
		if composeImage == "" {
			if config.Service.Image == "" {
				return fmt.Errorf("no image image provided")
			}
			return fmt.Errorf("no compose image provided")
		}
		p.Image = composeImage
		return nil
	}
}

func WithComposeStartupCheck(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		if config.Compose.HealthCheck != nil {
			probe, err := types.ParseHealthCheck(config.Compose.HealthCheck)
			if err != nil {
				return err
			}
			p.StartupProbe = probe
		}
		return nil
	}
}

func WithComposePorts(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		composePorts := config.Compose.Ports
		if len(composePorts) == 0 {
			return fmt.Errorf("no compose ports specified")
		}
		var ports []*plumagetemplate.ServicePort
		for _, port := range composePorts {
			targetPort, err := strconv.ParseInt(port.Published, 10, 16)

			if err != nil {
				return err
			}
			portInt := int(targetPort)
			grpc := false
			if port.AppProtocol == "grpc" {
				grpc = true
			}

			ports = append(ports, &plumagetemplate.ServicePort{
				ContainerPort: int(port.Target),
				HostPort:      portInt,
				Protocol:      port.Protocol,
				Name:          "",
				Grpc:          grpc,
			})
		}
		p.Ports = ports
		return nil
	}

}

func WithComposeEnv(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		return nil
	}
}

func WithMonitoringEnv(config *plumagetemplate.ServiceConfig) ServiceConfigFunc {
	return func(p *WebServiceProps) error {
		return nil
	}
}

func loadIngressConfig(host string, paths []*plumagetemplate.ServicePaths, loadBalancer bool, config *ingress.RouteConfig) {
	if host != "" {
		config.Host = host
	}
	if len(paths) > 0 {
		var ingressPaths []*ingress.ServicePaths
		for _, path := range paths {
			ingressPaths = append(ingressPaths, &ingress.ServicePaths{
				Path: path.Path,
				Port: path.Port,
			})
		}
		if loadBalancer != false {
			config.EnableLoadBalancer = true
		}
	}
}

func getConfigMonitoringEnv(cfg *config.MonitoringConfig) map[string]string {
	return nil
}
