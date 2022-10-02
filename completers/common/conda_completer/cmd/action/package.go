package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type pkg struct {
	Name    string
	Version string
}

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"list", "--json"}
		if name := cmd.Flag("name"); name != nil && name.Changed {
			args = append(args, "--name", name.Value.String())

		}
		if prefix := cmd.Flag("prefix"); prefix != nil && prefix.Changed {
			args = append(args, "--prefix", prefix.Value.String())

		}
		return carapace.ActionExecCommand("conda", args...)(func(output []byte) carapace.Action {
			var pkgs []pkg
			if err := json.Unmarshal(output, &pkgs); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, p := range pkgs {
				vals = append(vals, p.Name, p.Version)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
