package action

import (
	"encoding/json"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

func ActionPackages(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageSearch(cmd.Flag("registry").Value.String())
	})
}

func ActionPackageNames(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageNames(cmd.Flag("registry").Value.String())
	})
}

func ActionPackageVersions(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageVersions(npm.PackageOpts{Registry: cmd.Flag("registry").Value.String(), Package: pkg})
	})
}

func ActionPackageTags(cmd *cobra.Command, pkg string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return npm.ActionPackageTags(npm.PackageOpts{Registry: cmd.Flag("registry").Value.String(), Package: pkg})
	})
}

type packageJson struct {
	Scripts    map[string]string
	Workspaces []string
}

func loadPackageJson(c carapace.Context) (pj packageJson, err error) {
	var packageFile string
	if packageFile, err = util.FindReverse(c.Dir, "package.json"); err == nil {
		var content []byte
		if content, err = os.ReadFile(packageFile); err == nil {
			err = json.Unmarshal(content, &pj)
		}
	}
	return
}

func ActionScripts(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if pj, err := loadPackageJson(c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			vals := make([]string, 0)
			for name := range pj.Scripts {
				vals = append(vals, name)
			}
			return carapace.ActionValues(vals...)
		}
	})
}
