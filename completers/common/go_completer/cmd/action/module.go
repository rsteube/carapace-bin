package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionModules(includeIndirect bool) carapace.Action {
	return carapace.ActionExecCommand("go", "list", "-m", "-e", "-f", "{{.Indirect}}\t{{.Path}}\t{{.Version}}", "all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile("^(?P<indirect>.*)\t(?P<path>.*)\t(?P<version>.*)$")

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if includeIndirect || matches[1] == "false" {
					vals = append(vals, matches[2], matches[3])
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
