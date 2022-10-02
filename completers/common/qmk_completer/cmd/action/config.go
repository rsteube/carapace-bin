package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionConfigs() carapace.Action {
	return carapace.ActionExecCommand("qmk", "config", "--all")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			if splitted := strings.SplitN(line, "=", 2); len(splitted) == 2 {
				vals = append(vals, splitted...)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionConfigValues(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO more configs
		switch name {
		case "multibuild.keymap":
			return ActionKeymaps("")
		case "user.keyboard":
			return ActionKeyboards().Invoke(c).ToMultiPartsA("/")
		case "user.keymap":
			return ActionKeymaps("") // TODO user user.keyboard config
		case "setup.home":
			return carapace.ActionDirectories()
		default:
			return carapace.ActionValues()
		}
	})

}
