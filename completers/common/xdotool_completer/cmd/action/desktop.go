package action

import (
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/number"
)

func ActionDesktops() carapace.Action {
	return carapace.ActionExecCommand("xdotool", "get_num_desktops")(func(output []byte) carapace.Action {
		num, err := strconv.Atoi(strings.Split(string(output), "\n")[0])
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return number.ActionRange(number.RangeOpts{Format: "%d", Start: 0, End: num - 1})
	})
}
