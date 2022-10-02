package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resizeWindowCmd = &cobra.Command{
	Use:   "resize-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizeWindowCmd).Standalone()

	resizeWindowCmd.Flags().BoolS("A", "A", false, "TODO description")
	resizeWindowCmd.Flags().BoolS("D", "D", false, "TODO description")
	resizeWindowCmd.Flags().BoolS("L", "L", false, "TODO description")
	resizeWindowCmd.Flags().BoolS("R", "R", false, "TODO description")
	resizeWindowCmd.Flags().BoolS("U", "U", false, "TODO description")
	resizeWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	resizeWindowCmd.Flags().StringS("t", "t", "", "target-window")
	resizeWindowCmd.Flags().StringS("x", "x", "", "width")
	resizeWindowCmd.Flags().StringS("y", "y", "", "height")
	rootCmd.AddCommand(resizeWindowCmd)
}
