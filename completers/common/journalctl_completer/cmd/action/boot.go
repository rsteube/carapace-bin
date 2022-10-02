package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBoots() carapace.Action {
	return carapace.ActionExecCommand("journalctl", "--list-boots")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) > 2 {
				vals = append(vals, fields[0], strings.Join(fields[2:], " "))
				vals = append(vals, fields[1], strings.Join(fields[2:], " "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
