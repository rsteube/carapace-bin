package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/color"
)

func ActionColors() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			color.ActionHexColors(),
			color.ActionXtermColorNames(),
		).Invoke(c).Merge().ToA()
	})
}
