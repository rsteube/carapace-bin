package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionDevices() carapace.Action {
	return carapace.ActionExecCommand("avdmanager", "list", "device")(func(output []byte) carapace.Action {
		rName := regexp.MustCompile(`^id: (?P<id>\d+) or "(?P<name>.*)"$`)
		rDescription := regexp.MustCompile(`^    Name: (?P<description>.*)$`)

		lines := strings.Split(string(output), "\n")
		ids := make([]string, 0)
		names := make([]string, 0)
		for _, line := range lines {
			if rName.MatchString(line) {
				matches := rName.FindStringSubmatch(line)
				ids = append(ids, matches[1])
				names = append(names, matches[2])
			}
			if rDescription.MatchString(line) {
				matches := rDescription.FindStringSubmatch(line)
				ids = append(ids, matches[1])
				names = append(names, matches[1])
			}
		}
		return carapace.ActionValuesDescribed(append(ids, names...)...)
	})
}
