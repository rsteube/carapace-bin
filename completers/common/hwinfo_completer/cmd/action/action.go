package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionUniqueIds() carapace.Action {
	return carapace.ActionExecCommand("hwinfo", "--all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		ruid := regexp.MustCompile(`^  Unique ID: (?P<uid>.*)$`)
		rmodel := regexp.MustCompile(`^  Model: (?P<uid>.*)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if ruid.MatchString(line) {
				if len(vals)%2 != 0 {
					vals = append(vals, "")
				}
				vals = append(vals, ruid.FindStringSubmatch(line)[1])
			}
			if rmodel.MatchString(line) {
				vals = append(vals, rmodel.FindStringSubmatch(line)[1])
			}
		}
		if len(vals)%2 != 0 {
			vals = append(vals, "")
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
