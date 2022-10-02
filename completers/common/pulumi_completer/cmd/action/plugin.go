package action

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/third_party/golang.org/x/sys/execabs"
	"github.com/spf13/cobra"
)

type plugin struct {
	Name    string
	Kind    string
	Version string
}

func plugins(cwd string) (plugins []plugin, err error) {
	stderr := bytes.Buffer{}
	cmd := execabs.Command("pulumi", "--cwd", cwd, "plugin", "ls", "--json")
	cmd.Stderr = &stderr
	cmd.Env = append(cmd.Env, "PULUMI_SKIP_UPDATE_CHECK=1")
	if output, cmdErr := cmd.Output(); cmdErr != nil {
		err = errors.New(stderr.String())
	} else {
		err = json.Unmarshal(output, &plugins)
	}
	return
}

func ActionPluginKinds(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if plugins, err := plugins(cmd.Flag("cwd").Value.String()); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			kinds := make(map[string]bool)
			for _, plugin := range plugins {
				kinds[plugin.Kind] = true
			}

			vals := make([]string, 0)
			for kind := range kinds {
				vals = append(vals, kind)
			}
			return carapace.ActionValues(vals...)
		}
	})
}

func ActionPluginNames(cmd *cobra.Command, kind string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if plugins, err := plugins(cmd.Flag("cwd").Value.String()); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			names := make(map[string]bool)
			for _, plugin := range plugins {
				if kind == plugin.Kind {
					names[plugin.Name] = true
				}
			}

			vals := make([]string, 0)
			for name := range names {
				vals = append(vals, name)
			}
			return carapace.ActionValues(vals...)
		}
	})
}

func ActionPluginVersions(cmd *cobra.Command, kind string, name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if plugins, err := plugins(cmd.Flag("cwd").Value.String()); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			versions := make(map[string]bool)
			for _, plugin := range plugins {
				if kind == plugin.Kind &&
					name == plugin.Name {
					versions[plugin.Version] = true
				}
			}

			vals := make([]string, 0)
			for version := range versions {
				vals = append(vals, version)
			}
			return carapace.ActionValues(vals...)
		}
	})
}
