package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigKinds() carapace.Action {
	return carapace.ActionValues("service-defaults", "proxy-defaults")
}

func ActionConfigEntries(kind string) carapace.Action {
	return carapace.ActionExecCommand("consul", "config", "list", "--kind", kind)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
