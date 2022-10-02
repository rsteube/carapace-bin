package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a project/repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_viewCmd).Standalone()
	repo_viewCmd.Flags().StringP("branch", "b", "", "View a specific branch of the repository")
	repo_viewCmd.Flags().BoolP("web", "w", false, "Open a project in the browser")

	repoCmd.AddCommand(repo_viewCmd)

	carapace.Gen(repo_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				action.FakeRepoFlag(repo_viewCmd, c.Args[0])
				return action.ActionBranches(repo_viewCmd)
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(repo_viewCmd).PositionalCompletion(
		action.ActionRepo(repo_viewCmd),
	)
}
