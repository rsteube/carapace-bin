package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionResources() carapace.Action {
	return carapace.ActionExecCommand("terraform", "state", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
