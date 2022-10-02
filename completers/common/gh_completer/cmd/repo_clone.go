package cmd

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var repo_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_cloneCmd).Standalone()
	repo_cloneCmd.Flags().StringP("upstream-remote-name", "u", "upstream", "Upstream remote name when cloning a fork")
	repoCmd.AddCommand(repo_cloneCmd)

	carapace.Gen(repo_cloneCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_cloneCmd),
		carapace.ActionDirectories(),
	)

	carapace.Gen(repo_cloneCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			repo := ""
			if args := repo_cloneCmd.Flags().Args(); len(args) > 0 {
				repo = fmt.Sprintf("https://github.com/%v.git", args[0])
			}
			c.Args = append([]string{"clone", repo, ""}, c.Args...)
			return bridge.ActionCarapaceBin("git").Invoke(c).ToA()
		}),
	)
}
