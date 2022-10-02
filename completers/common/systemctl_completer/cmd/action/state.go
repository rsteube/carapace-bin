package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionStates() carapace.Action {
	return carapace.ActionExecCommand("systemctl", "--state=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if !strings.HasPrefix(line, "Available ") && line != "" {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
