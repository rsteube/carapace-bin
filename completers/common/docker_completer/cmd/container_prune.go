package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var container_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove all stopped containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_pruneCmd).Standalone()

	container_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")
	container_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	containerCmd.AddCommand(container_pruneCmd)
}
