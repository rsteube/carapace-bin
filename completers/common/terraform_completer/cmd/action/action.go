package action

import "github.com/rsteube/carapace"

func ActionBool() carapace.Action {
	return carapace.ActionValues("true", "false")
}
