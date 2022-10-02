package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionTeams(cmd *cobra.Command) carapace.Action {
	return actionExecVercel(cmd, "teams", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		if len(lines) > 2 {
			for _, line := range lines[2:] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], strings.Join(fields[1:], " "))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
