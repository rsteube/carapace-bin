package action

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
)

func ActionVerbs() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("winetricks", "list-all")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^=][^ ]+) +(?P<description>.*)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Cache(24 * time.Hour).Invoke(c).ToMultiPartsA("=")
	})
}
