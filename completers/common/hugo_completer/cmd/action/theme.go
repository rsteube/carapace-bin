package action

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type theme struct {
	Name        string
	Description string
}

func loadTheme(path string) (t theme, err error) {
	var content []byte
	if content, err = os.ReadFile(path); err == nil {
		err = toml.Unmarshal(content, &t)
	}
	return
}

func ActionThemes(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var themesDir string
		if flag := cmd.Flag("themesDir"); flag != nil && flag.Changed {
			themesDir = flag.Value.String()
		} else {
			path, err := configPath()
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}
			themesDir = fmt.Sprintf("%v/themes", filepath.Dir(path))
		}
		files, err := os.ReadDir(themesDir)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, file := range files {
			if file.IsDir() {
				t, err := loadTheme(fmt.Sprintf("%v/%v/theme.toml", themesDir, file.Name()))
				if err != nil {
					return carapace.ActionMessage(err.Error())
				}
				vals = append(vals, file.Name(), t.Description)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
