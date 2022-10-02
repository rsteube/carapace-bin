package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type user struct {
	Username string
	Name     string
}

type userQuery struct {
	Data struct {
		Users struct {
			Nodes []user
		}
	}
}

func ActionUsers(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 3 {
			return carapace.ActionMessage("search needs at least 3 characters")
		}

		query := fmt.Sprintf(`{ users(search: "%v") { nodes { username, name } } }`, c.CallbackValue)
		var queryResult userQuery
		return actionGraphql(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult.Data.Users.Nodes)*2)
			for _, user := range queryResult.Data.Users.Nodes {
				vals = append(vals, user.Username, user.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionProjectMembers(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 3 {
			return carapace.ActionMessage("search needs at least 3 characters")
		}
		query := fmt.Sprintf(`/projects/:fullpath/members/all?query=%v`, url.QueryEscape(c.CallbackValue))

		var queryResult []user
		return actionApi(cmd, query, &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, member := range queryResult {
				vals = append(vals, member.Username, member.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
