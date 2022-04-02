package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type pullrequest struct {
	Number  int
	Title   string
	IsDraft bool
	State   string
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
							s = style.BrightBlack
						} else {
							s = style.Green
						}
					case "CLOSED":
						s = style.Red
					case "MERGED":
						s = style.Magenta
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
