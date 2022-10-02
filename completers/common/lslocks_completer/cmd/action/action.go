package action

import "github.com/rsteube/carapace"

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"COMMAND", "command of the process holding the lock",
		"PID", "PID of the process holding the lock",
		"TYPE", "kind of lock",
		"SIZE", "size of the lock",
		"MODE", "lock access mode",
		"M", "mandatory state of the lock: 0 (none), 1 (set)",
		"START", "relative byte offset of the lock",
		"END", "ending offset of the lock",
		"PATH", "path of the locked file",
		"BLOCKER", "PID of the process blocking the lock",
	)
}
