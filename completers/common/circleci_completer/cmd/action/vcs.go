package action

import "github.com/rsteube/carapace"

func ActionVcsTypes() carapace.Action {
	return carapace.ActionValues("github", "bitbucket")
}
