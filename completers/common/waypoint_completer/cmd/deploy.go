package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy an application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	deployCmd.Flags().Bool("release", false, "Release this deployment immediately.")

	addGlobalOptions(deployCmd)
	addOperationOptions(deployCmd)

	rootCmd.AddCommand(deployCmd)
}
