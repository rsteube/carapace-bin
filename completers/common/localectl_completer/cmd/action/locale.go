package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionLocales() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-locales")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
