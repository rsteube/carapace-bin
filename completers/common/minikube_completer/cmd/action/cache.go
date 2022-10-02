package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionCachedImages() carapace.Action {
	return carapace.ActionExecCommand("minikube", "cache", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
