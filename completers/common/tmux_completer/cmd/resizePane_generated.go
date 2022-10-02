package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resizePaneCmd = &cobra.Command{
	Use:   "resize-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizePaneCmd).Standalone()

	resizePaneCmd.Flags().BoolS("D", "D", false, "TODO description")
	resizePaneCmd.Flags().BoolS("L", "L", false, "TODO description")
	resizePaneCmd.Flags().BoolS("M", "M", false, "TODO description")
	resizePaneCmd.Flags().BoolS("R", "R", false, "TODO description")
	resizePaneCmd.Flags().BoolS("T", "T", false, "TODO description")
	resizePaneCmd.Flags().BoolS("U", "U", false, "TODO description")
	resizePaneCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	resizePaneCmd.Flags().StringS("t", "t", "", "target-pane")
	resizePaneCmd.Flags().StringS("x", "x", "", "width")
	resizePaneCmd.Flags().StringS("y", "y", "", "height")
	rootCmd.AddCommand(resizePaneCmd)
}
