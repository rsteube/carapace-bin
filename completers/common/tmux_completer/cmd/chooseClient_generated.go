package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var chooseClientCmd = &cobra.Command{
	Use:   "choose-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseClientCmd).Standalone()

	chooseClientCmd.Flags().StringS("F", "F", "", "format")
	chooseClientCmd.Flags().StringS("K", "K", "", "key-format")
	chooseClientCmd.Flags().StringS("O", "O", "", "sort-order")
	chooseClientCmd.Flags().StringS("f", "f", "", "filter")
	chooseClientCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(chooseClientCmd)
}
