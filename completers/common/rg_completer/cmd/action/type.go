package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionTypes() carapace.Action {
	return carapace.ActionExecCommand("rg", "--type-list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, ": ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
