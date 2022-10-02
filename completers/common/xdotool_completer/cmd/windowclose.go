package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowcloseCmd = &cobra.Command{
	Use:   "windowclose",
	Short: "Close a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowcloseCmd).Standalone()

	rootCmd.AddCommand(windowcloseCmd)

	carapace.Gen(windowcloseCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
