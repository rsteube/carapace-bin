package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type project struct {
	Path        string
	Description string
}

func ActionGroupProjects(cmd *cobra.Command, groupID string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := fmt.Sprintf(`groups/%v/projects?order_by=updated_at&per_page=100&search=%v`, url.PathEscape(groupID), url.QueryEscape(c.CallbackValue))
		var queryResult []project
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, project := range queryResult {
				vals = append(vals, project.Path, project.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionUserProjects(cmd *cobra.Command, userID string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		query := fmt.Sprintf(`users/%v/projects?order_by=updated_at&per_page=100&search=%v`, url.PathEscape(userID), url.QueryEscape(c.CallbackValue))
		var queryResult []project
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, project := range queryResult {
				vals = append(vals, project.Path, project.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
