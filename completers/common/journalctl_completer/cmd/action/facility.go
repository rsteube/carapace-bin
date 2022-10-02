package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFacilities() carapace.Action {
	return carapace.ActionExecCommand("journalctl", "--facility=help")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1 : len(lines)-1]...)
	})
}
