package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionHostnames() carapace.Action {
	return carapace.ActionExecCommand("waypoint", "hostname", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.Split(line, "|"); len(splitted) == 3 {
				name := strings.TrimSpace(splitted[0])
				if name != "HOSTNAME" {
					vals = append(vals, name, strings.TrimSpace(splitted[1]))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
