package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listPanesCmd = &cobra.Command{
	Use:   "list-panes",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listPanesCmd).Standalone()

	listPanesCmd.Flags().StringS("F", "F", "", "format")
	listPanesCmd.Flags().BoolS("a", "a", false, "TODO description")
	listPanesCmd.Flags().StringS("f", "f", "", "filter")
	listPanesCmd.Flags().BoolS("s", "s", false, "TODO description")
	listPanesCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(listPanesCmd)
}
