package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowsizeCmd = &cobra.Command{
	Use:   "windowsize",
	Short: "Set the window size of the given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowsizeCmd).Standalone()

	windowsizeCmd.Flags().Bool("sync", false, "After sending the window size request, wait until the window is actually resized")
	windowsizeCmd.Flags().Bool("usehints", false, "Use window sizing hints (when available) to set width and height")
	rootCmd.AddCommand(windowsizeCmd)

	carapace.Gen(windowsizeCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
