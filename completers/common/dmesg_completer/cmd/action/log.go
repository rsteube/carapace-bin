package action

import "github.com/rsteube/carapace"

func ActionLogLevels() carapace.Action {
	return carapace.ActionValuesDescribed(
		"emerg", "system is unusable",
		"alert", "action must be taken immediately",
		"crit", "critical conditions",
		"err", "error conditions",
		"warn", "warning conditions",
		"notice", "normal but significant condition",
		"info", "informational",
		"debug", "debug-level messages",
	)
}
