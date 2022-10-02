package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var moveWindowCmd = &cobra.Command{
	Use:   "move-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveWindowCmd).Standalone()

	moveWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	moveWindowCmd.Flags().BoolS("b", "b", false, "TODO description")
	moveWindowCmd.Flags().BoolS("d", "d", false, "TODO description")
	moveWindowCmd.Flags().BoolS("k", "k", false, "TODO description")
	moveWindowCmd.Flags().BoolS("r", "r", false, "TODO description")
	moveWindowCmd.Flags().StringS("s", "s", "", "src-window")
	moveWindowCmd.Flags().StringS("t", "t", "", "dst-window")
	rootCmd.AddCommand(moveWindowCmd)
}
