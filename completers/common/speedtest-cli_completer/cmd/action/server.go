package action

import (
	"regexp"
	"strings"
	"time"

	"github.com/rsteube/carapace"
)

func ActionServers() carapace.Action {
	return carapace.ActionExecCommand("speedtest-cli", "--list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^ *(?P<id>\d+)\) (?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(5 * time.Minute)
}
