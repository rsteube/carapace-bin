package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/localectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setKeymapCmd = &cobra.Command{
	Use:   "set-keymap",
	Short: "Set console and X11 keyboard mappings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setKeymapCmd).Standalone()

	rootCmd.AddCommand(setKeymapCmd)

	carapace.Gen(setKeymapCmd).PositionalCompletion(
		action.ActionKeymaps(),
		action.ActionKeymaps(),
	)
}
