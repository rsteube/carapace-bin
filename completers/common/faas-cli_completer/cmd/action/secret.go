package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionSecrets() carapace.Action {
	return carapace.ActionExecCommand("faas-cli", "secret", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[2 : len(lines)-2]...)
	})
}
