package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var nextWindowCmd = &cobra.Command{
	Use:   "next-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextWindowCmd).Standalone()

	nextWindowCmd.Flags().BoolS("a", "a", false, "TODO description")
	nextWindowCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(nextWindowCmd)
}
