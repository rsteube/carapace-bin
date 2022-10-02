package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFunctions() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "list")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`^(?P<name>[^ ]+).*$`)

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

func ActionFunctionsStore() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "store", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)
		lines := strings.Split(string(output), "\n")
		index := strings.Index(lines[1], "DESCRIPTION")
		for _, line := range lines[2:] {
			if len(line) > index {
				vals = append(
					vals,
					strings.TrimSpace(string([]rune(line)[:index])),
					strings.TrimSpace(string([]rune(line)[index:])),
				)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
