package k3d

import (
	"encoding/json"
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type NodeOpts struct {
	Cluster string
	Running bool
	Stopped bool
}

func (n NodeOpts) Default() NodeOpts {
	n.Running = true
	n.Stopped = true
	return n
}

type node struct {
	Name          string
	Role          string
	RuntimeLabels struct {
		K3dCluster string `json:"k3d.cluster"`
	}
	State struct {
		Running bool
	}
}

func (n node) applies(opts NodeOpts) bool {
	switch {
	case opts.Cluster != "" && opts.Cluster != n.RuntimeLabels.K3dCluster:
		return false
	case opts.Running && n.State.Running:
		return true
	case opts.Stopped && !n.State.Running:
		return true
	default:
		return false
	}
}

func (n node) style() string {
	if n.State.Running {
		return style.Carapace.KeywordPositive
	}
	return style.Carapace.KeywordNegative
}

func ActionNodes(opts NodeOpts) carapace.Action {
	return carapace.ActionExecCommand("k3d", "node", "list", "--output", "json")(func(output []byte) carapace.Action {
		var nodes []node
		if err := json.Unmarshal(output, &nodes); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, n := range nodes {
			if n.applies(opts) {
				vals = append(vals, n.Name, fmt.Sprintf("%v.%v", n.RuntimeLabels.K3dCluster, n.Role), n.style())
			}
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}

func ActionNodeGroups() carapace.Action {
	return carapace.ActionValues("server", "servers", "agent", "agents", "loadbalancer", "all")
}

func ActionNodeFilter() carapace.Action { // TODO limit to specific cluster
	return carapace.ActionMultiPartsN(":", 3, func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionNodeGroups()
		case 1:
			return carapace.ActionValues() // TODO subset
		default:
			return carapace.ActionValues() // TODO suffix
		}
	})
}
