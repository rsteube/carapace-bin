package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type coc struct {
	Key  string
	Name string
}

func ActionCocs(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []coc
		return ApiV3Action(cmd, `codes_of_conduct`, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, c := range queryResult {
				vals = append(vals, c.Key, c.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
