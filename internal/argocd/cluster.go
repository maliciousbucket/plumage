package argocd

import (
	"context"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/cluster"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
)

func (c *Client) GetClusters(ctx context.Context) ([]v1alpha1.Cluster, error) {
	cl, err := c.clusterClient.List(ctx, &cluster.ClusterQuery{})
	if err != nil {
		return nil, err
	}

	return cl.Items, nil
}

func (c *Client) CreateCluster(ctx context.Context, name string) (*v1alpha1.Cluster, error) {
	cl, err := c.clusterClient.Create(ctx, &cluster.ClusterCreateRequest{
		Cluster: &v1alpha1.Cluster{
			Name:             name,
			Config:           v1alpha1.ClusterConfig{},
			Namespaces:       nil,
			Info:             v1alpha1.ClusterInfo{},
			ClusterResources: false,
		},
		Upsert: false,
	})
	if err != nil {
		return nil, err
	}
	return cl, nil
}
