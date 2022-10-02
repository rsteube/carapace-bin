package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var mousedownCmd = &cobra.Command{
	Use:   "mousedown",
	Short: "Send a mousedown for the given button",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mousedownCmd).Standalone()

	mousedownCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before clicking")
	mousedownCmd.Flags().String("delay", "", "Specify how long, in milliseconds, to delay between clicks")
	mousedownCmd.Flags().String("repeat", "", "Specify how many times to click")
	mousedownCmd.Flags().String("window", "", "Specify a window to send a click to")
	rootCmd.AddCommand(mousedownCmd)

	carapace.Gen(mousedownCmd).FlagCompletion(carapace.ActionMap{
		"window": action.ActionWindows(),
	})

	carapace.Gen(mousedownCmd).PositionalCompletion(
		os.ActionMouseButtons(),
	)
}
