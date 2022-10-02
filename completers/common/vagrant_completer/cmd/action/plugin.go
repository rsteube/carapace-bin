package action

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

type PluginOpts struct {
	Local  bool
	Global bool
}

func ActionPlugins(opts PluginOpts) carapace.Action {
	return carapace.ActionExecCommand("vagrant", "plugin", "list", "--local")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+) \((?P<version>[^,]+), (?P<location>.*)\)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				if (opts.Local && matches[3] == "local") ||
					(opts.Global && matches[3] == "global") {
					vals = append(vals, matches[1], matches[2])
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
