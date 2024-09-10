package kubeclient

import (
	"context"
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/release"
	"helm.sh/helm/v3/pkg/repo"
)

type HelmClient interface {
	InstallChart()
	UpgradeChart()
	GetDeployedRelease()
	UninstallRelease()
	AddChartRepo()
	UpdateChartRepo()
}

type HelmOptions struct {
	NameSpace        string
	RepositoryConfig string
	RepositoryCache  string
	RegistryConfig   string
}

type helmClient struct {
	Providers    getter.Providers
	ActionConfig *action.Configuration
	Settings     *cli.EnvSettings
	storage      *repo.File
}

func (c *helmClient) installChart(ctx context.Context, spec *ChartSpec) (*release.Release, error) {

	client := action.NewInstall(c.ActionConfig)
	setChartInstallOptions(spec, client)

	releaseName, _, err := client.NameAndChart([]string{spec.ChartName})
	if err != nil {
		return nil, err
	}

	client.ReleaseName = releaseName

	if client.Version == "" {
		client.Version = ">0.0.0-0"
	}

	helmChart, chartPath, err := c.getChart(spec.ChartName, &client.ChartPathOptions)
	if err != nil {
		return nil, err
	}

	helmChart, err = updateDependencies(c, helmChart, &client.ChartPathOptions, chartPath, client.DependencyUpdate, spec)

	return nil, nil
}

func (c *helmClient) getChart(name string, chartPathOptions *action.ChartPathOptions) (*chart.Chart, string, error) {
	chartPath, err := chartPathOptions.LocateChart(name, c.Settings)

	if err != nil {
		return nil, "", err
	}

	helmChart, err := loader.Load(chartPath)
	if err != nil {
		return nil, "", err
	}

	return helmChart, chartPath, nil
}

func (c *helmClient) chartExists(name string, namespace string) (bool, error) {
	releases, err := c.listReleases()
	if err != nil {
		return false, err
	}

	for _, r := range releases {
		if r.Name == name && r.Namespace == namespace {
			return true, nil
		}
	}
	return false, nil
}

func (c *helmClient) getRelease(name string) (*release.Release, error) {
	releaseClient := action.NewGet(c.ActionConfig)
	return releaseClient.Run(name)
}

func (c *helmClient) uninstallRelease(name string) error {
	releaseClient := action.NewUninstall(c.ActionConfig)

	resp, err := releaseClient.Run(name)
	if err != nil {
		return err
	}
	//TODO: Something...
	fmt.Printf("Uninstalled release %q\n", resp.Release.Name)
	return nil
}

func (c *helmClient) listReleases() ([]*release.Release, error) {
	listClient := action.NewList(c.ActionConfig)
	listClient.StateMask = action.ListAll
	return listClient.Run()
}

func updateDependencies(c *helmClient, helmChart *chart.Chart, chartPathOptions *action.ChartPathOptions, chartPath string,
	depUpdate bool, spec *ChartSpec) (*chart.Chart, error) {
	return nil, nil
}
