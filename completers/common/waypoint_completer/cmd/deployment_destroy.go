package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deployment_destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "estroy one or more deployments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployment_destroyCmd).Standalone()

	deployment_destroyCmd.Flags().Bool("all", false, "Delete ALL deployments, including released.")
	deployment_destroyCmd.Flags().Bool("force", false, "Yes to all confirmations.")

	addGlobalOptions(deployment_destroyCmd)
	addOperationOptions(deployment_destroyCmd)

	deploymentCmd.AddCommand(deployment_destroyCmd)
}
