package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var mouseupCmd = &cobra.Command{
	Use:   "mouseup",
	Short: "Send a mouseup for the given button",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mouseupCmd).Standalone()

	mouseupCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before clicking")
	mouseupCmd.Flags().String("delay", "", "Specify how long, in milliseconds, to delay between clicks")
	mouseupCmd.Flags().String("repeat", "", "Specify how many times to click")
	mouseupCmd.Flags().String("window", "", "Specify a window to send a click to")
	rootCmd.AddCommand(mouseupCmd)

	carapace.Gen(mouseupCmd).FlagCompletion(carapace.ActionMap{
		"window": action.ActionWindows(),
	})

	carapace.Gen(mouseupCmd).PositionalCompletion(
		os.ActionMouseButtons(),
	)
}
