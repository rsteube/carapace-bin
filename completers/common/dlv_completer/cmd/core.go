package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Examine a core dump.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coreCmd).Standalone()
	rootCmd.AddCommand(coreCmd)

	carapace.Gen(coreCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
