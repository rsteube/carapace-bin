package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionUnits() carapace.Action {
	return carapace.ActionExecCommand("systemctl", "list-units")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if line == "" {
				break
			}
			fields := strings.Fields(line)
			vals = append(vals, fields[0], strings.Join(fields[4:], " "))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
