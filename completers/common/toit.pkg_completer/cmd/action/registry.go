package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionRegistries() carapace.Action {
	return carapace.ActionExecCommand("toitpkg", "pkg", "registry", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^:]+): (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
