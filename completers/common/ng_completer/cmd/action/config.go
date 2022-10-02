package action

import (
	"encoding/json"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

type config struct {
	Projects map[string]struct {
		Architect map[string]struct {
			Configurations map[string]interface{}
		}
	}
}

func actionConfig(f func(cfg config) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "angular.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var cfg config
		if err := json.Unmarshal(content, &cfg); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(cfg)
	})
}

func ActionProjects() carapace.Action {
	return actionConfig(func(cfg config) carapace.Action {
		vals := make([]string, 0)
		for name := range cfg.Projects {
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionTargets(project string) carapace.Action {
	return actionConfig(func(cfg config) carapace.Action {
		vals := make([]string, 0)
		if p, ok := cfg.Projects[project]; ok {
			for name := range p.Architect {
				vals = append(vals, name)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

func ActionBuilderConfigurations(project string, target string) carapace.Action {
	return actionConfig(func(cfg config) carapace.Action {
		vals := make([]string, 0)
		if p, ok := cfg.Projects[project]; ok {
			for t, architect := range p.Architect {
				if target == "" || target == t {
					for name := range architect.Configurations {
						vals = append(vals, name)
					}
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
