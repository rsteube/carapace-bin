package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setDesktopCmd = &cobra.Command{
	Use:   "set_desktop",
	Short: "Change the current view to the specified desktop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setDesktopCmd).Standalone()

	setDesktopCmd.Flags().Bool("relative", false, "Use relative movements instead of absolute")
	rootCmd.AddCommand(setDesktopCmd)

	carapace.Gen(setDesktopCmd).PositionalCompletion(
		action.ActionDesktops(),
	)
}
