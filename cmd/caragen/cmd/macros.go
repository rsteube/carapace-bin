package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var macrosCmd = &cobra.Command{
	Use:   "macros",
	Short: "list macros",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := json.Marshal(macros())
		if err != nil {
			return err
		}

		fmt.Fprintln(cmd.OutOrStdout(), string(m))
		return nil
	},
}

func init() {
	carapace.Gen(macrosCmd).Standalone()

	rootCmd.AddCommand(macrosCmd)
}

type macro struct {
	Path        string
	ImportPath  string
	Package     string
	Function    string
	Description string
	Signature   string
}

func (m macro) FlagParsingDisabled() bool {
	return m.Package == "bridge"
}

func macros() map[string]macro {
	root, err := rootDir()
	if err != nil {
		panic(err.Error)
	}

	macros := make(map[string]macro)
	descriptions := make(map[string]string)
	r := regexp.MustCompile(`^func Action(?P<name>[^(]+)\((?P<arg>[^(]*)\) carapace.Action {$`)
	filepath.WalkDir(root+"/pkg/actions", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err.Error())
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				_macro := macro{}
				_macro.Package = strings.Replace(filepath.Dir(strings.TrimPrefix(path, root+"/pkg/actions/")), "/", ".", -1)

				if t := scanner.Text(); strings.HasPrefix(t, "func Action") {
					if matches := r.FindStringSubmatch(t); matches != nil {
						_macro.ImportPath = fmt.Sprintf("github.com/rsteube/carapace-bin/pkg/actions/%v", strings.Replace(_macro.Package, ".", "/", -1))
						_macro.Function = fmt.Sprintf("Action%v", matches[1])
						_macro.Signature = matches[2]
						macros[matches[1]] = _macro
					}
				} else if strings.HasPrefix(t, "// Action") {
					if splitted := strings.SplitN(strings.TrimPrefix(t, "// Action"), " ", 2); len(splitted) > 1 {
						descriptions[_macro.Package+"."+splitted[0]] = splitted[1]
					}
				}
			}

		}
		return nil
	})
	return macros
}

func (m macro) snippetCall() string {
	if arg := m.Signature; strings.Contains(arg, ",") {
		return "// TODO unsupported signature: " + m.Signature
	} else if arg == "" {
		return fmt.Sprintf(`"%v.%v": spec.MacroN(%v)%v,`, "TODO", "TODO", "TODO", m.FlagParsingDisabled())
	} else if strings.Contains(arg, "...") {
		return fmt.Sprintf(`"%v.%v": spec.MacroV(%v)%v,`, "TODO", "TODO", "TODO", m.FlagParsingDisabled())
	} else {
		return fmt.Sprintf(`"%v.%v": spec.MacroI(%v)%v,`, "TODO", "TODO", "TODO", m.FlagParsingDisabled())
	}
}

func (m macro) noFlag() string {
	if m.ImportPath == "github.com/rsteube/carapace-bin/pkg/actions/fs" {
		return ".NoFlag()"
	}
	return ""
}
