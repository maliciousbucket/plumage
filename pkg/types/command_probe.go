package types

import (
	"errors"
	"fmt"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	compose "github.com/compose-spec/compose-go/v2/types"
)

type CommandProbe struct {
	Commands []*string
	Delay    *string
	Timeout  *string
	Interval *string
	Retries  *float64
}

func ParseHealthCheck(h *compose.HealthCheckConfig) (*CommandProbe, error) {
	probe := CommandProbe{}
	if h.StartPeriod != nil {
		delay := h.StartPeriod.String()
		probe.Delay = &delay
	}

	if h.Timeout != nil {
		timeout := h.Timeout.String()
		probe.Timeout = &timeout
	}

	if h.Interval != nil {
		interval := h.Interval.String()
		probe.Interval = &interval
	}

	if h.Retries != nil {
		retries := float64(*h.Retries)
		probe.Retries = &retries
	}

	if len(h.Test) == 0 {
		return nil, errors.New("health check must have at least one test")
	}

	var commands []*string
	for _, command := range h.Test {
		cmd := command
		commands = append(commands, &cmd)
	}
	probe.Commands = commands
	return &probe, nil
}

func (p *CommandProbe) parseCommandProbeOpts() (*cdk8splus30.CommandProbeOptions, error) {
	var probeOpts cdk8splus30.CommandProbeOptions

	if len(p.Commands) == 0 {
		return nil, fmt.Errorf("no commands specified for probe")
	}

	if p.Retries != nil {
		probeOpts.FailureThreshold = p.Retries
	}

	if p.Delay != nil {
		probeOpts.InitialDelaySeconds = cdk8s.Duration_Parse(p.Delay)
	}

	if p.Interval != nil {
		probeOpts.PeriodSeconds = cdk8s.Duration_Parse(p.Interval)
	}

	if p.Timeout != nil {
		probeOpts.TimeoutSeconds = cdk8s.Duration_Parse(p.Timeout)
	}
	return &probeOpts, nil

}

func (p *CommandProbe) Probe() (cdk8splus30.Probe, error) {
	probeOpts, err := p.parseCommandProbeOpts()
	if err != nil {
		return nil, err
	}
	return cdk8splus30.Probe_FromCommand(&p.Commands, probeOpts), nil
}
