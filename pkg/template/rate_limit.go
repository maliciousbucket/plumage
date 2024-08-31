package template

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type RateLimitConfig struct {
	Average  int                `yaml:"average"`
	Burst    int                `yaml:"burst"`
	Period   string             `yaml:"period"`
	Strategy *RateLimitStrategy `yaml:"strategy"`
}

func (r *RateLimitConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		Average  int       `yaml:"average"`
		Burst    int       `yaml:"burst"`
		Period   string    `yaml:"period"`
		Strategy yaml.Node `yaml:"strategy"`
	}
	if err := unmarshal(&aux); err != nil {
		return err
	}

	r.Average = aux.Average
	r.Burst = aux.Burst
	r.Period = aux.Period

	var strat RateLimitStrategy

	if len(aux.Strategy.Content) > 0 {
		key := aux.Strategy.Content[0].Value
		switch key {
		case string(StrategyTypeIpDepth):
			var strategy IpDepthStrategy
			if err := aux.Strategy.Decode(&strategy); err != nil {
				return err
			}
			strat = strategy
			r.Strategy = &strat
		case string(StrategyTypeHost):
			var strategy HostStrategy
			if err := aux.Strategy.Decode(&strategy); err != nil {
				return err
			}
			strat = strategy
			r.Strategy = &strat
		case string(StrategyTypeHeader):
			var strategy RequestHeaderStrategy
			if err := aux.Strategy.Decode(&strategy); err != nil {
				return err
			}
			strat = strategy
			r.Strategy = &strat
		default:
			return fmt.Errorf("unknown strategy type: %s", key)

		}
	}
	return nil
}

type RateLimitStrategy interface {
	StrategyType() StrategyType
}

type StrategyType string

const (
	StrategyTypeIpDepth = StrategyType("ipDepth")
	StrategyTypeHost    = StrategyType("host")
	StrategyTypeHeader  = StrategyType("header")
)

func (IpDepthStrategy) StrategyType() StrategyType {
	return StrategyTypeIpDepth
}

func (HostStrategy) StrategyType() StrategyType {
	return StrategyTypeHost
}

func (RequestHeaderStrategy) StrategyType() StrategyType {
	return StrategyTypeHeader
}

type IpDepthStrategy struct {
	Depth int `yaml:"depth"`
}

type RequestHeaderStrategy struct {
	HeaderName string `yaml:"header_name"`
}

type HostStrategy struct {
	Host bool `yaml:"host"`
}
