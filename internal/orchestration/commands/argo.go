package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"github.com/joho/godotenv"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/maliciousbucket/plumage/internal/kubeclient"
	"github.com/maliciousbucket/plumage/internal/orchestration"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os/exec"
	"sync"
	"syscall"
	"time"
)

var (
	client argocd.Client
)

type ArgoClient interface {
	GetClusters(ctx context.Context) ([]v1alpha1.Cluster, error)
	CreateCluster(ctx context.Context, name string) (*v1alpha1.Cluster, error)
	AddProjectDestination(ctx context.Context, projectName string, server string, namespace string, name string) error
	AddApplicationDestination(ctx context.Context, appName string, server string, namespace string, name string) error
	CreateProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	DeleteProject(ctx context.Context, name string) error
	DeleteProjectWithApps(ctx context.Context, name string) error
	GetProject(ctx context.Context, name string) (*v1alpha1.AppProject, error)
	ListProjects(ctx context.Context) (*v1alpha1.AppProjectList, error)
	GetApplication(ctx context.Context, name string) (*v1alpha1.Application, error)
	ListApplications(ctx context.Context, params *argocd.AppQueryParams) (*v1alpha1.ApplicationList, error)
	CreateIngressApp(ctx context.Context) error
	CreateApplication(ctx context.Context) (*v1alpha1.Application, error)
	AddApplicationToProject(ctx context.Context, appName string, project string, validate bool) (*v1alpha1.ApplicationSpec, error)
	UpdateApplication(ctx context.Context, appName string) (*v1alpha1.Application, error)
	AddRepoCredentials(ctx context.Context) error
	SyncApplicationResources(ctx context.Context, name string) error
	SyncProject(ctx context.Context, name string) error

	CreateMonitoringProject(ctx context.Context) error
	CreateNetworkingProject(ctx context.Context) error
	CreateIngressProject(ctx context.Context) error
	CreateApplicationProject(ctx context.Context, app string) error
}

func SetArgoTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-argo-token",
		Short: "set argo token",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return newKubeClient()
		},
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			err := orchestration.SetArgoToken(ctx, kubernetesClient)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
}

// SetArgoToken TODO: Remove
func SetArgoToken() {
	setArgoToken()
}

func setArgoToken() {
	k8sClient, err := kubeclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = k8sClient.PatchArgoToLB(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.CreateGalahArgoAccount(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}
	secret, err := k8sClient.GetArgoPassword(ctx, "argocd")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)

	portForwardCmd := exec.Command("kubectl", "port-forward", "svc/argocd-helm-server", "8081:443", "-n", "argocd")
	go func() {
		defer wg.Done()
		err := portForwardCmd.Run()
		if err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				if exitErr.Sys().(syscall.WaitStatus).Signaled() && exitErr.Sys().(syscall.WaitStatus).Signal() == syscall.SIGKILL {
					log.Println("Port-forward Closed")
					return
				}
			}
			log.Fatal("Port-forward failed: ", err)
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Secret: " + secret)
	//pss := fmt.Sprintf("--password %s", secret)
	loginCommand := exec.Command("argocd", "login", "localhost:8081", "--username", "admin", "--password", secret, "--insecure")
	fmt.Println(loginCommand)
	aout, err := loginCommand.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err = loginCommand.Start(); err != nil {
		log.Fatal(err)
	}
	ad, err := io.ReadAll(aout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ad))

	tokenCmd := exec.Command("argocd", "account", "generate-token", "--account", "galah")
	fmt.Println(tokenCmd)
	stdout, err := tokenCmd.StdoutPipe()

	if err = tokenCmd.Start(); err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Token: %s\n", string(data))

	myEnv, err := godotenv.Read(".env")
	if err != nil {
		log.Fatal(err)
	}
	myEnv["ARGOCD_TOKEN"] = string(data)
	err = godotenv.Write(myEnv, ".env")
	if err != nil {
		log.Fatal(err)
	}

	if err := portForwardCmd.Process.Kill(); err != nil {
		log.Fatal("Failed to kill port-forward process: ", err)
	}

	wg.Wait()
}

func createAppsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create-apps",
	}
	return cmd
}

func ProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project",
		Short: "Manage ArgoCD projects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newArgoClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(createProjectCmd())
	cmd.AddCommand(createAppsCmd())
	cmd.AddCommand(getProjectCmd())
	cmd.AddCommand(addAppToProjectCmd())
	cmd.AddCommand(deleteProjectCmd())
	return cmd
}

func createProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create ArgoCD projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			monitoring, _ := cmd.Flags().GetBool("monitoring")
			networking, _ := cmd.Flags().GetBool("networking")
			crd, _ := cmd.Flags().GetBool("crd")
			gateway, _ := cmd.Flags().GetBool("gateway")
			appName, _ := cmd.Flags().GetBool("app")
			name, _ := cmd.Flags().GetString("name")
			ctx := context.Background()
			if monitoring {
				return argoClient.CreateMonitoringProject(ctx)
			}

			if networking {
				return argoClient.CreateNetworkingProject(ctx)
			}

			if crd {

			}

			if gateway {
				return argoClient.CreateIngressProject(ctx)
			}

			if appName {
				if name == "" {
					return fmt.Errorf("app name is required")
				}
				return argoClient.CreateApplicationProject(ctx, name)
			}
			return nil
		},
	}
	cmd.Flags().BoolP("monitoring", "m", false, "Create monitoring project")
	cmd.Flags().BoolP("networking", "n", false, "create networking project")
	cmd.Flags().BoolP("crd", "c", false, "create CRD project")
	cmd.Flags().BoolP("gateway", "g", false, "Create gateway project")
	cmd.Flags().BoolP("app", "p", false, "Create app project")
	cmd.Flags().StringP("name", "a", "", "App name")
	cmd.MarkFlagsMutuallyExclusive("monitoring", "networking", "crd", "gateway", "app")
	cmd.MarkFlagsRequiredTogether("name", "app")
	return cmd
}

func deleteProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete ArgoCD projects",
		RunE: func(cmd *cobra.Command, args []string) error {

			if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
				return err
			}
			withApps, _ := cmd.Flags().GetBool("withApps")
			project := args[0]
			if project == "" {
				return fmt.Errorf("project is required")
			}
			ctx := context.Background()
			if withApps {
				if promptUser("Are you sure?") {
					err := argoClient.DeleteProjectWithApps(ctx, project)
					if err != nil {
						log.Fatal(err)
					}
					log.Println("Successfully deleted ArgoCD project: ", project)
					return nil
				} else {
					log.Println("Aborting....")
					return nil
				}
			}

			if err := argoClient.DeleteProject(ctx, project); err != nil {
				return err
			}
			log.Printf("\nProject %s deleted", project)
			return nil
		},
	}
	cmd.Flags().BoolP("withApps", "a", false, "Delete all of the project's apps")

	return cmd
}

func getProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get ArgoCD projects",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newArgoClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			list, _ := cmd.Flags().GetBool("list")
			ctx := context.Background()
			if list {
				projects, err := argoClient.ListProjects(ctx)
				if err != nil {
					return err
				}
				for _, proj := range projects.Items {
					fmt.Printf("%s\n%s\n%+v\n", proj.Name, proj.Spec.Description, proj.Spec.SourceRepos)
				}
				return nil
			}
			name := cmd.Flag("name").Value.String()
			if name == "" {
				return errors.New("project name is required")
			}
			project, err := argoClient.GetProject(ctx, name)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n%s\n%+v\n", project.Name, project.Spec.Description, project.Spec.SourceRepos)

			return nil
		},
	}
	cmd.Flags().BoolP("list", "l", false, "List ArgoCD projects")
	cmd.Flags().StringP("name", "n", "", "Project name to search for")
	cmd.MarkFlagsMutuallyExclusive("name", "list")
	return cmd
}

func DeployAppCmd(filePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Deploy synthesised applications",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return newArgoClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			appName, err := orchestration.DeployApp(ctx, argoClient, filePath)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("\nApplication %s deployed successfully", appName)
			}

			return nil
		},
	}
	return cmd
}

func addAppToProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add an ArgoCD application to a Project",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}
			application := cmd.Flag("app").Value.String()
			project := cmd.Flag("project").Value.String()
			if application == "" || project == "" {
				return errors.New("both application and project names are required")
			}
			ctx := context.Background()
			_, err := argoClient.AddApplicationToProject(ctx, application, project, true)
			if err != nil {
				return err
			}
			log.Printf("Added %s to Project %s", application, project)
			return nil
		},
	}
	cmd.Flags().StringP("app", "a", "", "App name")
	cmd.Flags().StringP("project", "p", "", "Project name")
	_ = cmd.MarkFlagRequired("app")
	cmd.MarkFlagsRequiredTogether("app", "project")
	return cmd
}

var (
	syncProject bool
	syncApp     bool
)

func SyncCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync ArgoCD resources",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return newArgoClient()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}

			monitoring, _ := cmd.Flags().GetBool("monitoring")
			gateway, _ := cmd.Flags().GetBool("gateway")
			dashboards, _ := cmd.Flags().GetBool("dashboards")
			deployedService, _ := cmd.Flags().GetBool("deployedService")

			if monitoring {
				return syncArgoProject(argoClient, "galah-monitoring")
			}

			if gateway {
				return syncArgoProject(argoClient, "ingress")
			}

			if dashboards {
				return syncApplication(argoClient, "infra-routes")
			}

			if deployedService {
				if args[0] == "" {
					return errors.New("deployed service name is required")
				}
				return syncArgoProject(argoClient, args[0])
			}

			if syncProject {
				if args[0] == "" {
					return errors.New("project name is required")
				}
				return syncArgoProject(argoClient, args[0])
			}

			if syncApp {
				if args[0] == "" {
					return errors.New("app name is required")
				}
				return syncApplication(argoClient, args[0])
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&syncProject, "project", "p", false, "Project name")
	cmd.Flags().BoolVarP(&syncApp, "app", "a", false, "Application name")
	cmd.Flags().BoolP("monitoring", "m", false, "Sync monitoring project")
	cmd.Flags().BoolP("gateway", "g", false, "Sync gateway project")
	cmd.Flags().BoolP("dashboards", "d", false, "Sync dashboards project")
	//E.g. chirp
	cmd.Flags().BoolP("deployedService", "s", false, "Service name")
	cmd.MarkFlagsMutuallyExclusive("project", "app", "monitoring", "dashboards", "deployedService", "gateway")
	cmd.MarkFlagsOneRequired("project", "app", "monitoring", "dashboards", "deployedService", "gateway")

	cmd.AddCommand(syncAllCmd())

	return cmd

}

func syncApplication(client ArgoClient, name string) error {
	ctx := context.Background()
	if err := client.SyncApplicationResources(ctx, name); err != nil {
		return err
	}
	return nil
}

func syncArgoProject(client ArgoClient, name string) error {
	ctx := context.Background()
	if err := client.SyncProject(ctx, name); err != nil {
		return err
	}
	return nil
}

func syncAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "Sync all ArgoCD resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			if err := argoClient.SyncAllProjects(ctx); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
