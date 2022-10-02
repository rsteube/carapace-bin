package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lockClientCmd = &cobra.Command{
	Use:   "lock-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockClientCmd).Standalone()

	lockClientCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(lockClientCmd)
}
