package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "removes old allowed files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	rootCmd.AddCommand(pruneCmd)
}
