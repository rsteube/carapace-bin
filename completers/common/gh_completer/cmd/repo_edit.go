package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit repository settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_editCmd).Standalone()
	repo_editCmd.Flags().StringSlice("add-topic", []string{}, "Add repository topic")
	repo_editCmd.Flags().Bool("allow-forking", false, "Allow forking of an organization repository")
	repo_editCmd.Flags().String("default-branch", "", "Set the default branch `name` for the repository")
	repo_editCmd.Flags().Bool("delete-branch-on-merge", false, "Delete head branch when pull requests are merged")
	repo_editCmd.Flags().StringP("description", "d", "", "Description of the repository")
	repo_editCmd.Flags().Bool("enable-auto-merge", false, "Enable auto-merge functionality")
	repo_editCmd.Flags().Bool("enable-issues", false, "Enable issues in the repository")
	repo_editCmd.Flags().Bool("enable-merge-commit", false, "Enable merging pull requests via merge commit")
	repo_editCmd.Flags().Bool("enable-projects", false, "Enable projects in the repository")
	repo_editCmd.Flags().Bool("enable-rebase-merge", false, "Enable merging pull requests via rebase")
	repo_editCmd.Flags().Bool("enable-squash-merge", false, "Enable merging pull requests via squashed commit")
	repo_editCmd.Flags().Bool("enable-wiki", false, "Enable wiki in the repository")
	repo_editCmd.Flags().StringP("homepage", "h", "", "Repository home page `URL`")
	repo_editCmd.Flags().StringSlice("remove-topic", []string{}, "Remove repository topic")
	repo_editCmd.Flags().Bool("template", false, "Make the repository available as a template repository")
	repo_editCmd.Flags().String("visibility", "", "Change the visibility of the repository to {public,private,internal}")
	repoCmd.AddCommand(repo_editCmd)

	carapace.Gen(repo_editCmd).FlagCompletion(carapace.ActionMap{
		"add-topic": action.ActionTopicSearch(repo_editCmd),
		"default-branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				repo_editCmd.Flags().String("repo", c.Args[0], "")
				repo_editCmd.Flag("repo").Changed = true
			}
			return action.ActionBranches(repo_editCmd)
		}),
		"remove-topic": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				repo_editCmd.Flags().String("repo", c.Args[0], "")
				repo_editCmd.Flag("repo").Changed = true
			}
			return action.ActionRepoTopics(repo_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"visibility": carapace.ActionValues("public", "private", "internal"),
	})

	carapace.Gen(repo_editCmd).PositionalCompletion(
		action.ActionOwnerRepositories(repo_editCmd),
	)
}
