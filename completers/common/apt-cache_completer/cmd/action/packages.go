package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("apt-cache", "pkgnames")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
