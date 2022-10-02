package action

import "github.com/rsteube/carapace"

func ActionColumns() carapace.Action {
	return carapace.ActionValuesDescribed(
		"USER", "user name",
		"UID", "user ID",
		"GECOS", "full user name",
		"FAILED-TTY", "where did the login fail?",
		"HUSHED", "user's hush settings",
		"PWD-WARN", "days user is warned of password expiration",
		"PWD-CHANGE", "date of last password change",
		"PWD-MIN", "number of days required between changes",
		"PWD-MAX", "max number of days a password may remain unchanged",
		"PWD-EXPIR", "password expiration date",
		"CONTEXT", "the user's security context",
		"PROC", "number of processes run by the user",
	)
}
