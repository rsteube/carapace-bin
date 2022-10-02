package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type branch struct {
	Name   string
	Commit struct {
		Title string
	}
}

func ActionBranches(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var branches []branch
		query := fmt.Sprintf("/projects/:fullpath/repository/branches?search=^%v", url.QueryEscape(c.CallbackValue))
		return actionApi(cmd, query, &branches, func() carapace.Action {
			vals := make([]string, 0, len(branches)*2)
			for _, branch := range branches {
				vals = append(vals, branch.Name, branch.Commit.Title)
			}
			return carapace.ActionValuesDescribed(vals...).Style(style.Blue)
		})
	})
}
