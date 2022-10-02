package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionFields() carapace.Action {
	return carapace.ActionExecCommand("journalctl", "--fields")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
