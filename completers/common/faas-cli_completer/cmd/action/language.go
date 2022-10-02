package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionLanguageTemplatesNew() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "new", "--list")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`^- (?P<name>.*)$`)

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

func ActionLanguageTemplates() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "template", "store", "ls")(func(output []byte) carapace.Action {
		r := regexp.MustCompile(`^(?P<lang>[^ ]+) +(?P<source>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		lines := strings.Split(string(output), "\n")
		for _, line := range lines[2:] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[3])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
