package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionWindows() carapace.Action {
	return carapace.ActionExecCommand("xdotool", "search", "")(func(output []byte) carapace.Action {
		ids := strings.Split(string(output), "\n")
		ids = ids[:len(ids)-1]

		args := make([]string, 0, 2*len(ids))
		for _, id := range ids {
			args = append(args, "getwindowname", id)
		}
		return carapace.ActionExecCommand("xdotool", args...)(func(output []byte) carapace.Action {
			names := strings.Split(string(output), "\n")
			names = names[:len(names)-1]

			if len(ids) != len(names) {
				return carapace.ActionMessage("windows ids and names differ in amount")
			}

			vals := make([]string, 0, len(ids)*2)
			for index, id := range ids {
				vals = append(vals, id, names[index])
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
