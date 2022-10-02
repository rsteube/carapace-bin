package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var swapPaneCmd = &cobra.Command{
	Use:   "swap-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapPaneCmd).Standalone()

	swapPaneCmd.Flags().BoolS("D", "D", false, "TODO description")
	swapPaneCmd.Flags().BoolS("U", "U", false, "TODO description")
	swapPaneCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	swapPaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	swapPaneCmd.Flags().StringS("s", "s", "", "src-pane")
	swapPaneCmd.Flags().StringS("t", "t", "", "dst-pane")
	rootCmd.AddCommand(swapPaneCmd)
}
