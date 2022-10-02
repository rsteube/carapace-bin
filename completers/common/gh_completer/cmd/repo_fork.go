package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var repo_forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "Create a fork of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_forkCmd).Standalone()
	repo_forkCmd.Flags().Bool("clone", false, "Clone the fork")
	repo_forkCmd.Flags().String("fork-name", "", "Rename the forked repository")
	repo_forkCmd.Flags().String("org", "", "Create the fork in an organization")
	repo_forkCmd.Flags().Bool("remote", false, "Add a git remote for the fork")
	repo_forkCmd.Flags().String("remote-name", "origin", "Specify the name for the new remote")
	repoCmd.AddCommand(repo_forkCmd)

	carapace.Gen(repo_forkCmd).FlagCompletion(carapace.ActionMap{
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})

	carapace.Gen(repo_forkCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_forkCmd),
	)

	carapace.Gen(repo_forkCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			repo := ""
			if args := repo_forkCmd.Flags().Args(); len(args) > 0 {
				repo = fmt.Sprintf("https://github.com/%v.git", args[0])
			}
			c.Args = append([]string{"clone", repo, ""}, c.Args...)
			return bridge.ActionCarapaceBin("git").Invoke(c).ToA()
		}),
	)
}
