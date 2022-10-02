package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var builder_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove build cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(builder_pruneCmd).Standalone()

	builder_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused build cache, not just dangling ones")
	builder_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'until=24h')")
	builder_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	builder_pruneCmd.Flags().String("keep-storage", "", "Amount of disk space to keep for cache")
	builderCmd.AddCommand(builder_pruneCmd)
}
