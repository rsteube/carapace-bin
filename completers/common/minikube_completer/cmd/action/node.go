package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionNodes() carapace.Action {
	return carapace.ActionExecCommand("minikube", "node", "list")(func(output []byte) carapace.Action {
		vals := make([]string, 0)

		lines := strings.Split(string(output), "\n")
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			vals = append(vals, fields[0], fields[1])
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
