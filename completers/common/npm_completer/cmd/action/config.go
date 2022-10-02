package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionConfigKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {

		args := []string{"config", "list", "--json"}
		if flag := cmd.Flag("global"); flag != nil && flag.Changed {
			args = append(args, "--global")
		}

		return carapace.ActionExecCommand("npm", args...)(func(output []byte) carapace.Action {
			var config map[string]interface{}
			if err := json.Unmarshal(output, &config); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(config))
			for key := range config {
				vals = append(vals, key)
			}
			return carapace.ActionValues(vals...)
		})
	})
}
