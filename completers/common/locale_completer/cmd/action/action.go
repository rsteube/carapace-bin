package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionLocales() carapace.Action {
	return carapace.ActionExecCommand("locale", "-a")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[len(lines)-1], "\n")
	})
}

func ActionCharmaps() carapace.Action {
	return carapace.ActionExecCommand("locale", "-m")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[len(lines)-1], "\n")
	})
}
