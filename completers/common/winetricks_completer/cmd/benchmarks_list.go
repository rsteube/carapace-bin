package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var benchmarks_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list verbs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(benchmarks_listCmd).Standalone()

	benchmarksCmd.AddCommand(benchmarks_listCmd)
}
