package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var distTagCmd = &cobra.Command{
	Use:   "dist-tag",
	Short: "Modify package distribution tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTagCmd).Standalone()
	addWorkspaceFlags(distTagCmd)

	rootCmd.AddCommand(distTagCmd)
}
