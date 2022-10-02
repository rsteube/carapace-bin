package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a  GitLab Release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_deleteCmd).Standalone()
	release_deleteCmd.Flags().BoolP("with-tag", "t", false, "Delete associated tag")
	release_deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	releaseCmd.AddCommand(release_deleteCmd)

	carapace.Gen(release_deleteCmd).PositionalCompletion(
		action.ActionReleases(release_deleteCmd),
	)
}
