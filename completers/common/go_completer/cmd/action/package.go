package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPackages() carapace.Action {
	return carapace.ActionExecCommand("go", "list", "-f", "{{.ImportPath}}\n{{.Doc}} ", "all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValuesDescribed(lines[:len(lines)-1]...)
	})
}
