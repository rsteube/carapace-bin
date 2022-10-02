package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Start a debugger for the first matching coredump",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	rootCmd.AddCommand(debugCmd)
}
