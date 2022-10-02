package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var keydownCmd = &cobra.Command{
	Use:   "keydown",
	Short: "Send keydown event for a given keystroke",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keydownCmd).Standalone()

	keydownCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before sending keystrokes")
	keydownCmd.Flags().String("delay", "", "Delay between keystrokes")
	keydownCmd.Flags().String("window", "", "Send keystrokes to a specific window id")
	rootCmd.AddCommand(keydownCmd)

	carapace.Gen(keydownCmd).FlagCompletion(carapace.ActionMap{
		"window": action.ActionWindows(),
	})

	carapace.Gen(keydownCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("+", func(c carapace.Context) carapace.Action {
			return action.ActionKeys().Invoke(c).Filter(c.Parts).ToA()
		}),
	)
}
