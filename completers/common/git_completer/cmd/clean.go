package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove untracked files from the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolS("X", "X", false, "remove only ignored files")
	cleanCmd.Flags().BoolS("d", "d", false, "remove whole directories")
	cleanCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	cleanCmd.Flags().StringP("exclude", "e", "", "add <pattern> to ignore rules")
	cleanCmd.Flags().BoolP("force", "f", false, "force")
	cleanCmd.Flags().BoolP("interactive", "i", false, "interactive cleaning")
	cleanCmd.Flags().BoolP("quiet", "q", false, "do not print names of files removed")
	cleanCmd.Flags().BoolS("x", "x", false, "remove ignored files, too")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
