package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionChannels() carapace.Action {
	return carapace.ActionExecCommand("flutter", "channel", "--suppress-analytics")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			vals = append(vals, line[2:])
		}
		return carapace.ActionValues(vals...)
	})
}
