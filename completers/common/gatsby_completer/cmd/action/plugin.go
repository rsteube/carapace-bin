package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("gatsby", "plugin", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, "- ") {
				vals = append(vals, strings.TrimPrefix(line, "- "))
			}
		}
		return carapace.ActionValues(vals...)
	})
}
