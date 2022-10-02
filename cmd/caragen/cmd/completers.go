package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completersCmd = &cobra.Command{
	Use:       "completers",
	Short:     "list completers",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"common", "darwin", "linux", "unix", "windows"},
	RunE: func(cmd *cobra.Command, args []string) error {
		completers, err := applicableCompleters(args[0])
		if err != nil {
			return err
		}

		m, err := json.Marshal(completers)
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(m))
		return nil
	},
}

func init() {
	carapace.Gen(completersCmd).Standalone()

	rootCmd.AddCommand(completersCmd)

	carapace.Gen(completersCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"common", "all ${GOOS}",
			"darwin", "common, unix and macOS",
			"linux", "common, unix and linux",
			"unix", "common, linux and macOS",
			"windows", "common and windows",
		),
	)
}

type completer struct {
	Description string
	Path        string
}

func applicableCompleters(goos string) (map[string]completer, error) {
	gooses := make([]string, 0)
	switch goos {
	case "common":
		gooses = append(gooses, "common")
	case "darwin":
		gooses = append(gooses, "common", "unix", "darwin")
	case "linux":
		gooses = append(gooses, "common", "unix", "linux")
	case "unix":
		gooses = append(gooses, "common", "unix")
	case "windows":
		gooses = append(gooses, "common", "windows")
	}

	mergedCompleters := make(map[string]completer, 0)
	for _, goos := range gooses {
		completers, err := readCompleters(goos)
		if err != nil {
			return nil, err
		}

		for name, completer := range completers {
			mergedCompleters[name] = completer
		}
	}
	return mergedCompleters, nil
}

func readCompleters(goos string) (map[string]completer, error) {
	completers := make(map[string]completer)
	if root, err := rootDir(); err != nil {
		return nil, err
	} else {
		dir := fmt.Sprintf("%v/completers/%v", root, goos)
		if files, err := os.ReadDir(dir); err != nil {
			return nil, err

		} else {
			for _, file := range files {
				if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
					name := strings.TrimSuffix(file.Name(), "_completer")
					completers[name] = completer{
						Description: readDescription(fmt.Sprintf("%v/%v", dir, file.Name())),
						Path:        dir,
					}
				}
			}
		}
	}
	return completers, nil
}

func readDescription(path string) string {
	if content, err := os.ReadFile(fmt.Sprintf("%v/cmd/root.go", path)); err == nil {
		re := regexp.MustCompile("^\tShort: +\"(?P<description>.*)\",$")
		for _, line := range strings.Split(string(content), "\n") {
			if re.MatchString(line) {
				return re.FindStringSubmatch(line)[1]
			}
		}
	}
	return ""
}
