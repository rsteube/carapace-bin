package action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionEnvironments(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var configDir string
		if flag := cmd.Flag("configDir"); flag != nil && flag.Changed {
			configDir = flag.Value.String()
		} else {
			path, err := configPath()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			configDir = fmt.Sprintf("%v/config", filepath.Dir(path))
		}
		files, err := os.ReadDir(configDir)
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
