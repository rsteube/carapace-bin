package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove unneeded data from the repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()
	pruneCmd.Flags().BoolP("dry-run", "n", false, "do not modify the repository, just print what would be done")
	pruneCmd.Flags().String("max-repack-size", "", "maximum `size` to repack (allowed suffixes: k/K, m/M, g/G, t/T)")
	pruneCmd.Flags().String("max-unused", "5%", "tolerate given `limit` of unused data (absolute value in bytes with suffixes k/K, m/M, g/G, t/T, a value in % or the word 'unlimited')")
	pruneCmd.Flags().Bool("repack-cacheable-only", false, "only repack packs which are cacheable")
	rootCmd.AddCommand(pruneCmd)
}
