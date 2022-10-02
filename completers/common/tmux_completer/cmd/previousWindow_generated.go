package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var previousWindowCmd = &cobra.Command{
	Use:   "previous-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previousWindowCmd).Standalone()

	previousWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	previousWindowCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(previousWindowCmd)
}
