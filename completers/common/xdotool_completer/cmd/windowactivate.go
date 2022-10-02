package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowactivateCmd = &cobra.Command{
	Use:   "windowactivate",
	Short: "Activate the window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowactivateCmd).Standalone()

	windowactivateCmd.Flags().Bool("sync", false, "After sending the window activation, wait until the window is actually activated")
	rootCmd.AddCommand(windowactivateCmd)

	carapace.Gen(windowactivateCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
