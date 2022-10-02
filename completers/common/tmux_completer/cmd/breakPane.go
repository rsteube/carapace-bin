package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var breakPaneCmd = &cobra.Command{
	Use:   "break-pane",
	Short: "break a pane from an existing into a new window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(breakPaneCmd).Standalone()

	breakPaneCmd.Flags().StringS("F", "F", "", "specify output format")
	breakPaneCmd.Flags().BoolS("P", "P", false, "print information of new window after it has been created")
	breakPaneCmd.Flags().BoolS("a", "a", false, "move window to next index after")
	breakPaneCmd.Flags().BoolS("b", "b", false, "move window to next index before")
	breakPaneCmd.Flags().BoolS("d", "d", false, "don't make the new window become the active one")
	breakPaneCmd.Flags().StringS("n", "n", "", "specify window name")
	breakPaneCmd.Flags().StringS("s", "s", "", "specify source pane")
	breakPaneCmd.Flags().StringS("t", "t", "", "specify destination window")
	rootCmd.AddCommand(breakPaneCmd)

	// TODO output format
	carapace.Gen(breakPaneCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionPanes(),
		"t": tmux.ActionWindows(),
	})
}
