package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionKeymaps(keyboard string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"list-keymaps"}
		if keyboard != "" {
			args = append(args, "--keyboard", keyboard)
		}
		return carapace.ActionExecCommand("qmk", args...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			return carapace.ActionValues(lines[:len(lines)-1]...)
		})
	})
}
