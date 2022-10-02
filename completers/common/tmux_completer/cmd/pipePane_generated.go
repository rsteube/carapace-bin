package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pipePaneCmd = &cobra.Command{
	Use:   "pipe-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipePaneCmd).Standalone()

	pipePaneCmd.Flags().BoolS("I", "I", false, "TODO description")
	pipePaneCmd.Flags().BoolS("O", "O", false, "TODO description")
	pipePaneCmd.Flags().BoolS("o", "o", false, "TODO description")
	pipePaneCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(pipePaneCmd)
}
