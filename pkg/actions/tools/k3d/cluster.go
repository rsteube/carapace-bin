package k3d

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type cluster struct {
	Name           string
	ImageVolume    string
	ServersRunning int
	ServersCount   int
}

func (c cluster) style() string {
	switch {
	case c.ServersRunning == 0:
		return style.Carapace.KeywordNegative
	case c.ServersRunning < c.ServersCount:
		return style.Carapace.KeywordAmbiguous
	default:
		return style.Carapace.KeywordPositive
	}
}

func ActionClusters() carapace.Action {
	return carapace.ActionExecCommand("k3d", "cluster", "list", "--output", "json")(func(output []byte) carapace.Action {
		var clusters []cluster
		if err := json.Unmarshal(output, &clusters); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, c := range clusters {
			vals = append(vals, c.Name, c.ImageVolume, c.style())
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
