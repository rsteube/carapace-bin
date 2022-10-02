package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionMachines() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch(
			ActionGlobalMachines(),
		)

		if _, err := util.FindReverse(c.Dir, "Vagrantfile"); err == nil {
			batch = append(batch, ActionLocalMachines())
		}

		return batch.Invoke(c).Merge().ToA()
	})
}

func ActionLocalMachines() carapace.Action {
	// TODO filter by status
	return carapace.ActionExecCommand("vagrant", "status")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<name>[^ ]+)  +(?P<status>[^(]+) \((?P<provider>.*)\)$`)

		vals := make([]string, 0)
		for _, line := range lines {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1], fmt.Sprintf("%v %v", matches[3], matches[2]))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionGlobalMachines() carapace.Action {
	// TODO filter by status
	return carapace.ActionExecCommand("vagrant", "global-status")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)
		for _, line := range lines[2:] {
			if strings.TrimSpace(line) == "" {
				break
			}
			fields := strings.Fields(line)
			vals = append(vals, fields[0], strings.Join(fields[1:], " "))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
