package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

type release struct {
	Name       string
	Namespace  string
	Revision   string
	Updated    string
	Status     string
	Chart      string
	AppVersion string `json:"app_version"`
}

func ActionReleases() carapace.Action {
	return carapace.ActionExecCommand("helm", "list", "--output", "json")(func(output []byte) carapace.Action {
		var releases []release
		if err := json.Unmarshal(output, &releases); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0, len(releases)*2)
		for _, release := range releases {
			vals = append(vals, release.Name, release.Chart)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
