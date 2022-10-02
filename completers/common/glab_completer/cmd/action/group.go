package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type group struct {
	Path        string
	Description string
}

func ActionGroups(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := fmt.Sprintf(`groups?all_available=true&search=%v`, url.QueryEscape(c.CallbackValue))

		var queryResult []group
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, group := range queryResult {
				vals = append(vals, group.Path, group.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionSubgroups(cmd *cobra.Command, groupID string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := fmt.Sprintf(`groups/%v/subgroups?all_available=true&search=%v`, url.PathEscape(groupID), url.QueryEscape(c.CallbackValue))

		var queryResult []group
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			// TODO allow failure
			//			if strings.Contains(err.Error(), "404 Group Not Found") {
			//				return carapace.ActionValues() // fail silently for repo completion
			//			}
			vals := make([]string, 0, len(queryResult)*2)
			for _, group := range queryResult {
				vals = append(vals, group.Path, group.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
