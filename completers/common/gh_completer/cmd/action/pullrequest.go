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

type pullrequest struct {
	Number    int
	Title     string
	IsDraft   bool
	State     string
	Labels    []label
	Assignees []struct {
		Login string
	}
	RequestedReviewers []struct {
		Login string
	} `json:"requested_reviewers"`
}

type pullRequestsQuery struct {
	Data struct {
		Repository struct {
			PullRequests struct {
				Nodes []pullrequest
			}
		}
	}
}

type PullRequestOpts struct {
	Closed    bool
	Merged    bool
	Open      bool
	DraftOnly bool
}

func (p *PullRequestOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{p.Closed, p.Merged, p.Open} {
		if include {
			states = append(states, []string{"CLOSED", "MERGED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

func ActionPullRequests(cmd *cobra.Command, opts PullRequestOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult pullRequestsQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo){ pullRequests(first: 100, states: %v, orderBy: {direction: DESC, field: UPDATED_AT}) { nodes { number, title, isDraft, state } } }`, opts.states()), &queryResult, func() carapace.Action {
			pullrequests := queryResult.Data.Repository.PullRequests.Nodes
			vals := make([]string, 0, len(pullrequests)*2)
			for _, pullrequest := range pullrequests {
				if !opts.DraftOnly || pullrequest.IsDraft {

					s := style.Default
					switch pullrequest.State {
					case "OPEN":
						if pullrequest.IsDraft {
							s = styles.Gh.Draft
						} else {
							s = styles.Gh.StateOpen
						}
					case "CLOSED":
						s = styles.Gh.StateClosed
					case "MERGED":
						s = styles.Gh.StateMerged
					default:
					}

					vals = append(vals, strconv.Itoa(pullrequest.Number), pullrequest.Title, s)
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func ActionPullRequestFields() carapace.Action {
	return carapace.ActionValues(
		"additions",
		"baseRefName",
		"changedFiles",
		"commits",
		"deletions",
		"files",
		"headRefName",
		"headRepository",
		"headRepositoryOwner",
		"isCrossRepository",
		"isDraft",
		"maintainerCanModify",
		"mergeable",
		"mergeCommit",
		"mergedAt",
		"mergedBy",
		"mergeStateStatus",
		"potentialMergeCommit",
		"reviewDecision",
		"reviewRequests",
		"reviews",
		"statusCheckRollup",
	)
}

func actionPullRequests(cmd *cobra.Command, id string, f func(p pullrequest) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult pullrequest
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/pulls/%v`, repo.RepoOwner(), repo.RepoName(), id), &queryResult, func() carapace.Action {
			return f(queryResult)
		})

	})
}

func ActionPullRequestLabels(cmd *cobra.Command, id string) carapace.Action {
	return actionPullRequests(cmd, id, func(p pullrequest) carapace.Action {
		vals := make([]string, 0)
		for _, label := range p.Labels {
			vals = append(vals, label.Name, label.Description, "#"+label.Color)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

func ActionPullRequestAssignees(cmd *cobra.Command, id string) carapace.Action {
	return actionPullRequests(cmd, id, func(p pullrequest) carapace.Action {
		vals := make([]string, 0)
		for _, assignee := range p.Assignees {
			vals = append(vals, assignee.Login)
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionPullRequestReviewers(cmd *cobra.Command, id string) carapace.Action {
	return actionPullRequests(cmd, id, func(p pullrequest) carapace.Action {
		vals := make([]string, 0)
		for _, reviewer := range p.RequestedReviewers {
			vals = append(vals, reviewer.Login)
		}
		return carapace.ActionValues(vals...)
	})
}

type commit struct {
	Sha    string
	Commit struct {
		Message string
	}
}

func ActionPullRequestCommits(cmd *cobra.Command, id string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult []commit
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/pulls/%v/commits`, repo.RepoOwner(), repo.RepoName(), id), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, c := range queryResult {
				vals = append(vals, c.Sha, c.Commit.Message)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
