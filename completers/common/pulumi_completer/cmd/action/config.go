package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type config struct {
	Value  string
	Secret bool
}

func ActionConfigKeys(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		c.Setenv("PULUMI_SKIP_UPDATE_CHECK", "1")
		return carapace.ActionExecCommand("pulumi", "--cwd", cmd.Flag("cwd").Value.String(), "config", "--json")(func(output []byte) carapace.Action {
			var config map[string]config
			if err := json.Unmarshal(output, &config); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for name, c := range config {
				vals = append(vals, name, c.Value)
			}
			return carapace.ActionValuesDescribed(vals...).Invoke(c).ToMultiPartsA(":")
		}).Invoke(c).ToA()
	})
}
