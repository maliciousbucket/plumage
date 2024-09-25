package commands

import (
	"context"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/spf13/cobra"
	"log"
)

type ArgoProjectClient interface {
	CreateProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	AddApplicationToProject(ctx context.Context, appName string, project string, validate bool) (*v1alpha1.ApplicationSpec, error)
	GetProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
}

var (
	projectName = ""
)

func ArgoProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage ArgoCD projects",
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			_, err := newClient()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			getProj, _ := cmd.Flags().GetBool("get")
			del, _ := cmd.Flags().GetBool("delete")
			createProj, _ := cmd.Flags().GetBool("create")

			if projectName == "" {
				return fmt.Errorf("project name is required")
			}

			if getProj {
				err := getProject(projectName)
				if err != nil {
					return err
				}
			}
			if del {
				err := deleteProject(projectName)
				if err != nil {
					return err
				}
			}
			if createProj {
				if err := createProject(argoClient, projectName); err != nil {
					return err
				}
			}

			return nil

		},
	}
	cmd.Flags().BoolP("get", "g", false, "get a project by name")
	cmd.Flags().BoolP("delete", "d", false, "delete a project")
	cmd.Flags().BoolP("create", "c", false, "create a project")
	cmd.Flags().StringVarP(&projectName, "name", "n", "", "create a project by name")
	cmd.MarkFlagsMutuallyExclusive("get", "delete")

	cmd.AddCommand(addAppToProjectCmd(argoClient))
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

func createProject(client ArgoProjectClient, name string) error {
	ctx := context.Background()
	proj, err := client.CreateProject(ctx, name)
	if err != nil {
		return err
	}
	fmt.Printf("Project %s created\n", proj.Name)
	return nil
}

var (
	addAppName    = ""
	addAppProject = ""
)

func addAppToProjectCmd(client ArgoProjectClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an application to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			ctx := context.Background()
			result, err := client.AddApplicationToProject(ctx, addAppName, addAppProject, true)
			if err != nil {
				return fmt.Errorf("failed to add application to project %s: %v", addAppName, err)
			}

			fmt.Printf("Successfully added %s to project %s\n", addAppName, result.Project)

			return nil
		},
	}
	cmd.Flags().StringVar(&addAppName, "app", "", "application name")
	cmd.Flags().StringVar(&addAppProject, "project", "", "project name")
	cmd.MarkFlagsRequiredTogether("project", "app")
	if err := cmd.MarkFlagRequired("app"); err != nil {
		log.Fatalln(fmt.Errorf("failed to mark flag as required: %v", err))
	}
	if err := cmd.MarkFlagRequired("project"); err != nil {
		log.Fatalln(fmt.Errorf("failed to mark flag as required: %v", err))
	}

	return cmd
}

func infraProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "infra",
		Short: "Manage ArgoCD infra projects",
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			_, err := newClient()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	cmd.Flags().BoolP("monitoring", "m", false, "create monitoring project")
	cmd.Flags().BoolP("networking", "n", false, "create networking project")
	cmd.Flags().BoolP("crd", "c", false, "create CRD project")
	cmd.Flags().BoolP("gateway", "g", false, "create gateway project")
	return cmd
}
