package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionProtocols() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-protocols")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			if strings.HasPrefix(line, "  ") {
				vals = append(vals, strings.TrimSpace(line))
			}
		}
		return carapace.ActionValues(vals...)
	})
}
