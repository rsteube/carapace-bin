package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionStrategies(host string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--strats"}
		if host != "" {
			args = append(args, "--host", host)
		}

		return carapace.ActionExecCommand("dict", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^ (?P<name>[^ ]+) +(?P<description>.*)$`)

			vals := make([]string, 0)
			for _, line := range lines {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					vals = append(vals, matches[1], matches[2])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
