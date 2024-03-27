package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsMergeCmd = &cobra.Command{
	Use:   "docs-merge",
	Short: "[Alpha] Documentation for merging Resources (2-way merge).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsMergeCmd).Standalone()
	rootCmd.AddCommand(docsMergeCmd)
}
