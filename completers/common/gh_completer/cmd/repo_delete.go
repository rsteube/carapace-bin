package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_deleteCmd).Standalone()
	repo_deleteCmd.Flags().Bool("confirm", false, "confirm deletion without prompting")
	repoCmd.AddCommand(repo_deleteCmd)

	carapace.Gen(repo_deleteCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_deleteCmd),
	)
}
