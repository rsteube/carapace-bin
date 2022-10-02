package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var box_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "remove outdated boxes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_pruneCmd).Standalone()

	box_pruneCmd.Flags().BoolP("dry-run", "n", false, "Only print the boxes that would be removed.")
	box_pruneCmd.Flags().BoolP("force", "f", false, "Destroy without confirmation even when box is in use.")
	box_pruneCmd.Flags().BoolP("keep-active-boxes", "k", false, "When combined with `--force`, will keep boxes still actively in use.")
	box_pruneCmd.Flags().String("name", "", "The specific box name to check for outdated versions.")
	box_pruneCmd.Flags().StringP("provider", "p", "", "The specific provider type for the boxes to destroy.")
	boxCmd.AddCommand(box_pruneCmd)

	carapace.Gen(box_pruneCmd).FlagCompletion(carapace.ActionMap{
		"name":     action.ActionBoxes(),
		"provider": action.ActionProviders(),
	})
}
