package action

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
)

func passwordStore() (string, error) {
	if val, exists := os.LookupEnv("PASSWORD_STORE_DIR"); !exists {
		if home, err := os.UserHomeDir(); err != nil {
			return "", err
		} else {
			return home + "/.password-store", nil
		}
	} else {
		return val, nil
	}
}

func ActionPassDirectories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path, err := passwordStore(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			return fs.ActionSubDirectories(path)
		}
	})
}

func ActionPassNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path, err := passwordStore(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			names := make([]string, 0)
			filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".gpg") {
					names = append(names, strings.TrimSuffix(strings.TrimPrefix(walkPath, path+"/"), ".gpg"))
				}
				return nil
			})
			return carapace.ActionValues(names...)
		}
	})
}
