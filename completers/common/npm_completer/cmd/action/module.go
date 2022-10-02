package action

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func nodeModulesPath(c carapace.Context) (string, error) {
	return util.FindReverse(c.Dir, "node_modules")
}

func ActionModules() carapace.Action {
	return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
		path, err := nodeModulesPath(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		fullSegments := make([]string, 0)
		for _, part := range c.Parts {
			if strings.HasPrefix(part, "@") {
				fullSegments = append(fullSegments, part)
			} else {
				fullSegments = append(fullSegments, part, "node_modules")
			}
		}

		contents, err := os.ReadDir(fmt.Sprintf("%v/%v", path, strings.Join(fullSegments, "/")))
		if err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for _, content := range contents {
			if content.IsDir() && !strings.HasPrefix(content.Name(), ".") {
				vals = append(vals, content.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
