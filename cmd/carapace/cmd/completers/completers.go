package completers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/env"
	"github.com/rsteube/carapace/pkg/xdg"
	"gopkg.in/yaml.v3"
)

var (
	descriptions = make(map[string]string)
	names        = make([]string, 0)
)

func Names() []string {
	excludes := make(map[string]bool)
	for _, e := range env.Excludes() {
		excludes[e] = true
	}

	unique := map[string]bool{
		"carapace": true,
	}

	if os.Getenv("CARAPACE_ENV") != "0" {
		unique["get-env"] = true
		unique["set-env"] = true
		unique["unset-env"] = true
	}

	if _, ok := excludes["*"]; !ok {
		for _, name := range names {
			if _, ok := excludes[name]; !ok {
				unique[name] = true
			}
		}
	}
	//if specNames, err := Specs(); err == nil {
	if specNames := Specs(); true { // TODO use and handle err
		for _, name := range specNames {
			unique[name] = true
		}
	}

	combined := make([]string, 0, len(unique))
	for name := range unique {
		combined = append(combined, name)
	}
	sort.Strings(combined)
	return combined
}

func SpecPath(name string) (string, error) {
	userConfigDir, err := xdg.UserConfigDir()
	var path string
	if err == nil {
		path, err = specPath(userConfigDir, name)
		if err == nil {
			return path, nil
		}
	}

	if err != nil { // TODO have another look at the error handling here
		if !os.IsNotExist(err) {
			carapace.LOG.Println(err.Error())
		}

		configDirs, err := xdg.ConfigDirs()
		if err != nil {
			return "", err
		}
		for _, dir := range configDirs {
			path, err = specPath(dir, name)
			if err != nil {
				carapace.LOG.Println(err.Error())
				continue
			}
			return path, err
		}
	}
	return "", errors.New("no spec found")
}

func specPath(configDir string, name string) (string, error) {
	path := fmt.Sprintf("%v/carapace/specs/%v.yaml", configDir, name)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func OverlayPath(name string) (string, error) {
	// TODO code duplication
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	path := fmt.Sprintf("%v/carapace/overlays/%v.yaml", configDir, name)
	if _, err := os.Stat(path); err != nil {
		return "", err
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func Specs() []string {
	unique := make(map[string]bool)

	if userConfigDir, err := xdg.UserConfigDir(); err == nil {
		for _, spec := range specs(userConfigDir) {
			unique[spec] = true
		}
	}

	if configDirs, err := xdg.ConfigDirs(); err == nil {
		for _, dir := range configDirs {
			for _, spec := range specs(dir) {
				unique[spec] = true
			}
		}
	}

	result := make([]string, 0)
	for spec := range unique {
		result = append(result, spec)
	}
	return result
}

func specs(configDir string) []string {
	r := regexp.MustCompile(`^[0-9a-zA-Z_\-.]+\.yaml$`) // sanity check
	specs := make([]string, 0)
	dir := configDir + "/carapace/specs"
	if entries, err := os.ReadDir(dir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".yaml") && r.MatchString(entry.Name()) {
				specs = append(specs, strings.TrimSuffix(entry.Name(), ".yaml"))
			}
		}
	}
	return specs
}

func Description(name string) string {
	if d, err := specDescription(name); err == nil {
		return d
	}
	return descriptions[name]
}

func specDescription(name string) (string, error) {
	confDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(fmt.Sprintf("%v/carapace/specs/%v.yaml", confDir, name))
	if err != nil {
		return "", err
	}
	var s struct {
		Description string
	}

	err = yaml.Unmarshal(content, &s)
	if err != nil {
		return "", err
	}
	return s.Description, nil
}
