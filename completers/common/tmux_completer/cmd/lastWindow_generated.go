package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lastWindowCmd = &cobra.Command{
	Use:   "last-window",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lastWindowCmd).Standalone()

	lastWindowCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(lastWindowCmd)
}
