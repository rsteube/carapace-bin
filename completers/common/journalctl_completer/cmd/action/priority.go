package action

import "github.com/rsteube/carapace"

func ActionPriorities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"0", "emerg",
		"1", "alert",
		"2", "crit",
		"3", "err",
		"4", "warning",
		"5", "notice",
		"6", "info",
		"7", "debug",
		"emerg", "0",
		"alert", "1",
		"crit", "2",
		"err", "3",
		"warning", "4",
		"notice", "5",
		"info", "6",
		"debug", "7",
	)
}
