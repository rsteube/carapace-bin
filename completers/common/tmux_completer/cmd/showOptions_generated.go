package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showOptionsCmd = &cobra.Command{
	Use:   "show-options",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showOptionsCmd).Standalone()

	showOptionsCmd.Flags().BoolS("A", "A", false, "TODO description")
	showOptionsCmd.Flags().BoolS("H", "H", false, "TODO description")
	showOptionsCmd.Flags().BoolS("g", "g", false, "TODO description")
	showOptionsCmd.Flags().BoolS("p", "p", false, "TODO description")
	showOptionsCmd.Flags().BoolS("q", "q", false, "TODO description")
	showOptionsCmd.Flags().BoolS("s", "s", false, "TODO description")
	showOptionsCmd.Flags().StringS("t", "t", "", "target-pane")
	showOptionsCmd.Flags().BoolS("v", "v", false, "TODO description")
	showOptionsCmd.Flags().BoolS("w", "w", false, "TODO description")
	rootCmd.AddCommand(showOptionsCmd)
}
