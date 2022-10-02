package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
)

func ActionActivePackages() carapace.Action {
	return carapace.ActionExecCommand("flutter", "pub", "global", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, strings.SplitN(line, " ", 2)...)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionActivePackageExecutables(pkg string) carapace.Action {
	return carapace.ActionExecCommand("flutter", "pub", "global", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines[:len(lines)-1] {
			if splitted := strings.SplitN(line, " ", 2); splitted[0] == pkg {
				return pub.ActionHostedExecutables(pub.HostedExecutablesOpts{Name: pkg, Version: splitted[1]})
			}
		}
		return carapace.ActionValues()
	})
}
