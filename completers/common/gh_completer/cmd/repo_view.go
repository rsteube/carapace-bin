package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_viewCmd).Standalone()
	repo_viewCmd.Flags().StringP("branch", "b", "", "View a specific branch of the repository")
	repo_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	repo_viewCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	repo_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	repo_viewCmd.Flags().BoolP("web", "w", false, "Open a repository in the browser")
	repoCmd.AddCommand(repo_viewCmd)

	carapace.Gen(repo_viewCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				repo_viewCmd.Flags().String("repo", c.Args[0], "fake repo flag for ActionBranches")
				return action.ActionBranches(repo_viewCmd)
			} else {
				return carapace.ActionValues()
			}
		}),
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionRepositoryFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(repo_viewCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_viewCmd),
	)
}
