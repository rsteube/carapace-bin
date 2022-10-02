package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var behaveCmd = &cobra.Command{
	Use:   "behave",
	Short: "Bind an action to an event on a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(behaveCmd).Standalone()

	rootCmd.AddCommand(behaveCmd)

	carapace.Gen(behaveCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionWindows(),
			carapace.ActionValues("%1", "%@"),
		).ToA(),
		carapace.ActionValuesDescribed(
			"mouse-enter", "Fires when the mouse enters a window",
			"mouse-leave", "Fires when the mouse leaves a window",
			"mouse-click", "Fires when the mouse is clicked",
			"focus", "Fires when the window gets input focus",
			"blur", "Fires when the window loses focus",
		),
	)
}
