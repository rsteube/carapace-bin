package action

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type topic struct {
	Name             string
	ShortDescription string `json:"short_description"`
}

type topicQuery struct {
	Items []topic
}

func ActionTopicSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("topic search needs at least two characters")
		}

		var queryResult topicQuery
		return ApiV3Action(cmd, fmt.Sprintf(`search/topics?per_page=100&q=%v`, url.QueryEscape(c.CallbackValue)), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult.Items)*2)
			for _, topic := range queryResult.Items {
				vals = append(vals, topic.Name, topic.ShortDescription)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionTopics(cmd *cobra.Command, user string) carapace.Action {
	return carapace.ActionExecCommand("gh", "repo", "list", "--limit", "500", "--json", "repositoryTopics", "--jq", "[ .[].repositoryTopics | select(. != null) | .[].name ] | sort | unique | .[]", user)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

type repoTopicsQuery struct {
	Names []string
}

func ActionRepoTopics(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult repoTopicsQuery
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/topics`, repo.RepoOwner(), repo.RepoName()), &queryResult, func() carapace.Action {
			return carapace.ActionValues(queryResult.Names...)
		})
	})
}
