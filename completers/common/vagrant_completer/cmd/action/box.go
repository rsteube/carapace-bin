package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBoxes() carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionBoxProviders(name string, version string) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if name == matches[1] &&
					(version == "" || version == matches[3]) {
					vals = append(vals, matches[2])
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionBoxVersions(name string, provider string) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "box", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<provider>[^,]+), (?P<version>.*)\)$`)
		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if name == matches[1] &&
					(provider == "" || provider == matches[2]) {
					vals = append(vals, matches[3])
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
