package action

import (
	"strings"

	"github.com/rsteube/carapace"
)

func ActionKeymaps() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-keymaps")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

func ActionKeymapModels() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-x11-keymap-models")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

func ActionKeymapLayouts() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-x11-keymap-layouts")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

func ActionKeymapVariants() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-x11-keymap-variants")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

func ActionKeymapOptions() carapace.Action {
	return carapace.ActionExecCommand("localectl", "list-x11-keymap-options")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}
