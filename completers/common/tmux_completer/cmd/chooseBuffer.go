package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var chooseBufferCmd = &cobra.Command{
	Use:   "choose-buffer",
	Short: "put a pane into buffer choice mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseBufferCmd).Standalone()

	chooseBufferCmd.Flags().StringS("F", "F", "", "specify format for each list item")
	chooseBufferCmd.Flags().StringS("K", "K", "", "specify format for each shortcut key")
	chooseBufferCmd.Flags().BoolS("N", "N", false, "start without the preview")
	chooseBufferCmd.Flags().StringS("O", "O", "", "initial sort order")
	chooseBufferCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	chooseBufferCmd.Flags().StringS("f", "f", "", "filter items")
	chooseBufferCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	chooseBufferCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(chooseBufferCmd)

	// TODO formats
	carapace.Gen(chooseBufferCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("time", "name", "size"),
		"t": tmux.ActionWindows(),
	})
}
