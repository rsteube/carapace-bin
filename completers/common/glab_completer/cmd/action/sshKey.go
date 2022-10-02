package action

import (
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type sshKey struct {
	Id    int
	Title string
}

func ActionSshKeyIds(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []sshKey
		return actionApi(cmd, "/user/keys", &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult)*2)
			for _, sshKey := range queryResult {
				vals = append(vals, strconv.Itoa(sshKey.Id), sshKey.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
