package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var displayMenuCmd = &cobra.Command{
	Use:   "display-menu",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(displayMenuCmd).Standalone()

	displayMenuCmd.Flags().BoolS("O", "O", false, "TODO description")
	displayMenuCmd.Flags().StringS("T", "T", "", "title")
	displayMenuCmd.Flags().StringS("c", "c", "", "target-client")
	displayMenuCmd.Flags().StringS("t", "t", "", "target-pane")
	displayMenuCmd.Flags().StringS("x", "x", "", "position")
	displayMenuCmd.Flags().StringS("y", "y", "", "position")
	rootCmd.AddCommand(displayMenuCmd)
}
