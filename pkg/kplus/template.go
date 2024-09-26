package kplus

import (
	"errors"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	kplus "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	"github.com/maliciousbucket/plumage/pkg/resilience"
	"gopkg.in/yaml.v3"
)

type Template struct {
	Name      string            `yaml:"name"`
	Namespace string            `yaml:"namespace"`
	Services  []ServiceTemplate `yaml:"services"`
}

type ServiceTemplate struct {
	Namespace          string                           `yaml:"namespace"`
	Name               string                           `yaml:"name"`
	Image              string                           `yaml:"image"`
	Args               []string                         `yaml:"args"`
	Commands           []string                         `yaml:"commands"`
	Paths              []ServicePath                    `yaml:"paths"`
	Ports              []Port                           `yaml:"ports"`
	LivenessProbe      Probe                            `yaml:"liveness_probe,omitempty"`
	ReadinessProbe     Probe                            `yaml:"readiness_probe,omitempty"`
	HealthCheckProbe   Probe                            `yaml:"health_check_probe,omitempty"`
	VolumeMounts       map[string]string                `yaml:"volumeMounts,omitempty"`
	FileMounts         []map[string]string              `yaml:"fileMounts,omitempty"`
	EmptyDirs          []string                         `yaml:"emptyDirs,omitempty"`
	WorkingDir         string                           `yaml:"workingDir"`
	Env                map[string]string                `yaml:"env,omitempty"`
	EnvFile            string                           `yaml:"envFile"`
	Monitoring         *MonitoringTemplate              `yaml:"monitoring,omitempty"`
	Replicas           int                              `yaml:"replicas"`
	Scaling            ScalingTemplate                  `yaml:"scaling"`
	DefaultMiddleware  []string                         `yaml:"defaultMiddleware,omitempty"`
	DefaultAutoScaling DefaultAutoScaling               `yaml:"defaultAutoScaling,omitempty"`
	CircuitBreaker     *resilience.CircuitBreakerConfig `yaml:"circuitBreaker,omitempty"`
	Retry              *resilience.RetryConfig          `yaml:"retry,omitempty"`
	RateLimit          *resilience.RateLimitConfig      `yaml:"rateLimit,omitempty"`
}

type ServicePath struct {
	Path string `yaml:"path"`
	Port int    `yaml:"port"`
}

type Port struct {
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Protocol string `yaml:"protocol,omitempty"`
	GRPC     bool   `yaml:"grpc"`
}

type MonitoringTemplate struct {
	MonitoringEnv []string          `yaml:"env"`
	Aliases       map[string]string `yaml:"aliases"`
	ScrapePort    int               `yaml:"scrapePort"`
	ScrapePath    string            `yaml:"scrapePath"`
}

type Probe struct {
	Type  string       `yaml:"type"`
	Probe ServiceProbe `yaml:"probe"`
}

type DefaultAutoScaling struct {
	Type           string         `yaml:"type"`
	DefaultScaling DefaultScaling `yaml:"scaling"` //vertical / horizontal

}

func (s *DefaultAutoScaling) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}
	scalingType, ok := raw["type"].(string)
	if !ok {
		return errors.New("unknown scaling type")
	}
	//TODO: test
	//data, ok := raw[scaling].(map[string]interface{})
	//if !ok {
	//	return errors.New("unknown or missing scaling type")
	//}

	s.Type = scalingType
	switch scalingType {
	case "horizontal":
		var horizontal DefaultHorizontalScaling
		if err := unmarshal(&horizontal); err != nil {
			return err
		}
	case "vertical":
		var vertical DefaultVerticalScaling
		if err := unmarshal(&vertical); err != nil {
			return err
		}
	default:
		return errors.New("unknown scaling type")
	}
	return nil
}

type DefaultScaling interface {
	ScalingType() ScalingType
}

type DefaultHorizontalScaling struct {
	MinReplicas int `yaml:"minReplicas"`
	MaxReplicas int `yaml:"maxReplicas"`
}

func (s *DefaultHorizontalScaling) ScalingType() ScalingType {
	return ScalingTypeHorizontal
}

type DefaultVerticalScaling struct {
	ControlLimits bool `yaml:"controlLimits"`
	MinCpuMillis  int  `yaml:"minCpuMillis"`
	MaxCpuMillis  int  `yaml:"maxCpuMillis"`
	MinMemoryMi   int  `yaml:"minMemoryMb"`
	MaxMemoryMi   int  `yaml:"maxMemoryMb"`
}

func (s *DefaultVerticalScaling) ScalingType() ScalingType {
	return ScalingTypeVertical
}

func (p *Probe) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw map[string]interface{}
	if err := unmarshal(&raw); err != nil {
		return err
	}

	probeType, ok := raw["type"].(string)
	if !ok {
		return errors.New("unknown probe type")
	}

	data, ok := raw["probe"].(map[string]interface{})
	if !ok {
		return errors.New("unknown or missing probe")
	}
	p.Type = probeType
	switch probeType {
	case "http":
		var httpProbe HttpProbe
		if err := decodeNode(data, &httpProbe); err != nil {
			return err
		}
	case "tcp":
		var probe TCPProbe
		if err := decodeNode(data, &probe); err != nil {
			return err
		}
	case "command":
		var commandProbe CommandProbe
		if err := decodeNode(data, &commandProbe); err != nil {
			return err
		}
	default:
		return errors.New("unknown probe type")
	}
	return nil
}

type ServiceProbe interface {
	toKplusProbe() kplus.Probe
}

func decodeNode(data map[string]interface{}, target interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlData, target)
}

func ToKplusProbe(p ServiceProbe) kplus.Probe {
	return p.toKplusProbe()
}

type HttpProbe struct {
	Path                string `yaml:"path"`
	Port                int    `yaml:"port"`
	Host                string `yaml:"host"`
	Threshold           int    `yaml:"threshold"`
	InitialDelaySeconds int    `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int    `yaml:"timeoutSeconds"`
	PeriodSeconds       int    `yaml:"periodSeconds"`
	HTTPS               bool   `yaml:"https"`
}

func (p *HttpProbe) toKplusProbe() kplus.Probe {
	scheme := kplus.ConnectionScheme_HTTP
	if p.HTTPS {
		scheme = kplus.ConnectionScheme_HTTPS
	}
	return kplus.Probe_FromHttpGet(jsii.String(p.Path), &kplus.HttpGetProbeOptions{
		FailureThreshold:    jsii.Number(p.Threshold),
		InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(p.InitialDelaySeconds)),
		PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(p.PeriodSeconds)),
		TimeoutSeconds:      cdk8s.Duration_Seconds(jsii.Number(p.TimeoutSeconds)),
		Host:                jsii.String(p.Host),
		Port:                jsii.Number(p.Port),
		Scheme:              scheme,
	})
}

type TCPProbe struct {
	Port                int    `yaml:"port"`
	Host                string `yaml:"host"`
	Threshold           int    `yaml:"threshold"`
	InitialDelaySeconds int    `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int    `yaml:"timeoutSeconds"`
	PeriodSeconds       int    `yaml:"periodSeconds"`
}

func (p *TCPProbe) toKplusProbe() kplus.Probe {
	return kplus.Probe_FromTcpSocket(&kplus.TcpSocketProbeOptions{
		FailureThreshold:    jsii.Number(p.Threshold),
		InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(p.InitialDelaySeconds)),
		PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(p.PeriodSeconds)),
		TimeoutSeconds:      cdk8s.Duration_Seconds(jsii.Number(p.TimeoutSeconds)),
		Host:                jsii.String(p.Host),
		Port:                jsii.Number(p.Port),
	})
}

type CommandProbe struct {
	Commands            []string `yaml:"commands"`
	Threshold           int      `yaml:"threshold"`
	InitialDelaySeconds int      `yaml:"initialDelaySeconds"`
	TimeoutSeconds      int      `yaml:"timeoutSeconds"`
	PeriodSeconds       int      `yaml:"periodSeconds"`
}

func (p *CommandProbe) toKplusProbe() kplus.Probe {
	commands := StringSliceToK8s(p.Commands)
	return kplus.Probe_FromCommand(commands, &kplus.CommandProbeOptions{
		FailureThreshold:    jsii.Number(p.Threshold),
		InitialDelaySeconds: cdk8s.Duration_Seconds(jsii.Number(p.InitialDelaySeconds)),
		PeriodSeconds:       cdk8s.Duration_Seconds(jsii.Number(p.PeriodSeconds)),
		TimeoutSeconds:      cdk8s.Duration_Seconds(jsii.Number(p.TimeoutSeconds)),
	})
}
