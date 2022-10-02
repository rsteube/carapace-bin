package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEnvs() carapace.Action {
	return carapace.ActionExecCommand("qmk", "env")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if splitted := strings.SplitN(line, "=", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
