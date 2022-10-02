package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getwindowclassnameCmd = &cobra.Command{
	Use:   "getwindowclassname",
	Short: "Prints the class name for the window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowclassnameCmd).Standalone()

	rootCmd.AddCommand(getwindowclassnameCmd)

	carapace.Gen(getwindowclassnameCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
