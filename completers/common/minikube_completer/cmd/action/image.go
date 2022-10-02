package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionImages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("minikube", "image", "ls")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...).Invoke(c).ToMultiPartsA("/").Invoke(c).ToMultiPartsA(":")
		})
	})
}
