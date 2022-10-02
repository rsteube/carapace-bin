package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionTemporaryFiles() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "vi -r 2>&1")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^On (?P<date>.*) saved \d+ lines of file "(?P<file>.*)"$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matched := r.FindStringSubmatch(line)
				vals = append(vals, matched[2], matched[1])
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
