package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Perform the build, deploy, and release steps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upCmd).Standalone()

	upCmd.Flags().Bool("prune", false, "Prune old unreleased deployments.")
	upCmd.Flags().String("prune-retain", "", "The number of unreleased deployments to keep.")

	addGlobalOptions(upCmd)
	addOperationOptions(upCmd)

	rootCmd.AddCommand(upCmd)
}
