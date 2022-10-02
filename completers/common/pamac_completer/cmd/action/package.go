package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
)

func ActionPackageGroups() carapace.Action {
	return carapace.ActionExecCommand("pamac", "list", "--groups")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		for _, line := range lines[:len(lines)-1] {
			vals = append(vals, line, "[group]")
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionPackageSearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.CallbackValue) == 0 {
			return carapace.ActionValues()
		}
		return carapace.ActionExecCommand("pamac", "search", "--aur", c.CallbackValue)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ ]+) +(?P<installed>\[Installed\])? +(?P<version>[^ ]+) +(?P<repo>[^ ]+) $`)

			current := ""
			packages := make(map[string][]string)
			for _, line := range lines[:len(lines)-1] {
				if r.MatchString(line) {
					matches := r.FindStringSubmatch(line)
					current = matches[1]
					packages[current] = []string{fmt.Sprintf("[%v]", matches[4])}
				} else {
					packages[current] = append(packages[current], strings.TrimSpace(line))
				}
			}

			vals := make([]string, 0)
			for key, value := range packages {
				vals = append(vals, key, strings.Join(value, " "))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionInstalledPackages(explicit bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		arg := "-i"
		if explicit {
			arg = "-e"
		}
		return carapace.ActionExecCommand("pamac", "list", arg)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], fields[1])
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})

	})
}
