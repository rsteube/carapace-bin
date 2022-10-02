package cmd

import (
	"github.com/spf13/cobra"
)

var merge_baseCmd = &cobra.Command{
	Use:   "merge-base",
	Short: "Find as good common ancestors as possible for a merge",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	merge_baseCmd.Flags().BoolP("all", "a", false, "output all common ancestors")
	merge_baseCmd.Flags().Bool("fork-point", false, "find where <commit> forked from reflog of <ref>")
	merge_baseCmd.Flags().Bool("independent", false, "list revs not reachable from others")
	merge_baseCmd.Flags().Bool("is-ancestor", false, "is the first one ancestor of the other?")
	merge_baseCmd.Flags().Bool("octopus", false, "find ancestors for a single n-way merge")
	rootCmd.AddCommand(merge_baseCmd)
}
