package action

import (
	"github.com/rsteube/carapace"
)

func ActionDeployments() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		conf, err := loadConfig()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, deployment := range conf.Deployment.Targets {
			vals = append(vals, deployment.Name, deployment.Url)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
