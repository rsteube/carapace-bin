package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEncoders() carapace.Action {
	return carapace.ActionExecCommand("ffmpeg", "-hide_banner", "-encoders")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^.{7} (?P<encoder>[^ ]+) +(?P<description>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines[4 : len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], matches[2])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
