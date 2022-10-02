package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// ActionSearchMultiRepo handles the stringslice `repo` flag of the search commands.
// For each repo given function is invoked with a fake repo flag and the return values are merged.
func ActionSearchMultiRepo(cmd *cobra.Command, f func(cmd *cobra.Command) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repos, err := cmd.Flags().GetStringSlice("repo")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if len(repos) == 0 {
			repos = append(repos, "")
		}

		batch := carapace.Batch()
		for _, repo := range repos {
			name := repo
			batch = append(batch, carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				dummyCmd := &cobra.Command{}
				dummyCmd.Flags().String("repo", name, "fake repo flag")
				if name != "" {
					dummyCmd.Flag("repo").Changed = true
				}
				return f(dummyCmd)
			}))
		}
		return batch.ToA()
	})
}

func ActionSearchIssueFields() carapace.Action {
	return carapace.ActionValues(
		"assignees",
		"author",
		"authorAssociation",
		"body",
		"closedAt",
		"commentsCount",
		"createdAt",
		"id",
		"isLocked",
		"isPullRequest",
		"labels",
		"number",
		"repository",
		"state",
		"title",
		"updatedAt",
		"url",
	)
}

func ActionSearchRepositoryFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"defaultBranch",
		"description",
		"forksCount",
		"fullName",
		"hasDownloads",
		"hasIssues",
		"hasPages",
		"hasProjects",
		"hasWiki",
		"homepage",
		"id",
		"isArchived",
		"isDisabled",
		"isFork",
		"isPrivate",
		"language",
		"license",
		"name",
		"openIssuesCount",
		"owner",
		"pushedAt",
		"size",
		"stargazersCount",
		"updatedAt",
		"visibility",
		"watchersCount",
	)
}

func ActionSearchRange(a carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := ""
		if strings.HasPrefix(c.CallbackValue, "<=") || strings.HasPrefix(c.CallbackValue, ">=") {
			prefix = c.CallbackValue[:2]
		} else if strings.HasPrefix(c.CallbackValue, "<") || strings.HasPrefix(c.CallbackValue, ">") {
			prefix = c.CallbackValue[:1]
		}

		return carapace.Batch(
			carapace.ActionValuesDescribed(
				">", "after",
				">=", "on or after",
				"<", "before",
				"<=", "on or before",
			),
			carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				if prefix != "" {
					c.CallbackValue = strings.TrimPrefix(c.CallbackValue, prefix)
					return a.Invoke(c).Prefix(prefix).ToA()

				}
				return carapace.ActionMultiParts("..", func(c carapace.Context) carapace.Action {
					if len(c.Parts) > 1 {
						return carapace.ActionValues()
					}
					return a
				})
			}),
		).ToA()
	})
}
