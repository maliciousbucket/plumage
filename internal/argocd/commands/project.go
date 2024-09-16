package commands

import (
	"context"
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
			del, _ := cmd.Flags().GetBool("delete")
			if get {
				if len(args) == 0 {
					log.Fatalln(fmt.Errorf("project name must be specified"))
				}
				err := getProject(args[0])
				if err != nil {
					log.Fatalln(err)
				}
			}
			if del {
				if len(args) == 0 {
					log.Fatalln(fmt.Errorf("project name must be specified"))
				}
				err := deleteProject(args[0])
				if err != nil {
					log.Fatalln(err)
				}
			}

		},
	}
	cmd.Flags().BoolP("get", "g", false, "get a project by name")
	cmd.Flags().BoolP("delete", "d", false, "delete a project")
	cmd.MarkFlagsMutuallyExclusive("get", "delete")
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
	ctx := context.Background()
	res, err := client.GetProject(ctx, name)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func deleteProject(name string) error {
	conn, err := argocd.GetConnection()
	if err != nil {
		return err
	}
	client, err := argocd.NewClient(conn)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = client.DeleteProject(ctx, name)
	if err != nil {
		return fmt.Errorf("failed to delete project %s: %v", name, err)
	}
	return nil
}
