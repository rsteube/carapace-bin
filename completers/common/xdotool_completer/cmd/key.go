package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/xdotool_completer/cmd/action"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Type a given keystroke",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyCmd).Standalone()

	keyCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before sending keystrokes")
	keyCmd.Flags().String("delay", "", "Delay between keystrokes")
	keyCmd.Flags().String("window", "", "Send keystrokes to a specific window id")
	rootCmd.AddCommand(keyCmd)

	carapace.Gen(keyCmd).FlagCompletion(carapace.ActionMap{
		"window": action.ActionWindows(),
	})

	carapace.Gen(keyCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("+", func(c carapace.Context) carapace.Action {
			return action.ActionKeys().Invoke(c).Filter(c.Parts).ToA()
		}),
	)
}
