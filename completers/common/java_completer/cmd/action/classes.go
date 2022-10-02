package action

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

func ActionClasspathClasses(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		paths := make([]string, 0)
		if f := cmd.Flag("cp"); f.Changed {
			paths = append(paths, strings.Split(f.Value.String(), ":")...)
		}
		if f := cmd.Flag("classpath"); f.Changed {
			paths = append(paths, strings.Split(f.Value.String(), ":")...)
		}
		if f := cmd.Flag("jar"); f.Changed {
			paths = append(paths, f.Value.String())
		}
		if value, exists := os.LookupEnv("CLASSPATH"); exists {
			paths = append(paths, strings.Split(value, ":")...)
		}

		files := make([]string, 0)
		for _, path := range paths {
			if f, err := os.Stat(path); err == nil {
				if !f.IsDir() {
					files = append(files, path)
				} else {
					if fileInfos, err := os.ReadDir(path); err == nil {
						for _, file := range fileInfos {
							files = append(files, fmt.Sprintf("%v/%v", path, file.Name()))
						}
					}
				}
			}
		}

		actions := make([]carapace.Action, 0)
		for _, file := range files {
			if strings.HasSuffix(file, ".jar") ||
				strings.HasSuffix(file, ".zip") {
				actions = append(actions, fs.ActionJarFileClasses(file))
			}
		}
		return carapace.Batch(actions...).Invoke(c).Merge().ToMultiPartsA(".")
	})
}

func ActionLocalClasses() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := os.ReadDir(".")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, entry := range entries {
			if strings.HasSuffix(entry.Name(), ".class") {
				vals = append(vals, strings.TrimSuffix(entry.Name(), ".class"))
			}
		}
		return carapace.ActionValues(vals...)
	})
}
