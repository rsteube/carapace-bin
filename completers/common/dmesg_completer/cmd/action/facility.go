package action

import "github.com/rsteube/carapace"

func ActionFacilities() carapace.Action {
	return carapace.ActionValuesDescribed(
		"kern", "kernel messages",
		"user", "random user-level messages",
		"mail", "mail system",
		"daemon", "system daemons",
		"auth", "security/authorization messages",
		"syslog", "messages generated internally by syslogd",
		"lpr", "line printer subsystem",
		"news", "network news subsystem",
	)
}
