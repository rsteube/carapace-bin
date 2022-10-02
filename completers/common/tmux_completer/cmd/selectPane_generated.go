package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var selectPaneCmd = &cobra.Command{
	Use:   "select-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectPaneCmd).Standalone()

	selectPaneCmd.Flags().BoolS("D", "D", false, "TODO description")
	selectPaneCmd.Flags().BoolS("L", "L", false, "TODO description")
	selectPaneCmd.Flags().BoolS("M", "M", false, "TODO description")
	selectPaneCmd.Flags().BoolS("R", "R", false, "TODO description")
	selectPaneCmd.Flags().StringS("T", "T", "", "title")
	selectPaneCmd.Flags().BoolS("U", "U", false, "TODO description")
	selectPaneCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	selectPaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	selectPaneCmd.Flags().BoolS("e", "e", false, "TODO description")
	selectPaneCmd.Flags().BoolS("l", "l", false, "TODO description")
	selectPaneCmd.Flags().BoolS("m", "m", false, "TODO description")
	selectPaneCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(selectPaneCmd)
}
