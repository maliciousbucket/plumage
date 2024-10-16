package helm

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	helmc "github.com/mittwald/go-helm-client"
	"github.com/mittwald/go-helm-client/values"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/repo"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type ChartsConfig struct {
	ConfigDir string         `yaml:"configDir"`
	FileName  string         `yaml:"fileName"`
	Charts    []*ChartConfig `yaml:"charts"`
}

type ChartConfig struct {
	Repository   string `yaml:"repository"`
	remoteFile   bool
	Namespace    string                 `yaml:"namespace"`
	Name         string                 `yaml:"chartName"`
	ReleaseName  string                 `yaml:"releaseName"`
	Version      string                 `yaml:"version"`
	Replace      bool                   `yaml:"replace"`
	ValuesFiles  []string               `yaml:"valuesFiles"`
	Values       map[string]interface{} `yaml:"values"`
	valuesString []string
	LocalFile    bool              `yaml:"localFile"`
	SkipCRDs     bool              `yaml:"skipCRDs"`
	UpgradeCRDs  bool              `yaml:"upgradeCRDs"`
	Labels       map[string]string `yaml:"labels"`
	Lint         bool              `yaml:"lint"`
}

func InstallCRDChartsFromConfig(ctx context.Context, client *helmClient, cfg *ChartsConfig) error {
	var chartClient *helmClient
	if client == nil {
		clientCfg := &ClientCfg{}
		newHelmClient, err := newClient(clientCfg)
		if err != nil {
			return err
		}
		chartClient = newHelmClient
	} else {
		chartClient = client
	}

	if cfg.ConfigDir != "" && cfg.FileName != "" {
		path := filepath.Join(cfg.ConfigDir, cfg.FileName)
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			return fmt.Errorf("file %s does not exist", path)
		}
		if info.IsDir() {
			return fmt.Errorf("file %s is a directory", path)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if cfg.Charts == nil {
			cfg.Charts = []*ChartConfig{}
		}
		var charts []*ChartConfig
		err = yaml.Unmarshal(data, &charts)
		if err != nil {
			return err
		}
		cfg.Charts = append(cfg.Charts, charts...)
	}
	for _, chart := range cfg.Charts {
		if err := chartClient.installChart(ctx, chart); err != nil {
			return err
		}
	}
	log.Println("Successfully installed CRD charts")
	return nil
}

func InstallCRDChartsFromFile(ctx context.Context, client *helmClient, configDir, fileName string) error {
	var chartClient *helmClient
	if client == nil {
		cfg := &ClientCfg{}
		newHelmClient, err := newClient(cfg)
		if err != nil {
			return err
		}
		chartClient = newHelmClient
	} else {
		chartClient = client
	}

	charts, err := loadChartConfigs(configDir, fileName)
	if err != nil {
		return err
	}

	for _, chart := range charts {
		if err = chartClient.installChart(ctx, chart); err != nil {
			return err
		}
	}
	log.Println("Successfully installed Helm CRD charts")
	return nil
}

func (c *helmClient) InstallChart(ctx context.Context, chart *ChartConfig) error {
	return c.installChart(ctx, chart)
}

func (c *helmClient) installChart(ctx context.Context, chart *ChartConfig) error {
	opts := getChartOpts(chart)
	spec, err := chart.chartSpec(c, opts...)
	if err != nil {
		return err
	}

	res, err := c.Client.InstallOrUpgradeChart(ctx, spec, &helmc.GenericHelmOptions{
		RollBack: c.Client,
	})
	if err != nil {
		return err
	}

	log.Printf("Successfully installed helm chart %s in namespace %s\n", res.Name, res.Namespace)
	return nil
}

func loadChartConfigs(configDir, file string) ([]*ChartConfig, error) {
	path := filepath.Join(configDir, file)
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, fmt.Errorf("%s is a directory", path)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var chartsConfig ChartsConfig
	if err = yaml.Unmarshal(data, &chartsConfig); err != nil {
		return nil, err
	}

	if chartsConfig.ConfigDir != "" && chartsConfig.FileName != "" {
		extraPath := filepath.Join(configDir, chartsConfig.FileName)
		info, err = os.Stat(extraPath)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			return nil, fmt.Errorf("%s is a directory", chartsConfig.FileName)
		}
		chartData, readErr := os.ReadFile(extraPath)
		if readErr != nil {
			return nil, readErr
		}
		if chartsConfig.Charts == nil {
			chartsConfig.Charts = []*ChartConfig{}
		}
		var extraCharts []*ChartConfig
		if err = yaml.Unmarshal(chartData, &extraCharts); err != nil {
			return nil, err
		}
		chartsConfig.Charts = append(chartsConfig.Charts, extraCharts...)

	}

	return chartsConfig.Charts, nil

}

func getChartOpts(chart *ChartConfig) []chartOpts {
	opts := []chartOpts{}
	if chart.LocalFile {
		opts = append(opts, fromLocalArchive(chart.Repository))
	} else {
		opts = append(opts, fromRemote(chart.Repository))
	}
	opts = append(opts, withCRDOpts(chart.SkipCRDs, chart.UpgradeCRDs))

	if chart.ValuesFiles != nil && len(chart.ValuesFiles) > 0 {
		opts = append(opts, withValuesFiles(chart.ValuesFiles))
	}

	if chart.Values != nil {
		opts = append(opts, withValues(chart.Values))
	}
	return opts
}

func (chart *ChartConfig) chartSpec(client *helmClient, opts ...chartOpts) (*helmc.ChartSpec, error) {
	var chartConfigErr error
	if opts != nil {
		for _, opt := range opts {
			err := opt(chart)
			if err != nil {
				chartConfigErr = errors.Join(chartConfigErr, err)
			}
		}
	}

	if chartConfigErr != nil {
		return nil, chartConfigErr
	}
	if chart.LocalFile == false {
		if chart.remoteFile == false {
			repoErr := fromRemoteRepository(client)(chart)
			if repoErr != nil {
				return nil, repoErr
			}
		} else {
			fileErr := fromRemoteFile()(chart)
			if fileErr != nil {
				return nil, fileErr
			}
		}
	} else {
		chart.Name = chart.Repository
	}

	generateNamespace := false
	if chart.Namespace == "" {
		generateNamespace = true
	}

	valuesOpts := values.Options{}

	if chart.valuesString != nil && len(chart.valuesString) > 0 {
		valuesOpts.Values = chart.valuesString
	}

	if chart.ValuesFiles != nil && len(chart.ValuesFiles) > 0 {
		valuesOpts.ValueFiles = chart.ValuesFiles
	}

	spec := &helmc.ChartSpec{
		ChartName:       chart.Name,
		ReleaseName:     chart.ReleaseName,
		Namespace:       chart.Namespace,
		Version:         chart.Version,
		Replace:         chart.Replace,
		CleanupOnFail:   true,
		Labels:          chart.Labels,
		CreateNamespace: false,
		GenerateName:    generateNamespace,
		ValuesOptions:   valuesOpts,
	}

	if chart.Lint {
		if err := client.Client.LintChart(spec); err != nil {
			return nil, err
		}
	}
	return spec, nil
}

type chartOpts func(chart *ChartConfig) error

func withValuesFiles(files []string) chartOpts {
	return func(chart *ChartConfig) error {
		result := []string{}
		for _, file := range files {
			if file == "" {
				continue
			}
			info, err := os.Stat(file)
			if err != nil {
				return fmt.Errorf("failed to stat values file %s: %w", file, err)
			}
			if info.IsDir() {
				return fmt.Errorf("values file %s is a directory", file)
			}
			result = append(result, file)
		}
		chart.ValuesFiles = result
		return nil
	}
}

func fromLocalArchive(filepath string) chartOpts {
	return func(chart *ChartConfig) error {
		info, err := os.Stat(filepath)
		if err != nil {
			return fmt.Errorf("failed to stat %s chart file %s: %w", chart.Name, filepath, err)
		}

		if info.IsDir() {
			return fmt.Errorf("chart file %s is a directory", filepath)
		}
		chart.Repository = filepath

		return nil
	}
}

func fromRemote(repository string) chartOpts {
	return func(chart *ChartConfig) error {
		if repository == "" {
			return fmt.Errorf("repository is required. Chart: %s", chart.Name)
		}
		remote, err := url.Parse(repository)
		if err != nil {
			return fmt.Errorf("failed to parse repository %s: %w", repository, err)
		}

		if strings.HasSuffix(remote.Path, "gz") || strings.HasSuffix(remote.Path, "yaml") || strings.HasSuffix(remote.Path, "yml") {
			chart.remoteFile = true
		}

		if err != nil {
			return err
		}
		return nil
	}
}

func fromRemoteFile() chartOpts {
	return func(chart *ChartConfig) error {
		if !strings.HasSuffix(chart.Repository, "gz") && !strings.HasSuffix(chart.Repository, ".yaml") && !strings.HasSuffix(chart.Repository, ".yml") {
			return fmt.Errorf("repository %s does not contain a .gz, .yaml, or .yml file", chart.Repository)
		}

		chart.Name = chart.Repository
		return nil
	}
}

func fromRemoteRepository(c *helmClient) chartOpts {
	return func(chart *ChartConfig) error {
		err := c.addChartRepo(chart.Name, chart.Repository)
		if err != nil {
			return err
		}
		chartName := fmt.Sprintf("%s/%s", chart.Name, chart.ReleaseName)
		chart.Name = chartName
		return nil
	}
}

func withCRDOpts(skip, upgrade bool) chartOpts {
	return func(chart *ChartConfig) error {
		if skip && upgrade {
			return fmt.Errorf("only one of skip or upgrade CRDs can be enabled")
		}
		chart.SkipCRDs = skip
		chart.UpgradeCRDs = upgrade
		return nil
	}
}

func withValues(values map[string]interface{}) chartOpts {
	return func(chart *ChartConfig) error {
		result := []string{}

		var formatValues func(prefix string, v interface{}) error
		formatValues = func(prefix string, v interface{}) error {
			switch val := v.(type) {
			case map[string]interface{}:
				for nestedKey, nestedValue := range val {
					err := formatValues(prefix+"."+nestedKey, nestedValue)
					if err != nil {
						return err
					}
				}
			case []interface{}:
				for i, elem := range val {
					err := formatValues(fmt.Sprintf("%s[%d]", prefix, i), elem)
					if err != nil {
						return err
					}
				}
			case []map[string]interface{}:
				for i, elem := range val {
					for nestedKey, nestedValue := range elem {
						err := formatValues(fmt.Sprintf("%s[%d].%s", prefix, i, nestedKey), nestedValue)
						if err != nil {
							return err
						}
					}
				}
			default:
				b := new(bytes.Buffer)
				_, err := fmt.Fprintf(b, "%s=%v", prefix, val)
				if err != nil {
					return err
				}
				result = append(result, b.String())
			}
			return nil
		}

		for k, v := range values {
			err := formatValues(k, v)
			if err != nil {
				return err
			}
		}
		chart.valuesString = result
		return nil
	}
}

func (c *helmClient) addChartRepo(name, repository string) error {
	if repository == "" {
		return fmt.Errorf("repository address must be specified")
	}

	if name == "" {
		return fmt.Errorf("chart name must be specified")
	}

	chartRepo := repo.Entry{Name: name, URL: repository}

	if err := c.Client.AddOrUpdateChartRepo(chartRepo); err != nil {
		return fmt.Errorf("failed to add chart repo %s: %w", chartRepo.Name, err)
	}
	if err := c.Client.UpdateChartRepos(); err != nil {
		return fmt.Errorf("failed to update chart repos: %w", err)
	}

	log.Printf("\nchart repo added %s\n", chartRepo.Name)
	return nil

}
