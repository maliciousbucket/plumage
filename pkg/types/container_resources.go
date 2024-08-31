package types

import (
	"fmt"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2"
	compose "github.com/compose-spec/compose-go/v2/types"
)

type ContainerCPU struct {
	CpuMillis *float64
}

func (c *ContainerCPU) Validate() error {
	if c.CpuMillis != nil {
		return fmt.Errorf("cpu millis and cpu cannot be specified at the same time")
	}

	return nil
}

type ContainerCpuResources struct {
	Request *ContainerCPU
	Limit   *ContainerCPU
}

func (c *ContainerCpuResources) Validate() error {
	if c.Request != nil {
		if err := c.Request.Validate(); err != nil {
			return err
		}
	}

	if c.Limit != nil {
		if err := c.Limit.Validate(); err != nil {
			return err
		}
	}

	if c.Request.CpuMillis != nil && c.Limit.CpuMillis != nil {
		if *c.Request.CpuMillis > *c.Limit.CpuMillis {
			return fmt.Errorf("cpu request cannot be higher than cpu limit. Request: %v, Limit: %v", c.Request.CpuMillis, c.Limit.CpuMillis)
		}
	}

	return nil
}

type ContainerMemoryResources struct {
	MemoryRequest *float64
	MemoryLimit   *float64
}

func (c *ContainerMemoryResources) Validate() error {
	if c.MemoryRequest != nil && c.MemoryLimit != nil {
		if *c.MemoryRequest > *c.MemoryLimit {
			return fmt.Errorf("memory request cannot be higher than memory limit. Request: %v, Limit: %v", c.MemoryRequest, c.MemoryLimit)
		}
	}
	return nil
}

type ContainerResources struct {
	Memory *ContainerMemoryResources
	Cpu    *ContainerCpuResources
}

func (r *ContainerResources) Validate() error {
	if r.Memory != nil {
		if err := r.Memory.Validate(); err != nil {
			return err
		}
	}
	if r.Cpu != nil {
		if err := r.Cpu.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (r *ContainerResources) CpuRequest() *float64 {
	if r.Cpu != nil {
		return r.CpuRequest()
	}

	return nil
}

func (r *ContainerResources) CpuLimit() *float64 {
	if r.Cpu != nil {
		return r.CpuLimit()
	}
	return nil
}

func (r *ContainerResources) MemoryRequest() *float64 {
	if r.Memory != nil {
		return r.Memory.MemoryRequest
	}
	return nil
}

func (r *ContainerResources) MemoryLimit() *float64 {
	if r.Memory != nil {
		return r.Memory.MemoryLimit
	}
	return nil
}

func CpuResources(s *compose.ServiceConfig) (*ContainerCpuResources, error) {
	return &ContainerCpuResources{}, nil
}

func MemoryResources(s *compose.ServiceConfig) (*ContainerMemoryResources, error) {
	return nil, nil
}

func (r *ContainerResources) ToK8sContainerResources() *cdk8splus30.ContainerResources {
	resources := cdk8splus30.ContainerResources{
		Cpu: &cdk8splus30.CpuResources{
			Limit:   cdk8splus30.Cpu_Millis(r.CpuLimit()),
			Request: cdk8splus30.Cpu_Millis(r.CpuRequest()),
		},
		EphemeralStorage: nil,
		Memory: &cdk8splus30.MemoryResources{
			Limit:   cdk8s.Size_Mebibytes(r.MemoryLimit()),
			Request: cdk8s.Size_Mebibytes(r.MemoryRequest()),
		},
	}
	return &resources
}
