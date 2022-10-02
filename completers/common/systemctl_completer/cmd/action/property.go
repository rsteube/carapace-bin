package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionProperties(units ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"show"}
		args = append(args, units...)
		return carapace.ActionExecCommand("systemctl", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "=", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
