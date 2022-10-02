package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type discussionCategory struct {
	Name        string
	Description string
}

type discussionCategoriesQuery struct {
	Data struct {
		Repository struct {
			DiscussionCategories struct {
				Nodes []discussionCategory
			}
		}
	}
}

func ActionDiscussionCategories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult discussionCategoriesQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ discussionCategories( first: 100) { nodes { name, description } } }`, &queryResult, func() carapace.Action {
			categories := queryResult.Data.Repository.DiscussionCategories.Nodes
			vals := make([]string, len(categories)*2)
			for index, category := range categories {
				vals[index*2] = category.Name
				vals[index*2+1] = category.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
