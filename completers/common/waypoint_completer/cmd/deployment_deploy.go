package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deployment_deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a pushed artifact",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployment_deployCmd).Standalone()

	deployment_deployCmd.Flags().Bool("release", false, "Release this deployment immediately.")

	addGlobalOptions(deployment_deployCmd)
	addOperationOptions(deployment_deployCmd)

	deploymentCmd.AddCommand(deployment_deployCmd)
}
