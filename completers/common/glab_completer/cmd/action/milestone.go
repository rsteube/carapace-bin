package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type milestone struct {
	Title       string
	Description string
}

func ActionMilestones(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []milestone
		return actionApi(cmd, "projects/:fullpath/milestones", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, milestone := range queryResult {
				vals = append(vals, milestone.Title, milestone.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
