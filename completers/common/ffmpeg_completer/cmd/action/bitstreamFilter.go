package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBitstreamFilters() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-bsfs")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if line != "" {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}
