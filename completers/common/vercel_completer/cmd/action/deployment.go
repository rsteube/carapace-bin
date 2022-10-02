package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionDeployments(cmd *cobra.Command) carapace.Action {
	return actionExecVercel(cmd, "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		if len(lines) > 1 {
			for _, line := range lines[1:] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[1], strings.Join(fields[2:], " "))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
