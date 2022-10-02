package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowfocusCmd = &cobra.Command{
	Use:   "windowfocus",
	Short: "Focus a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowfocusCmd).Standalone()

	windowfocusCmd.Flags().Bool("sync", false, "After sending the window focus request, wait until the window is actually focused")
	rootCmd.AddCommand(windowfocusCmd)

	carapace.Gen(windowfocusCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
