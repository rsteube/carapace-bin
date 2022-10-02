package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionAvds() carapace.Action {
	return carapace.ActionExecCommand("avdmanager", "list", "avd")(func(output []byte) carapace.Action {
		rName := regexp.MustCompile(`^ +Name: (?P<name>.*)$`)
		rPath := regexp.MustCompile(`^ +Path: (?P<path>.*)$`)

		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if rName.MatchString(line) {
				matches := rName.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
			if rPath.MatchString(line) {
				matches := rPath.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
