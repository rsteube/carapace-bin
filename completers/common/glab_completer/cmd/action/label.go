package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type label struct {
	Name        string
	Description string
	Color       string
}

func ActionLabels(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []label
		return actionApi(cmd, fmt.Sprintf("projects/:fullpath/labels?per_page=100&search=%v", url.QueryEscape(c.CallbackValue)), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, label := range queryResult {
				vals = append(vals, label.Name, label.Description, label.Color)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
			//}).Cache(1*time.Hour, cache.String(repo.FullName())) // TODO paginate instead of search and re-add caching
		})
	})
}
