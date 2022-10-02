package action

import (
	"github.com/rsteube/carapace"
)

func ActionExecKinds() carapace.Action {
	return carapace.ActionValues(
		"auto.local",
		"auto.inline",
		"cli",
	)
}
