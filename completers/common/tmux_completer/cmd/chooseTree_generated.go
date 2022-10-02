package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var chooseTreeCmd = &cobra.Command{
	Use:   "choose-tree",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseTreeCmd).Standalone()

	chooseTreeCmd.Flags().StringS("F", "F", "", "format")
	chooseTreeCmd.Flags().StringS("K", "K", "", "key-format")
	chooseTreeCmd.Flags().StringS("O", "O", "", "sort-order")
	chooseTreeCmd.Flags().StringS("f", "f", "", "filter")
	chooseTreeCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(chooseTreeCmd)
}
