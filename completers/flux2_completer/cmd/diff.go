package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff a flux resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()
	rootCmd.AddCommand(diffCmd)
}
