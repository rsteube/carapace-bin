package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type stack struct {
	Name             string
	Current          bool
	UpdateInProgress bool
}

type StackOpts struct {
	All              bool
	UpdateInProgress bool
}

func ActionStacks(cmd *cobra.Command, opts StackOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"--cwd", cmd.Flag("cwd").Value.String(), "stack", "ls", "--json"}
		if opts.All {
			args = append(args, "--all")
		}

		c.Setenv("PULUMI_SKIP_UPDATE_CHECK", "1")
		return carapace.ActionExecCommand("pulumi", args...)(func(output []byte) carapace.Action {
			var stacks []stack
			if err := json.Unmarshal(output, &stacks); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, s := range stacks {
				if opts.UpdateInProgress && !s.UpdateInProgress {
					continue
				}
				vals = append(vals, s.Name)
			}
			return carapace.ActionValues(vals...)
		}).Invoke(c).ToA()
	})
}
