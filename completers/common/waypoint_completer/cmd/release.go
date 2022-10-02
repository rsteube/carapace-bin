package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Open a deployment to traffic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.Flags().String("deployment", "", "Release the specified deployment.")
	releaseCmd.Flags().Bool("prune", false, "Prune old unreleased deployments.")
	releaseCmd.Flags().String("prune-retain", "", "The number of unreleased deployments to keep.")
	releaseCmd.Flags().Bool("repeat", false, "Re-release if deploy is already released.")

	addGlobalOptions(releaseCmd)
	addOperationOptions(releaseCmd)

	rootCmd.AddCommand(releaseCmd)
}
