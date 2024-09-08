package commands

import (
	"fmt"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/spf13/cobra"
	"log"
)

func ArgoProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage ArgoCD projects",
		Run: func(cmd *cobra.Command, args []string) {
			get, _ := cmd.Flags().GetBool("get")
			if get {
				if len(args) == 0 {
					log.Fatalln(fmt.Errorf("project name must be specified"))
				}
				err := getProject(args[0])
				if err != nil {
					log.Fatalln(err)
				}
			}

		},
	}
	cmd.Flags().BoolP("get", "g", false, "get a project by name")
	return cmd
}

func getProject(name string) error {
	conn, err := argocd.GetConnection()
	if err != nil {
		return err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return err
	}
	project, err := client.GetProject(name)
	if err != nil {
		return err
	}
	fmt.Println(project)
	return nil
}
