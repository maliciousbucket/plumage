package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/spf13/cobra"
	"log"
)

var (
	name      = ""
	nameSpace = ""
	project   = ""
	repo      = ""
)

type ArgoAppClient interface {
	ListApplications(ctx context.Context, params *argocd.AppQueryParams) (*v1alpha1.ApplicationList, error)
	GetApplication(ctx context.Context, name string) (*v1alpha1.Application, error)
}

func ArgoApplicationCmd(client ArgoAppClient) *cobra.Command {
	appCmd := &cobra.Command{
		Use:   "app",
		Short: "Manage applications",
		Run: func(cmd *cobra.Command, args []string) {
			get, _ := cmd.Flags().GetBool("get")
			list, _ := cmd.Flags().GetBool("list")
			if get {
				if len(args) == 0 {
					log.Fatalln(errors.New("no application name specified"))
				}
				err := getAppByName(client, args[0])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if list {
				err := listApps(client)
				if err != nil {
					log.Fatalln(err)
				}
			}
		},
	}
	appCmd.Flags().BoolP("get", "g", false, "Get application")
	appCmd.Flags().BoolP("list", "l", false, "List applications")
	appCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "", "Namespace for application")
	appCmd.Flags().StringVarP(&project, "project", "p", "", "Project for application")
	appCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository for application")
	appCmd.Flags().StringVarP(&name, "name", "a", "", "Name for application")

	appCmd.MarkFlagsMutuallyExclusive("get", "list")

	return appCmd
}

func getAppByName(client ArgoAppClient, name string) error {
	ctx := context.Background()
	app, err := client.GetApplication(ctx, name)
	if err != nil {
		return err
	}
	fmt.Println(app)
	return nil

}

func listApps(client ArgoAppClient) error {
	params := argocd.AppQueryParams{
		Options: []argocd.AppQueryFunc{},
	}

	buildAppQueryParams(&params)
	ctx := context.Background()
	apps, err := client.ListApplications(ctx, &params)
	if err != nil {
		return err
	}
	fmt.Println(apps)
	return nil
}

var (
	createAppNamespace = ""
	createAppName      = ""
)

func createAppCmd(client ArgoAppClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new ArgoCD application",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}

		},
	}
	cmd.Flags().StringVarP(&createAppNamespace, "namespace", "n", "", "Namespace for application")
	cmd.Flags().StringVarP(&createAppName, "name", "a", "", "Name for application")
	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")
	return cmd
}

func buildAppQueryParams(params *argocd.AppQueryParams) {
	if project != "" {
		params.Options = append(params.Options, argocd.WithProject(project))
	}
	if repo != "" {
		params.Options = append(params.Options, argocd.WithRepository(repo))
	}
	if name != "" {
		params.Options = append(params.Options, argocd.WithName(name))
	}
	if nameSpace != "" {
		params.Options = append(params.Options, argocd.WithNamespace(nameSpace))
	}
}
