package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bugCmd = &cobra.Command{
	Use:   "bug",
	Short: "start a bug report",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bugCmd).Standalone()

	rootCmd.AddCommand(bugCmd)
}
