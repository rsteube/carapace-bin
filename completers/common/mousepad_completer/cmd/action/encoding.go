package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionEncoding() carapace.Action {
	return carapace.ActionExecCommand("mousepad", "--list-encodings")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
