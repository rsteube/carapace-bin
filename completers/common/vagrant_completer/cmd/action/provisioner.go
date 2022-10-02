package action

import (
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionProvisioners() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "Vagrantfile")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		r := regexp.MustCompile(`^ *[^ ]+\.vm\.provision +"(?P<provisioner>[^"]+)".*`)
		vals := make([]string, 0)
		for _, line := range strings.Split(string(content), "\n") {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				vals = append(vals, matches[1])
			}
		}
		return carapace.ActionValues(vals...)
	})
	// TODO cache
}
