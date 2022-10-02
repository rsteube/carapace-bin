package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEngines() carapace.Action {
	return carapace.ActionExecCommand("curl", "--engine", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			vals = append(vals, strings.TrimSpace(line))
		}
		return carapace.ActionValues(vals...)
	})
}
