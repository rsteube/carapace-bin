package action

import (
	"github.com/rsteube/carapace"
)

func ActionWorkspaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		pj, err := loadPackageJson(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(pj.Workspaces...)
	})
}
