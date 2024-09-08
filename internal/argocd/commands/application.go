package commands

import (
	"errors"
	"fmt"
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

func ArgoApplicationCmd() *cobra.Command {
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
				err := getAppByName(args[0])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if list {
				err := listApps()
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

func getAppByName(name string) error {
	conn, err := argocd.GetConnection()
	if err != nil {
		return err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return err
	}
	app, err := client.GetApplication(name)
	if err != nil {
		return err
	}
	fmt.Println(app)
	return nil

}

func listApps() error {
	conn, err := argocd.GetConnection()
	if err != nil {
		return err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return err
	}
	params := argocd.AppQueryParams{
		Options: []argocd.AppQueryFunc{},
	}

	buildAppQueryParams(&params)
	apps, err := client.ListApplications(&params)
	if err != nil {
		return err
	}
	fmt.Println(apps)
	return nil
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
