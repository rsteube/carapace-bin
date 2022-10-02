package action

import (
	"encoding/json"
	"strconv"

	"github.com/rsteube/carapace"
)

type revision struct {
	Revision    int
	Description string
}

func ActionRevisions(release string) carapace.Action {
	return carapace.ActionExecCommand("helm", "history", "--output", "json", release)(func(output []byte) carapace.Action {
		var revisions []revision
		if err := json.Unmarshal(output, &revisions); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, rev := range revisions {
			vals = append(vals, strconv.Itoa(rev.Revision), rev.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
