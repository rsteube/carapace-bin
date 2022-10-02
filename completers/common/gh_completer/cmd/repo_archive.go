package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_archiveCmd).Standalone()
	repo_archiveCmd.Flags().BoolP("confirm", "y", false, "Skip the confirmation prompt")
	repoCmd.AddCommand(repo_archiveCmd)

	carapace.Gen(repo_archiveCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_archiveCmd),
	)
}
