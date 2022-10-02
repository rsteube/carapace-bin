package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove object files and cached files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolS("cache", "cache", false, "remove the entire go build cache")
	cleanCmd.Flags().BoolS("i", "i", false, "remove the corresponding installed archive or binary")
	cleanCmd.Flags().BoolS("modcache", "modcache", false, "remove the entire module download cache")
	cleanCmd.Flags().BoolS("r", "r", false, "apply recursively to all the dependencies")
	cleanCmd.Flags().BoolS("testcache", "testcache", false, "expire all test results in the go build cache")
	addBuildFlags(cleanCmd)
	rootCmd.AddCommand(cleanCmd)
}
