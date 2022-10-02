package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
)

type filesystem struct {
	Filesystem string
	MountPoint string `json:"mount_point"`
}

func ActionFSNames() carapace.Action {
	return carapace.ActionExecCommand("dfc", "-a", "-e", "json")(func(output []byte) carapace.Action {
		var fs struct {
			Filesystems []filesystem
		}
		if err := json.Unmarshal(output, &fs); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, f := range fs.Filesystems {
			vals = append(vals, f.Filesystem, f.MountPoint)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
