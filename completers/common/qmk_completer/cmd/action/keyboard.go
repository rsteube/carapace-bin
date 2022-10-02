package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionKeyboards() carapace.Action {
	return carapace.ActionExecCommand("qmk", "list-keyboards")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
