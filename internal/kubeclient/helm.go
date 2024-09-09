package kubeclient

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
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
	storage      *repo.File
}

func (c *helmClient) getChart(name string) (*chart.Chart, error) {
	return nil, nil
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
