package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type environment struct {
	Name        string
	ExternalUrl string `json:"external_url"`
}

func ActionEnvironments(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []environment
		return actionApi(cmd, "projects/:fullpath/environments", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, env := range queryResult {
				vals = append(vals, env.Name, env.ExternalUrl)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
