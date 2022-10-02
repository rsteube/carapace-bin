package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var clearHistoryCmd = &cobra.Command{
	Use:   "clear-history",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearHistoryCmd).Standalone()
	clearHistoryCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(clearHistoryCmd)
}
