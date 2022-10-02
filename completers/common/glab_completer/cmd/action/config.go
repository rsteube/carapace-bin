package action

import (
	"os"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
)

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if hosts, err := hosts(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return carapace.ActionValues(hosts...)
		}
	})
}

type glabConfig struct {
	Hosts map[string]interface{}
}

func hosts() ([]string, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	hosts := make([]string, 0)
	for host := range config.Hosts {
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func loadConfig() (config *glabConfig, err error) {
	var dir string
	if dir, err = os.UserConfigDir(); err == nil {
		var content []byte
		if content, err = os.ReadFile(dir + "/glab-cli/config.yml"); err == nil {
			err = yaml.Unmarshal(content, &config)
		}
	}
	return
}

func ActionConfigKeys() carapace.Action {
	return carapace.ActionValuesDescribed(
		"browser", "if unset, defaults to environment variables",
		"editor", "if unset, defaults to environment variables.",
		"gitlab_uri", "if unset, defaults to https://gitlab.com",
		"glab_pager", "the terminal pager program to send standard output to",
		"glamour_style", "Your desired markdown renderer style.",
		"token", "Your gitlab access token, defaults to environment variables",
		"visual", "alternative for editor. if unset, defaults to environment variables.",
	)
}

func ActionConfigValues(key string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		actions := map[string]carapace.Action{
			"token":         carapace.ActionValues(),
			"gitlab_uri":    carapace.ActionValues(),
			"glab_pager":    carapace.ActionValues("bat --style grid", "more", "most", "less"),
			"browser":       carapace.ActionFiles(),
			"editor":        carapace.ActionFiles(),
			"visual":        carapace.ActionFiles(),
			"glamour_style": carapace.ActionValues("dark", "light", "notty"),
		}

		return actions[key]
	})
}
