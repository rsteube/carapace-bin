package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowkillCmd = &cobra.Command{
	Use:   "windowkill",
	Short: "Kill a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowkillCmd).Standalone()

	rootCmd.AddCommand(windowkillCmd)

	carapace.Gen(windowkillCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
