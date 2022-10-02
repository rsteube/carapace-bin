package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionBootloaders() carapace.Action {
	return carapace.ActionExecCommand("sh", "-c", "qmk flash --bootloaders;true")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		r := regexp.MustCompile(`^\t(?P<bootloader>[^ ]+)$`)
		for _, line := range lines {
			if r.MatchString(line) {
				vals = append(vals, r.FindStringSubmatch(line)[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
}
