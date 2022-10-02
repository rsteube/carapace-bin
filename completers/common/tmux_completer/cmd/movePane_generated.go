package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var movePaneCmd = &cobra.Command{
	Use:   "move-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(movePaneCmd).Standalone()

	movePaneCmd.Flags().BoolS("b", "b", false, "TODO description")
	movePaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	movePaneCmd.Flags().BoolS("f", "f", false, "TODO description")
	movePaneCmd.Flags().BoolS("h", "h", false, "TODO description")
	movePaneCmd.Flags().StringS("l", "l", "", "size")
	movePaneCmd.Flags().StringS("s", "s", "", "src-pane")
	movePaneCmd.Flags().StringS("t", "t", "", "dst-pane")
	movePaneCmd.Flags().BoolS("v", "v", false, "TODO description")
	rootCmd.AddCommand(movePaneCmd)
}
