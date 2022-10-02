package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

func ActionContents(cmd *cobra.Command, path string, branch string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return gh.ActionContents(gh.ContentOpts{Host: repo.RepoHost(), Owner: repo.RepoOwner(), Name: repo.RepoName(), Branch: branch, Path: path})
	})
}
