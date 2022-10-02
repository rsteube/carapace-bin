package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionRaftPeerIds(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO client flags from cmd
		return carapace.ActionExecCommand("consul", "operator", "raft", "list-peers")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1 : len(lines)-1] {
				fields := strings.Fields(line)
				vals = append(vals, fields[1], strings.Join(append(fields[:2], fields[2:]...), " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionRaftPeerAddresses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO client flags from cmd
		return carapace.ActionExecCommand("consul", "operator", "raft", "list-peers")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[1 : len(lines)-1] {
				fields := strings.Fields(line)
				vals = append(vals, fields[2], strings.Join(append(fields[:3], fields[3:]...), " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
