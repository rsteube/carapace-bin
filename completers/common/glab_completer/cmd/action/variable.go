package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type variable struct {
	Key   string
	Value string
}

func ActionVariables(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []variable
		return actionApi(cmd, "projects/:fullpath/variables", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, v := range queryResult {
				vals = append(vals, v.Key, v.Value)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
