package main

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"github.com/maliciousbucket/plumage/cmd"
	"github.com/maliciousbucket/plumage/pkg/config"
)

var (
	Cfg *config.AppConfig
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func main() {
	appConfig, err := config.NewAppConfig("./config/compose", "galah-testbed", "", "")
	if err != nil {
		panic(err)
	}
	Cfg = appConfig
	cmd.Execute(Cfg)
}
