package action

import (
	"encoding/json"
	"fmt"

	"github.com/rsteube/carapace"
)

type addon struct {
	Profile string
	Status  string
}

func ActionAddons() carapace.Action {
	return carapace.ActionExecCommand("minikube", "addons", "list", "--output", "json")(func(output []byte) carapace.Action {
		var addons map[string]addon
		if err := json.Unmarshal(output, &addons); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(addons)*2)
		for name, addon := range addons {
			vals = append(vals, name, fmt.Sprintf("%v - %v", addon.Profile, addon.Status))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
