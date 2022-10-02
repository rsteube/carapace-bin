package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowstateCmd = &cobra.Command{
	Use:   "windowstate",
	Short: "Change a property on a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowstateCmd).Standalone()
	windowstateCmd.Flags().String("add", "", "add a property")
	windowstateCmd.Flags().String("remove", "", "remove a property")
	windowstateCmd.Flags().String("toggle", "", "toggle a property")

	rootCmd.AddCommand(windowstateCmd)

	carapace.Gen(windowstateCmd).FlagCompletion(carapace.ActionMap{
		"add":    action.ActionProperties(),
		"remove": action.ActionProperties(),
		"toggle": action.ActionProperties(),
	})

	carapace.Gen(windowstateCmd).PositionalCompletion(
		action.ActionWindows(),
	)
}
