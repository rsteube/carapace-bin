package action

import (
	"os"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
)

func LoadAliases() (aliases map[string]string, err error) {
	var dir string
	if dir, err = os.UserConfigDir(); err == nil {
		var content []byte
		if content, err = os.ReadFile(dir + "/glab-cli/aliases.yml"); err == nil {
			err = yaml.Unmarshal(content, &aliases)
		}
	}
	return
}

func ActionAliases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		aliases, err := LoadAliases()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for alias, desc := range aliases {
			vals = append(vals, alias, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
