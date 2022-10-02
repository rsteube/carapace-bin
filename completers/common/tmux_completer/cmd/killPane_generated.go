package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var killPaneCmd = &cobra.Command{
	Use:   "kill-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killPaneCmd).Standalone()

	killPaneCmd.Flags().BoolS("a", "a", false, "TODO description")
	killPaneCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(killPaneCmd)
}
