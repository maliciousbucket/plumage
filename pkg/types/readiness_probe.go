package types

import (
	"errors"
	compose "github.com/compose-spec/compose-go/v2/types"
)

type ReadinessProbe struct {
	Commands []*string
	Delay    *string
	Timeout  *string
	Interval *string
	Retries  *float64
}

func ParseHealthCheck(h *compose.HealthCheckConfig) (*ReadinessProbe, error) {
	probe := ReadinessProbe{}
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
