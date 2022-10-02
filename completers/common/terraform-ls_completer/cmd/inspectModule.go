package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var inspectModuleCmd = &cobra.Command{
	Use:   "inspect-module",
	Short: "Lists available debug items",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectModuleCmd).Standalone()
	inspectModuleCmd.Flags().BoolS("verbose", "verbose", false, "whether to enable verbose output")

	rootCmd.AddCommand(inspectModuleCmd)

	carapace.Gen(inspectModuleCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
