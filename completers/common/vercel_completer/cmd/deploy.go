package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Performs a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployCmd).Standalone()

	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
