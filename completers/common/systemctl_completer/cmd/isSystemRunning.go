package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var isSystemRunningCmd = &cobra.Command{
	Use:   "is-system-running",
	Short: "Check whether system is fully running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isSystemRunningCmd).Standalone()

	rootCmd.AddCommand(isSystemRunningCmd)
}
