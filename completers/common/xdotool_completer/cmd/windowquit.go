package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowquitCmd = &cobra.Command{
	Use:   "windowquit",
	Short: "Close a window gracefully",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowquitCmd).Standalone()

	rootCmd.AddCommand(windowquitCmd)

	carapace.Gen(windowquitCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
