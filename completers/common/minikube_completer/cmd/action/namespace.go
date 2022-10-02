package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionNamespaces() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("minikube", "kubectl", "get", "namespaces")(func(output []byte) carapace.Action {
			vals := make([]string, 0)

			lines := strings.Split(string(output), "\n")
			for _, line := range lines[1 : len(lines)-1] {
				vals = append(vals, strings.Fields(line)[0])
			}
			return carapace.ActionValues(vals...)
		})
	})
}
