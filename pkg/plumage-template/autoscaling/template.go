package autoscaling

import (
	"errors"
	"github.com/maliciousbucket/plumage/pkg/kplus"
	"gopkg.in/yaml.v3"
)

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

	data, ok := raw["scaling"].(map[string]interface{})
	if !ok {
		return errors.New("unknown or missing scaling type")
	}

	s.Type = scalingType
	switch scalingType {
	case "horizontal":
		var horizontal DefaultHorizontalScaling
		if err := decodeNode(data, &horizontal); err != nil {
			return err
		}
		s.DefaultScaling = &horizontal
	case "vertical":
		var vertical DefaultVerticalScaling
		if err := decodeNode(data, &vertical); err != nil {
			return err
		}
		s.DefaultScaling = &vertical
	default:
		return errors.New("unknown scaling type")
	}
	return nil
}

type DefaultScaling interface {
	ScalingType() kplus.ScalingType
}

type DefaultHorizontalScaling struct {
	MinReplicas int `yaml:"minReplicas"`
	MaxReplicas int `yaml:"maxReplicas"`
}

func (s *DefaultHorizontalScaling) ScalingType() kplus.ScalingType {
	return ScalingTypeHorizontal
}

type DefaultVerticalScaling struct {
	ControlLimits bool `yaml:"controlLimits"`
	MinCpuMillis  int  `yaml:"minCpuMillis"`
	MaxCpuMillis  int  `yaml:"maxCpuMillis"`
	MinMemoryMi   int  `yaml:"minMemoryMb"`
	MaxMemoryMi   int  `yaml:"maxMemoryMb"`
}

func (s *DefaultVerticalScaling) ScalingType() kplus.ScalingType {
	return ScalingTypeVertical
}

func decodeNode(data map[string]interface{}, target interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlData, target)
}

const (
	ScalingTypeVertical   = kplus.ScalingType("vertical")
	ScalingTypeHorizontal = kplus.ScalingType("horizontal")
	ScalingTypeUnknown    = kplus.ScalingType("unknown")
)
