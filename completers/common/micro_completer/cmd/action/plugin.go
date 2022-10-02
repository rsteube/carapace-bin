package action

import (
	"strings"
	"time"

	"github.com/rsteube/carapace"
)

func ActionAvailablePlugins() carapace.Action {
	return carapace.ActionExecCommand("micro", "-plugin", "available")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[1 : len(lines)-1]...)
	}).Cache(5 * time.Minute)
}

func ActionInstalledPlugins() carapace.Action {
	return carapace.ActionExecCommand("micro", "-plugin", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[1 : len(lines)-1] {
			splitted := strings.SplitN(line, " ", 2)
			vals = append(vals, splitted[0], splitted[1][1:len(splitted[1])-1])
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
