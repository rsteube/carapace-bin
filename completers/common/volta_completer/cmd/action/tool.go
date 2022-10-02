package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionTools() carapace.Action {
	return carapace.ActionExecCommand("volta", "list", "--format", "plain")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if fields := strings.Fields(line); len(fields) > 1 {
				vals = append(vals, fields[1], fields[0])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
