package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEmulators() carapace.Action {
	return carapace.ActionExecCommand("flutter", "--suppress-analytics", "emulators")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<id>[^ ]+) +• (?P<name>[^•]+) +• (?P<oem>[^ •]+) +• (?P<platform>[^ •]+)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], fmt.Sprintf("%v • %v • %v", matches[2], matches[3], matches[4]))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})

}
