package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move or rename a file, a directory, or a symlink",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mvCmd).Standalone()

	mvCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	mvCmd.Flags().BoolP("force", "f", false, "force move/rename even if target exists")
	mvCmd.Flags().BoolS("k", "k", false, "skip move/rename errors")
	mvCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(mvCmd)

	carapace.Gen(mvCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
