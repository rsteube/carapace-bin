package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hasSessionCmd = &cobra.Command{
	Use:   "has-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hasSessionCmd).Standalone()

	hasSessionCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(hasSessionCmd)
}
