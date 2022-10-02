package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type tasks struct {
	Nodes []struct {
		Label string
		Type  string
	}
}

func ActionTasks(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--tasks-json"}
		if f := cmd.Flag("cwd"); f.Changed {
			args = append(args, "--cwd", f.Value.String())
		}
		if f := cmd.Flag("gulpfile"); f.Changed {
			args = append(args, "--gulpfile", f.Value.String())
		}

		return carapace.ActionExecCommand("gulp", args...)(func(output []byte) carapace.Action {
			var t tasks
			if err := json.Unmarshal(output, &t); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			vals := make([]string, 0)
			for _, node := range t.Nodes {
				if node.Type == "task" {
					vals = append(vals, node.Label)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
