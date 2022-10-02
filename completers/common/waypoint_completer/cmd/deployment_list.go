package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deployment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployment_listCmd).Standalone()

	deployment_listCmd.Flags().Bool("json", false, "Output the deployment information as JSON.")
	deployment_listCmd.Flags().Bool("long-ids", false, "Show long identifiers rather than sequence numbers.")
	deployment_listCmd.Flags().Bool("url", false, "Display deployment URL.")
	deployment_listCmd.Flags().Bool("verbose", false, "Display more details about each deployment.")
	deployment_listCmd.Flags().Bool("workspace-all", false, "List builds in all workspaces for this project and application.")

	addGlobalOptions(deployment_listCmd)
	addFilterOptions(deployment_listCmd)

	deploymentCmd.AddCommand(deployment_listCmd)
}
