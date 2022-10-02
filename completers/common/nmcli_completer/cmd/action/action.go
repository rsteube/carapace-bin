package action

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionYesNo() carapace.Action {
	return carapace.ActionValues("yes", "no").StyleF(style.ForKeyword)
}
