package action

import "github.com/rsteube/carapace"

func ActionLogLevels() carapace.Action {
	return carapace.ActionValues(
		"0", "1", "2", "3", "4", "5", "6", "7",
		"emerg", "alert", "crit", "err", "warning", "notice", "info", "debug",
	)
}
