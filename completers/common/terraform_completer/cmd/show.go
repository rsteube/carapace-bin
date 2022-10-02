package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current state or a saved plan",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolS("json", "json", false, "output in a machine-readable form")
	showCmd.Flags().BoolS("no-color", "no-color", false, "output without color")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
