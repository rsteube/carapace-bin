package action

import (
	"time"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
)

func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("conda", "config", "--show")(func(output []byte) carapace.Action {
		var config map[string]interface{}
		if err := yaml.Unmarshal(output, &config); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for key := range config {
			vals = append(vals, key)
		}
		return carapace.ActionValues(vals...)
	}).Cache(24 * time.Hour)
}
