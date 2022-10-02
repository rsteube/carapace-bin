package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var joinPaneCmd = &cobra.Command{
	Use:   "join-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinPaneCmd).Standalone()

	joinPaneCmd.Flags().BoolS("b", "b", false, "TODO description")
	joinPaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	joinPaneCmd.Flags().BoolS("f", "f", false, "TODO description")
	joinPaneCmd.Flags().BoolS("h", "h", false, "TODO description")
	joinPaneCmd.Flags().StringS("l", "l", "", "size")
	joinPaneCmd.Flags().StringS("s", "s", "", "src-pane")
	joinPaneCmd.Flags().StringS("t", "t", "", "dst-pane")
	joinPaneCmd.Flags().BoolS("v", "v", false, "TODO description")
	rootCmd.AddCommand(joinPaneCmd)
}
