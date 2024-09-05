package resilience

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type RateLimitConfig struct {
	Average  int               `yaml:"average"`
	Burst    int               `yaml:"burst"`
	Period   string            `yaml:"period"`
	Strategy RateLimitStrategy `yaml:"strategy"`
}

func (r *RateLimitConfig) AverageRequests() int {
	return r.Average
}
func (r *RateLimitConfig) BurstRequests() int {
	return r.Burst
}

func (r *RateLimitConfig) RatePeriod() string {
	return r.Period
}

func (r *RateLimitConfig) LimitStrategy() RateLimitStrategy {
	return r.Strategy
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

	if len(aux.Strategy.Content) > 0 {
		key := aux.Strategy.Content[0].Value
		switch key {
		case string(StrategyTypeIpDepth):
			var strategy IpDepthStrategy
			if err := aux.Strategy.Content[1].Decode(&strategy); err != nil {
				return err
			}
			r.Strategy = &strategy
		case string(StrategyTypeHost):
			var strategy HostStrategy
			if err := aux.Strategy.Content[1].Decode(&strategy); err != nil {
				return err
			}

			r.Strategy = &strategy
		case string(StrategyTypeHeader):
			var strategy RequestHeaderStrategy
			if err := aux.Strategy.Content[1].Decode(&strategy); err != nil {
				return err
			}

			r.Strategy = &strategy
		default:
			return fmt.Errorf("unknown strategy type: %s", key)

		}
	}
	return nil
}
func (r *RateLimitConfig) Validate() error {
	return nil
}

type RateLimitStrategy interface {
	StrategyType() StrategyType
}

type StrategyType string

const (
	StrategyTypeIpDepth = StrategyType("ipDepth")
	StrategyTypeHost    = StrategyType("hostStrategy")
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
	HeaderName string `yaml:"headerName"`
}

type HostStrategy struct {
	Host bool `yaml:"host"`
}
