package action

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
)

func ActionOrbs() carapace.Action {
	//  TODO somehow retrieven namespaces first for better performance
	return carapace.ActionExecCommand("circleci", "orb", "list", "--uncertified")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<version>.*)\)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1:]...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(24 * time.Hour)
}

func ActionOrbCategories() carapace.Action {
	return carapace.ActionExecCommand("circleci", "orb", "list-categories")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
