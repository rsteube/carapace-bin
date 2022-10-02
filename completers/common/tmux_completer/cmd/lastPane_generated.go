package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lastPaneCmd = &cobra.Command{
	Use:   "last-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lastPaneCmd).Standalone()

	lastPaneCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	lastPaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	lastPaneCmd.Flags().BoolS("e", "e", false, "TODO description")
	lastPaneCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(lastPaneCmd)
}
