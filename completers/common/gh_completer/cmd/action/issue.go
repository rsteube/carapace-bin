package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/styles"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type issue struct {
	Number    int
	Title     string
	State     string
	Labels    []label
	Assignees []struct {
		Login string
	}
}

type issueQuery struct {
	Data struct {
		Repository struct {
			Issues struct {
				Nodes []issue
			}
		}
	}
}

type IssueOpts struct {
	Closed bool
	Open   bool
}

func (i *IssueOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{i.Closed, i.Open} {
		if include {
			states = append(states, []string{"CLOSED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

func ActionIssues(cmd *cobra.Command, opts IssueOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult issueQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo){ issues(first: 100, states: %v, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, state } } }`, opts.states()), &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.Issues.Nodes
			vals := make([]string, 0)
			for _, issue := range issues {
				s := style.Default
				switch issue.State {
				case "OPEN":
					s = styles.Gh.StateOpen
				case "CLOSED":
					s = styles.Gh.StateClosed
				case "MERGED":
					s = styles.Gh.StateMerged
				default:
				}

				vals = append(vals, strconv.Itoa(issue.Number), issue.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

type pinnedIssueQuery struct {
	Data struct {
		Repository struct {
			PinnedIssues struct {
				Nodes []struct {
					Issue issue
				}
			}
		}
	}
}

func ActionPinnedIssues(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult pinnedIssueQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ pinnedIssues(first: 3) { nodes { issue { number, title, state } } } }`, &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.PinnedIssues.Nodes
			vals := make([]string, 0)
			for _, issue := range issues {
				s := style.Default
				switch issue.Issue.State {
				case "OPEN":
					s = styles.Gh.StateOpen
				case "CLOSED":
					s = styles.Gh.StateClosed
				case "MERGED":
					s = styles.Gh.StateMerged
				default:
				}

				vals = append(vals, strconv.Itoa(issue.Issue.Number), issue.Issue.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func ActionIssueFields() carapace.Action {
	return carapace.ActionValues(
		"assignees",
		"author",
		"body",
		"closed",
		"comments",
		"createdAt",
		"closedAt",
		"id",
		"labels",
		"milestone",
		"number",
		"projectCards",
		"reactionGroups",
		"state",
		"title",
		"updatedAt",
		"url",
	)
}

func actionIssue(cmd *cobra.Command, issueId string, f func(i issue) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult issue
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/issues/%v`, repo.RepoOwner(), repo.RepoName(), issueId), &queryResult, func() carapace.Action {
			return f(queryResult)
		})

	})
}

func ActionIssueLabels(cmd *cobra.Command, issueId string) carapace.Action {
	return actionIssue(cmd, issueId, func(i issue) carapace.Action {
		vals := make([]string, 0)
		for _, label := range i.Labels {
			vals = append(vals, label.Name, label.Description, "#"+label.Color)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

func ActionIssueAssignees(cmd *cobra.Command, issueId string) carapace.Action {
	return actionIssue(cmd, issueId, func(i issue) carapace.Action {
		vals := make([]string, 0)
		for _, assignee := range i.Assignees {
			vals = append(vals, assignee.Login)
		}
		return carapace.ActionValues(vals...)
	})
}
