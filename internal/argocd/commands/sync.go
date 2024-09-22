package commands

import (
	"context"
	"github.com/maliciousbucket/plumage/internal/argocd"
	"github.com/spf13/cobra"
)

type ArgoSyncClient interface {
	SyncApplicationResources(ctx context.Context, name string) error
	SyncProject(ctx context.Context, name string) error
}

var (
	syncProject bool
	syncApp     bool
)

func SyncCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sync",
		Short: "Sync ArgoCD resources",
		Args:  cobra.MinimumNArgs(1),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			_, err := newClient()
			if err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateFlagGroups(); err != nil {
				return err
			}

			conn, err := argocd.GetConnection()
			if err != nil {
				return err
			}
			client, err := argocd.NewClient(conn)
			if err != nil {
				return err
			}

			if syncProject {
				return syncArgoProject(client, args[0])
			}

			if syncApp {
				return syncApplication(client, args[0])
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&syncProject, "project", "p", false, "project name")
	cmd.Flags().BoolVarP(&syncApp, "app", "a", false, "application name")
	cmd.MarkFlagsMutuallyExclusive("project", "app")
	cmd.MarkFlagsOneRequired("project", "app")

	return cmd

}

func syncApplication(client ArgoSyncClient, name string) error {
	ctx := context.Background()
	if err := client.SyncApplicationResources(ctx, name); err != nil {
		return err
	}
	return nil
}

func syncArgoProject(client ArgoSyncClient, name string) error {
	ctx := context.Background()
	if err := client.SyncProject(ctx, name); err != nil {
		return err
	}
	return nil
}
