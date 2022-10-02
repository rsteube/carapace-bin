package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getwindowpidCmd = &cobra.Command{
	Use:   "getwindowpid",
	Short: "Output the PID owning a given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowpidCmd).Standalone()

	rootCmd.AddCommand(getwindowpidCmd)

	carapace.Gen(getwindowpidCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
