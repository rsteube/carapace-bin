package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volume_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove all unused local volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_pruneCmd).Standalone()

	volume_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'label=<label>')")
	volume_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	volumeCmd.AddCommand(volume_pruneCmd)
}
