package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Print first matching coredump to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()

	rootCmd.AddCommand(dumpCmd)
}
