package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionNamespaces() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "namespaces")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`^ - (?P<namespace>.*)$`)

		vals := make([]string, 0)
		lines := strings.Split(string(output), "\n")
		for _, line := range lines[1:] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
