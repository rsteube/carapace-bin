package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var findWindowCmd = &cobra.Command{
	Use:   "find-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findWindowCmd).Standalone()

	findWindowCmd.Flags().BoolS("C", "C", false, "TODO description")
	findWindowCmd.Flags().BoolS("N", "N", false, "TODO description")
	findWindowCmd.Flags().BoolS("T", "T", false, "TODO description")
	findWindowCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	findWindowCmd.Flags().BoolS("i", "i", false, "TODO description")
	findWindowCmd.Flags().BoolS("r", "r", false, "TODO description")
	findWindowCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(findWindowCmd)
}
