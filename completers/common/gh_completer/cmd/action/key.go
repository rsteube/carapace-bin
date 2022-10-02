package action

import (
	"fmt"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type deployKey struct {
	Id    int
	Title string
	Key   string
}

func ActionDeployKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult []deployKey
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/keys`, repo.RepoOwner(), repo.RepoName()), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, key := range queryResult {
				vals = append(vals, strconv.Itoa(key.Id), fmt.Sprintf("%v %v", key.Title, truncateMiddle(26, key.Key)))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func truncateMiddle(maxWidth int, t string) string {
	if len(t) <= maxWidth {
		return t
	}

	ellipsis := "..."
	if maxWidth < len(ellipsis)+2 {
		return t[0:maxWidth]
	}

	halfWidth := (maxWidth - len(ellipsis)) / 2
	remainder := (maxWidth - len(ellipsis)) % 2
	return t[0:halfWidth+remainder] + ellipsis + t[len(t)-halfWidth:]
}
