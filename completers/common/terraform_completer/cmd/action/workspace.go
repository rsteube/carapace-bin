package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionWorkspaces() carapace.Action {
	return carapace.ActionExecCommand("terraform", "workspace", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^[* ] (?P<name>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
