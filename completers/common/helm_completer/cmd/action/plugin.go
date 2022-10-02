package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("helm", "plugin", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		versionIndex := strings.Index(lines[0], "VERSION")
		descriptionIndex := strings.Index(lines[0], "DESCRIPTION")

		vals := make([]string, 0, len(lines)-2)
		for _, line := range lines[1 : len(lines)-1] {
			name := strings.TrimSpace(line[0:versionIndex])
			description := strings.TrimSpace(line[descriptionIndex:])
			vals = append(vals, name, description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
