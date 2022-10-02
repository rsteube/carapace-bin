package action

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type pkgFile struct {
	Dependencies map[string]struct {
		Url     string `yaml:"url,omitempty"`
		Version string `yaml:"version,omitempty"`
		Path    string `yaml:"path,omitempty"`
	}
}

func ActionDependencies(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path := ""
		if f := cmd.Flag("project-root"); f != nil && f.Changed {
			path = f.Value.String()
		} else {
			path = c.Dir // TODO use preinvoke for project-root ?
		}
		location, err := util.FindReverse(path, "package.yaml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(location)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var pf pkgFile
		if err := yaml.Unmarshal(content, &pf); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for name, d := range pf.Dependencies {
			if d.Url != "" {
				vals = append(vals, name, d.Url)
			} else {
				vals = append(vals, name, d.Path)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
