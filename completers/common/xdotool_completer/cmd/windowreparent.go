package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var windowreparentCmd = &cobra.Command{
	Use:   "windowreparent",
	Short: "Reparent a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowreparentCmd).Standalone()

	rootCmd.AddCommand(windowreparentCmd)

	carapace.Gen(windowreparentCmd).PositionalCompletion(
		action.ActionWindows(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionWindows().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
