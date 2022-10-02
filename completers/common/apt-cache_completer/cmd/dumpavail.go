package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dumpavailCmd = &cobra.Command{
	Use:   "dumpavail",
	Short: "prints out an available list to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpavailCmd).Standalone()

	rootCmd.AddCommand(dumpavailCmd)
}
