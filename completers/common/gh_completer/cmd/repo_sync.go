package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_syncCmd).Standalone()
	repo_syncCmd.Flags().StringP("branch", "b", "", "Branch to sync (default: main branch)")
	repo_syncCmd.Flags().Bool("force", false, "Hard reset the branch of the destination repository to match the source repository")
	repo_syncCmd.Flags().StringP("source", "s", "", "Source repository")
	repoCmd.AddCommand(repo_syncCmd)

	carapace.Gen(repo_syncCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if flag := repo_syncCmd.Flag("source"); flag.Changed {
				repo_syncCmd.Flags().String("repo", flag.Value.String(), "fake repo flag")
				repo_syncCmd.Flag("repo").Changed = true
			}
			return action.ActionBranches(repo_syncCmd)
		}),
		"source": action.ActionOwnerRepositories(repo_syncCmd),
	})

	carapace.Gen(repo_syncCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_syncCmd),
	)
}
