package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var displayPopupCmd = &cobra.Command{
	Use:   "display-popup",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayPopupCmd).Standalone()

	displayPopupCmd.Flags().BoolS("C", "C", false, "TODO description")
	displayPopupCmd.Flags().BoolS("E", "E", false, "TODO description")
	displayPopupCmd.Flags().StringS("c", "c", "", "target-client")
	displayPopupCmd.Flags().StringS("d", "d", "", "start-directory")
	displayPopupCmd.Flags().StringS("h", "h", "", "height")
	displayPopupCmd.Flags().StringS("t", "t", "", "target-pane")
	displayPopupCmd.Flags().StringS("w", "w", "", "width")
	displayPopupCmd.Flags().StringS("x", "x", "", "position")
	displayPopupCmd.Flags().StringS("y", "y", "", "position")
	rootCmd.AddCommand(displayPopupCmd)
}
