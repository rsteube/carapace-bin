package action

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type mergeRequest struct {
	Iid   int
	Title string
	State string
}

func ActionMergeRequests(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		stateQuery := ""
		if state != "" {
			stateQuery = "&state=" + url.QueryEscape(state)
		}

		var queryResult []mergeRequest
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/merge_requests?order_by=updated_at&per_page=100%v", stateQuery), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, mr := range queryResult {
				var s string
				switch mr.State {
				case "opened":
					s = style.Green
				case "closed":
					s = style.Red
				case "locked":
					s = style.Cyan
				case "merged":
					s = style.Magenta
				default:
					s = style.Default
				}

				vals = append(vals, strconv.Itoa(mr.Iid), mr.Title, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}

func ActionMergeRequestsAndBranches(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			ActionBranches(cmd),
			ActionMergeRequests(cmd, state),
		).Invoke(c).Merge().ToA()
	})
}
