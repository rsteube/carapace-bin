package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionProjects() carapace.Action {
	return carapace.ActionExecCommand("waypoint", "project", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
