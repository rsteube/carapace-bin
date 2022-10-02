package action

import (
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type tag struct {
	Name   string
	Commit struct {
		Title string
	}
}

func ActionTags(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []tag
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/repository/tags?order_by=updated&per_page=100&search=%v", c.CallbackValue), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, tag := range queryResult {
				vals = append(vals, tag.Name, tag.Commit.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
