package action

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionSearch(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) < 2 {
			return carapace.ActionMessage("search needs at least 2 characters")
		}

		args := append([]string{"search"}, defaultArgs(cmd)...)
		args = append(args, fmt.Sprintf("/^%v/", c.CallbackValue))
		return carapace.ActionExecCommand("brew", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines {
				if !strings.HasPrefix(line, "==>") {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
