package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionNodes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO client flags from cmd
		// TODO filter options (datacenter,...)
		return carapace.ActionExecCommand("consul", "catalog", "nodes")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1 : len(lines)-1] {
				fields := strings.Fields(line)
				vals = append(vals, fields[0], strings.Join(fields[1:], " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionNodeIdentity(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionNodes(cmd)
		case 1:
			return ActionDatacenters(cmd)
		default:
			return carapace.ActionValues()
		}
	})
}
