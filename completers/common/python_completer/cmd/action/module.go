package action

import (
	"strings"
	"time"

	"github.com/rsteube/carapace"
)

func ActionModules() carapace.Action {
	return carapace.ActionExecCommand("python", "-c", "help('modules')")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[3 : len(lines)-4] {
			vals = append(vals, strings.Fields(line)...)
		}
		return carapace.ActionValues(vals...)
	}).Cache(24 * time.Hour)
}
