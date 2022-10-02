package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionThemes() carapace.Action {
	return carapace.ActionExecCommand("vivid", "themes")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
