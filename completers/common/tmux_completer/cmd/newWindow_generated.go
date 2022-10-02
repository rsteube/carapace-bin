package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var newWindowCmd = &cobra.Command{
	Use:   "new-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newWindowCmd).Standalone()

	newWindowCmd.Flags().StringS("F", "F", "", "format")
	newWindowCmd.Flags().BoolS("P", "P", false, "TODO description")
	newWindowCmd.Flags().BoolS("S", "S", false, "TODO description")
	newWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	newWindowCmd.Flags().BoolS("b", "b", false, "TODO description")
	newWindowCmd.Flags().StringS("c", "c", "", "start-directory")
	newWindowCmd.Flags().BoolS("d", "d", false, "TODO description")
	newWindowCmd.Flags().StringS("e", "e", "", "environment")
	newWindowCmd.Flags().BoolS("k", "k", false, "TODO description")
	newWindowCmd.Flags().StringS("n", "n", "", "window-name")
	newWindowCmd.Flags().StringS("t", "t", "", "target-window")
	rootCmd.AddCommand(newWindowCmd)
}
