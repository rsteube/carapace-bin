package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/localectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setX11KeymapCmd = &cobra.Command{
	Use:   "set-x11-keymap",
	Short: "Set X11 and console keyboard mappings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setX11KeymapCmd).Standalone()

	rootCmd.AddCommand(setX11KeymapCmd)

	carapace.Gen(setX11KeymapCmd).PositionalCompletion(
		action.ActionKeymapLayouts(),
		action.ActionKeymapModels(),
		action.ActionKeymapVariants(),
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionKeymapOptions().Invoke(c).Filter(c.Parts).ToMultiPartsA(":")
		}),
	)
}
