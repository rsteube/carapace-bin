package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEnvironments() carapace.Action {
	return carapace.ActionExecCommand("conda", "info", "--envs")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^# ]+)[ *]+(?P<path>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
