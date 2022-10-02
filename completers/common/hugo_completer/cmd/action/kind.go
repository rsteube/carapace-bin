package action

import "github.com/rsteube/carapace"

func ActionKinds() carapace.Action {
	return carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
		return carapace.ActionValues("page", "home", "section", "taxonomy", "term").Invoke(c).Filter(c.Parts).ToA()
	})
}
