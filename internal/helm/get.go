package helm

import (
	"fmt"
	"helm.sh/helm/v3/pkg/release"
)

type ChartRelease struct {
	ReleaseMeta *ReleaseMeta
	Info        *release.Info
	Labels      map[string]string
}

type ReleaseMeta struct {
	Name      string
	Namespace string
	Version   int
}

func (c *helmClient) GetRelease(name string) (*ChartRelease, error) {
	return c.getRelease(name)
}

func (c *helmClient) getRelease(name string) (*ChartRelease, error) {
	res, err := c.Client.GetRelease(name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving name %s: %s", name, err)
	}
	if res == nil {
		return nil, fmt.Errorf("release %s not found", name)
	}
	info := &ChartRelease{
		ReleaseMeta: &ReleaseMeta{
			Name:      res.Name,
			Namespace: res.Namespace,
			Version:   res.Version,
		},
		Info:   res.Info,
		Labels: res.Labels,
	}
	return info, nil
}

func (c *helmClient) ListReleases() ([]*ReleaseMeta, error) {
	return c.listReleases()
}

func (c *helmClient) listReleases() ([]*ReleaseMeta, error) {
	releases, err := c.Client.ListDeployedReleases()
	if err != nil {
		return nil, fmt.Errorf("error listing releases: %s", err)
	}
	if releases == nil || len(releases) == 0 {
		return nil, fmt.Errorf("no releases found")
	}
	result := make([]*ReleaseMeta, 0, len(releases))
	for _, r := range releases {
		result = append(result, &ReleaseMeta{
			Name:      r.Name,
			Namespace: r.Namespace,
			Version:   r.Version,
		})
	}
	return result, nil
}

func (c *helmClient) GetReleaseValues(name string, allValues bool) (map[string]interface{}, error) {
	return c.getReleaseValues(name, allValues)
}

func (c *helmClient) getReleaseValues(name string, allValues bool) (map[string]interface{}, error) {
	res, err := c.Client.GetReleaseValues(name, allValues)
	if err != nil {
		return nil, fmt.Errorf("error retrieving release %s: %s", name, err)
	}
	if res == nil {
		return nil, fmt.Errorf("release %s not found", name)
	}

	return res, nil
}
