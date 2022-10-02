package action

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionContexts() carapace.Action {
	return carapace.ActionExecCommand("waypoint", "context", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if splitted := strings.Split(line, "|"); len(splitted) == 4 {
				name := strings.TrimSpace(splitted[1])
				if name != "NAME" {
					platform := strings.TrimSpace(splitted[2])
					address := strings.TrimSpace(splitted[3])
					vals = append(vals, name, fmt.Sprintf("%v %v", platform, address))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
