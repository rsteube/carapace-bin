package action

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func ActionHostPort() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return carapace.ActionValues() // TODO local ips
		case 1:
			return net.ActionPorts()
		default:
			return carapace.ActionValues()
		}
	})
}

func ActionLintRules() carapace.Action {
	return carapace.ActionExecCommand("deno", "lint", "--rules")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines {
			if strings.HasPrefix(line, " - ") {
				vals = append(vals, strings.TrimPrefix(line, " - "))
			}
		}
		return carapace.ActionValues(vals...)
	})
}

type config struct {
	Tasks map[string]string
}

func ActionTasks(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			var err error
			if path, err = util.FindReverse(c.Dir, "deno.json"); err != nil {
				if path, err = util.FindReverse(c.Dir, "deno.jsonc"); err != nil {
					return carapace.ActionMessage(err.Error())
				}
			}
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var _config config
		if err := json.Unmarshal(content, &_config); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for k, v := range _config.Tasks {
			vals = append(vals, k, v)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
