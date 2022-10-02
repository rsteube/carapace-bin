package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove unused images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_pruneCmd).Standalone()

	image_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused images, not just dangling ones")
	image_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'until=<timestamp>')")
	image_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	imageCmd.AddCommand(image_pruneCmd)
}
