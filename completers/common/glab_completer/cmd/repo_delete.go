package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing repository on GitLab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deleteCmd).Standalone()
	repo_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt and immediately delete the repository.")
	repoCmd.AddCommand(repo_deleteCmd)

	carapace.Gen(repo_deleteCmd).PositionalCompletion(
		action.ActionRepo(repo_deleteCmd),
	)
}
