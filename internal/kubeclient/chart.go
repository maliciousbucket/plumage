package kubeclient

import "helm.sh/helm/v3/pkg/action"

type ChartSpec struct {
	ReleaseName     string
	Version         string
	ChartName       string
	Namespace       string
	CreateNamespace bool
	IncludeCRDs     bool
	UpgradeCRDs     bool
	Wait            bool
}

func setChartInstallOptions(spec *ChartSpec, opts *action.Install) {
	opts.ReleaseName = spec.ReleaseName
	opts.Version = spec.Version
	opts.IncludeCRDs = spec.IncludeCRDs
	opts.CreateNamespace = spec.CreateNamespace
	opts.Wait = spec.Wait
}
