package types

import (
	"fmt"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
)

type HttpProbe struct {
	Path                *string
	Port                *float64
	InitialDelaySeconds *float64
	IntervalSeconds     *float64
	TimeoutSeconds      *float64
	Retries             *float64
}

func (p *HttpProbe) parseLivenessProbeOpts() (*cdk8splus30.HttpGetProbeOptions, error) {
	if p.Path == nil || *p.Path == "" {
		return nil, fmt.Errorf("live check path is required")
	}

	if p.Port == nil || *p.Port == 0 {
		return nil, fmt.Errorf("live check port is required")
	}

	var probeOpts cdk8splus30.HttpGetProbeOptions

	if p.InitialDelaySeconds != nil {
		probeOpts.InitialDelaySeconds = cdk8s.Duration_Seconds(p.InitialDelaySeconds)
	}

	if p.IntervalSeconds != nil {
		probeOpts.PeriodSeconds = cdk8s.Duration_Seconds(p.IntervalSeconds)
	}

	if p.Retries != nil {
		probeOpts.FailureThreshold = p.Retries
	}

	if p.TimeoutSeconds != nil {
		probeOpts.TimeoutSeconds = cdk8s.Duration_Seconds(p.TimeoutSeconds)
	}

	probeOpts.Port = p.Port
	return &probeOpts, nil
}

func (p *HttpProbe) Probe() (cdk8splus30.Probe, error) {
	probeOpts, err := p.parseLivenessProbeOpts()
	if err != nil {
		return nil, err
	}
	return cdk8splus30.Probe_FromHttpGet(p.Path, probeOpts), nil
}
