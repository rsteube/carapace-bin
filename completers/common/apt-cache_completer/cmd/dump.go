package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "shows a short listing of every package in the cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()

	rootCmd.AddCommand(dumpCmd)
}
