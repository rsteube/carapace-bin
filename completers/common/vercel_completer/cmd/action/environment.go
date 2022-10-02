package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionEnvironmentVariables(cmd *cobra.Command) carapace.Action {
	return actionExecVercel(cmd, "env", "ls")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		if len(lines) > 2 {
			for _, line := range lines[2:] {
				if fields := strings.Fields(line); len(fields) > 2 {
					vals = append(vals, fields[0], strings.Join(fields[2:], " "))
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionEnvironments() carapace.Action {
	return carapace.ActionValues(
		"development",
		"preview",
		"production",
	)
}
