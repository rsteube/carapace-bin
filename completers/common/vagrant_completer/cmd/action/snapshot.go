package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionSnapshots(machine string) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "snapshot", "list", machine)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1 : len(lines)-1]...)
	})
}
