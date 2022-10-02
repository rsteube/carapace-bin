package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var displayPanesCmd = &cobra.Command{
	Use:   "display-panes",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayPanesCmd).Standalone()

	displayPanesCmd.Flags().BoolS("N", "N", false, "TODO description")
	displayPanesCmd.Flags().BoolS("b", "b", false, "TODO description")
	displayPanesCmd.Flags().StringS("d", "d", "", "duration")
	displayPanesCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(displayPanesCmd)
}
