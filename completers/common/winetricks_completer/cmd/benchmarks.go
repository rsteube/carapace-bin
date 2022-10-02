package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var benchmarksCmd = &cobra.Command{
	Use:   "benchmarks",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchmarksCmd).Standalone()

	rootCmd.AddCommand(benchmarksCmd)
}
