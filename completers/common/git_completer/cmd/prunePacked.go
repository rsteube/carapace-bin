package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var prunePackedCmd = &cobra.Command{
	Use:   "prune-packed",
	Short: "Remove extra objects that are already in pack files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prunePackedCmd).Standalone()

	prunePackedCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	prunePackedCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	rootCmd.AddCommand(prunePackedCmd)
}
