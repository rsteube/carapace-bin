package action

import (
	"fmt"
	"os"

	"github.com/rsteube/carapace"
)

func ActionPrefixes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home, err := os.UserHomeDir()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		files, err := os.ReadDir(fmt.Sprintf("%v/.local/share/wineprefixes/", home))
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if file.IsDir() {
				vals = append(vals, file.Name())
			}
		}
		return carapace.ActionValues(vals...)
	})
}
