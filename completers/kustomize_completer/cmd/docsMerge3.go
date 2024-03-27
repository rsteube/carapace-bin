package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsMerge3Cmd = &cobra.Command{
	Use:   "docs-merge3",
	Short: "[Alpha] Documentation for merging Resources (3-way merge).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsMerge3Cmd).Standalone()
	rootCmd.AddCommand(docsMerge3Cmd)
}
