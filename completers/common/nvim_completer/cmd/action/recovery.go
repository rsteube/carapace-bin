package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionTemporaryFiles() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "nvim -r 2>&1")(func(output []byte) carapace.Action {
		rDate := regexp.MustCompile(`^ +owned by: .*dated: (?P<date>.*)$`)
		rFile := regexp.MustCompile(`^ +file name: (?P<file>.*)$`)

		lines := strings.Split(string(output), "\n")
		dates := make([]string, 0)
		files := make([]string, 0)
		for _, line := range lines {
			if rDate.MatchString(line) {
				dates = append(dates, rDate.FindStringSubmatch(line)[1])
			}
			if rFile.MatchString(line) {
				files = append(files, rFile.FindStringSubmatch(line)[1])
			}
		}

		vals := make([]string, 0)
		for index, file := range files {
			vals = append(vals, file, dates[index])
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
