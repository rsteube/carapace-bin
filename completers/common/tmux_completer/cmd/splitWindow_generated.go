package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var splitWindowCmd = &cobra.Command{
	Use:   "split-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitWindowCmd).Standalone()

	splitWindowCmd.Flags().StringS("F", "F", "", "format")
	splitWindowCmd.Flags().BoolS("I", "I", false, "TODO description")
	splitWindowCmd.Flags().BoolS("P", "P", false, "TODO description")
	splitWindowCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	splitWindowCmd.Flags().BoolS("b", "b", false, "TODO description")
	splitWindowCmd.Flags().StringS("c", "c", "", "start-directory")
	splitWindowCmd.Flags().BoolS("d", "d", false, "TODO description")
	splitWindowCmd.Flags().StringS("e", "e", "", "environment")
	splitWindowCmd.Flags().BoolS("h", "h", false, "TODO description")
	splitWindowCmd.Flags().StringS("l", "l", "", "size")
	splitWindowCmd.Flags().StringS("t", "t", "", "target-pane")
	splitWindowCmd.Flags().BoolS("v", "v", false, "TODO description")
	rootCmd.AddCommand(splitWindowCmd)
}
