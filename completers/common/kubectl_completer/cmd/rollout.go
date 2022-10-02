package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rolloutCmd = &cobra.Command{
	Use:   "rollout",
	Short: "Manage the rollout of a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolloutCmd).Standalone()

	rootCmd.AddCommand(rolloutCmd)
}
