package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func actionExecVercel(cmd *cobra.Command, arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{}
			if f := cmd.Root().Flag("local-config"); f.Changed {
				args = append(args, "--local-config", f.Value.String())
			}
			if f := cmd.Root().Flag("global-config"); f.Changed {
				args = append(args, "--global-config", f.Value.String())
			}
			if f := cmd.Root().Flag("token"); f.Changed {
				args = append(args, "--token", f.Value.String())
			}
			args = append(args, arg...)

			return carapace.ActionExecCommand("vercel", args...)(func(output []byte) carapace.Action {
				return f(output)
			})
		})
	}
}
