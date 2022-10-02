package action

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionInputFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.CallbackValue, "+-") {
			return carapace.ActionExecCommand("pandoc", "--list-input-formats")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}
		fields := extensionFields(c.CallbackValue)
		fields = fields[:len(fields)-1] // omit currently completed extension
		return ActionExtensions(fields[0]).Invoke(c).Filter(fields[1:]).Prefix(strings.Join(fields, "")).ToA()
	})
}

func ActionOutputFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if !strings.ContainsAny(c.CallbackValue, "+-") {
			return carapace.ActionExecCommand("pandoc", "--list-output-formats")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				return carapace.ActionValues(lines[:len(lines)-1]...)
			})
		}
		fields := extensionFields(c.CallbackValue)
		fields = fields[:len(fields)-1] // omit currently completed extension
		return ActionExtensions(fields[0]).Invoke(c).Filter(fields[1:]).Prefix(strings.Join(fields, "")).ToA()
	})
}

func extensionFields(s string) []string {
	fields := []string{s}
	for {
		if last := fields[len(fields)-1]; len(last) == 0 {
			break
		} else {
			if index := strings.IndexAny(last[1:], "+-"); index > -1 {
				fields = append(fields[:len(fields)-1], last[:index+1], last[index+1:])
			} else {
				break
			}
		}
	}
	return fields
}

func ActionFormats() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			ActionInputFormats(),
			ActionOutputFormats(),
		).Invoke(c).Merge().ToA()
	})
}

func ActionExtensions(format string) carapace.Action {
	return carapace.ActionExecCommand("pandoc", "--list-extensions", format)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

func ActionHighlightStyles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if util.HasPathPrefix(c.CallbackValue) {
			return carapace.ActionFiles(".theme")
		}
		return carapace.ActionValues("pygments", "kate", "monochrome", "breezeDark", "espresso", "zenburn", "haddock", "tango")
	})
}
