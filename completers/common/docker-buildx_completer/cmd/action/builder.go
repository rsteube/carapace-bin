package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBuilders() carapace.Action {
	return carapace.ActionExecCommand("docker", "buildx", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) [*]? +(?P<driver>[^ ]+) *$`)

		vals := make([]string, 0)
		for _, line := range lines[1:] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
