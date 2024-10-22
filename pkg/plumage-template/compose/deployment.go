package compose

import (
	plumagetemplate "github.com/maliciousbucket/plumage/pkg/plumage-template"
	"github.com/maliciousbucket/plumage/pkg/types"
)

type DeploymentProps struct {
	Name           string
	Image          string
	Commands       []string
	Args           []string
	StartupProbe   *types.CommandProbe
	HealthCheck    *plumagetemplate.HttpProbe
	Ports          []plumagetemplate.ServicePort
	Resources      *plumagetemplate.ServiceResources
	Monitoring     *plumagetemplate.MonitoringConfig
	InitContainers []*plumagetemplate.InitContainer
	MinReplicas    int
	Env            map[string]string
}
