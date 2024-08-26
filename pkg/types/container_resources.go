package types

import (
	compose "github.com/compose-spec/compose-go/v2/types"
)

type ContainerCPU struct {
	CpuMillis  *float64
	CpuPercent *float64
}

type ContainerCpuResources struct {
	Request *ContainerCPU
	Limit   *ContainerCPU
}

type ContainerMemoryResources struct {
	MemoryRequest *string
	MemoryLimit   *string
}

type ContainerResources struct {
	Memory *ContainerMemoryResources
	Cpu    *ContainerCpuResources
}

func CpuResources(s *compose.ServiceConfig) (*ContainerCpuResources, error) {
	return &ContainerCpuResources{}, nil
}

func MemoryResources(s *compose.ServiceConfig) (*ContainerMemoryResources, error) {
	return nil, nil
}
