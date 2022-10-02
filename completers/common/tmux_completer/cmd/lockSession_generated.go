package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lockSessionCmd = &cobra.Command{
	Use:   "lock-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockSessionCmd).Standalone()

	lockSessionCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(lockSessionCmd)
}
