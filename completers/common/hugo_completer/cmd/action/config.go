package action

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/rsteube/carapace-bin/pkg/util"
	"gopkg.in/yaml.v3"
)

func configPath() (path string, err error) {
	var wd string
	if wd, err = os.Getwd(); err == nil {
		if path, err = util.FindReverse(wd, "config.toml"); err != nil {
			if path, err = util.FindReverse(wd, "config.yaml"); err != nil {
				path, err = util.FindReverse(wd, "config.json")
			}
		}
	}
	return
}

type config struct {
	Deployment struct {
		Targets []struct {
			Name string
			Url  string `json:"URL"`
		}
	}
}

func loadConfig() (c config, err error) {
	var path string
	if path, err = configPath(); err == nil {
		var content []byte
		if content, err = os.ReadFile(path); err == nil {
			switch filepath.Ext(path) {
			case ".json":
				err = json.Unmarshal(content, &c)
			case ".toml":
				err = toml.Unmarshal(content, &c)
			case ".yaml":
				err = yaml.Unmarshal(content, &c)
			case ".yml":
				err = yaml.Unmarshal(content, &c)
			}
		}
	}
	return
}
