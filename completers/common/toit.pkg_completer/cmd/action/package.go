package action

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
)

type pkg struct {
	Name        string
	Description string
	Version     string
	Url         string
}

func packageAction(f func(map[string]pkg) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("toitpkg", "pkg", "list", "--output", "json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		pkgs := make(map[string]pkg)

		var p pkg
		for _, line := range lines {
			if strings.HasPrefix(line, "{") {
				if err := json.Unmarshal([]byte(line), &p); err != nil {
					return carapace.ActionMessage(err.Error())
				}
				pkgs[p.Name] = p
			}
		}
		return f(pkgs)
	})
}

func ActionPackages() carapace.Action {
	return packageAction(func(m map[string]pkg) carapace.Action {
		vals := make([]string, 0)
		for _, p := range m {
			vals = append(vals, p.Name, p.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

func ActionPackageVersions(id string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(id, "github.com/") {
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + id, Tags: true})
		}

		if strings.Contains(id, "/") && !strings.Contains(id, ".") { // assume github package
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://github.com/" + id, Tags: true})
		}

		return packageAction(func(m map[string]pkg) carapace.Action {
			if p, ok := m[id]; ok {
				if strings.HasPrefix(p.Url, "github.com/") {
					return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + p.Url, Tags: true})
				}
			}
			return carapace.ActionValues()
		})
	})
}
