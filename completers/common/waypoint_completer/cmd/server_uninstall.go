package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall the Waypoint server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_uninstallCmd).Standalone()

	server_uninstallCmd.Flags().Bool("auto-approve", false, "Auto-approve server uninstallation.")
	server_uninstallCmd.Flags().Bool("delete-context", false, "Delete the context for the server once it's uninstalled.")
	server_uninstallCmd.Flags().String("ecs-cluster", "", "Configures the Cluster to uninstall.")
	server_uninstallCmd.Flags().String("ecs-region", "", "Configures which AWS region to uninstall from.")
	server_uninstallCmd.Flags().Bool("ignore-runner-error", false, "Ignore any errors encountered while uninstalling runners.")
	server_uninstallCmd.Flags().String("k8s-context", "", "The Kubernetes context to unisntall the Waypoint server from.")
	server_uninstallCmd.Flags().String("k8s-namespace", "", "Namespace in Kubernetes to uninstall the Waypoint server from.")
	server_uninstallCmd.Flags().String("platform", "", "Platform to uninstall the Waypoint server from.")
	server_uninstallCmd.Flags().Bool("snapshot", false, "Enable or disable taking a snapshot of Waypoint server prior to uninstall.")
	server_uninstallCmd.Flags().String("snapshot-name", "", "Filename to write the snapshot to.")

	addGlobalOptions(server_uninstallCmd)

	serverCmd.AddCommand(server_uninstallCmd)
}
