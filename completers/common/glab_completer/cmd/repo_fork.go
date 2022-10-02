package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Create a fork of a GitLab repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_forkCmd).Standalone()
	repo_forkCmd.Flags().BoolP("clone", "c", false, "Clone the fork {true|false}")
	repo_forkCmd.Flags().StringP("name", "n", "", "The name assigned to the resultant project after forking")
	repo_forkCmd.Flags().StringP("path", "p", "", "The path assigned to the resultant project after forking")
	repo_forkCmd.Flags().Bool("remote", false, "Add remote for fork {true|false}")
	repoCmd.AddCommand(repo_forkCmd)

	// TODO path completion

	carapace.Gen(repo_forkCmd).PositionalCompletion(
		action.ActionRepo(repo_forkCmd),
	)
}
