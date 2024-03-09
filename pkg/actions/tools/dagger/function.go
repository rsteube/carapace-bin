package dagger

import (
	"encoding/json"
	"os"
	"time"
	"unicode"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/cache"
	"github.com/rsteube/carapace/pkg/cache/key"
	"github.com/spf13/cobra"
)

type function struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Args        []struct {
		Name        string `json:"Name"`
		Description string `json:"Description"`
		TypeDef     struct {
			Kind     string `json:"Kind"`
			Optional bool   `json:"Optional"`
			AsObject *struct {
				Name             string `json:"Name"`
				Functions        any    `json:"Functions"`
				Fields           any    `json:"Fields"`
				Constructor      any    `json:"Constructor"`
				SourceModuleName string `json:"SourceModuleName"`
			} `json:"AsObject"`
			AsInterface any `json:"AsInterface"`
			AsInput     any `json:"AsInput"`
			AsList      any `json:"AsList"`
		} `json:"TypeDef"`
		DefaultValue string `json:"DefaultValue"`
	} `json:"Args"`
}

type daggerFile struct {
	Name   string
	Sdk    string
	Source string
}

// ActionFunctions completes functions and their arguments
//
//	grep-dir (dagger call grep-dir --directory-arg . --pattern GrepDir)
//	--directory-arg
func ActionFunctions() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile("dagger.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var d daggerFile
		if err := json.Unmarshal(content, &d); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		output, err := cache.Cache(24*time.Hour, key.FolderStats(d.Source))(func() ([]byte, error) {
			return c.Command("dagger", "functions", "--json").Output()
		})
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var functions []function
		if err := json.Unmarshal(output, &functions); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		cmd := &cobra.Command{}
		carapace.Gen(cmd).Standalone()
		for _, f := range functions {
			f.Name = toKebab(f.Name)
			subCmd := &cobra.Command{
				Use:   f.Name,
				Short: f.Description,
				Run:   func(cmd *cobra.Command, args []string) {},
			}
			carapace.Gen(subCmd).Standalone()

			for _, arg := range f.Args {
				arg.Name = toKebab(arg.Name)

				// TODO transform camelcase to kebab
				switch arg.TypeDef.Kind { // TODO more types
				case "STRING_KIND", "OBJECT_KIND":
					subCmd.Flags().String(arg.Name, arg.DefaultValue, arg.Description)
				default:
					return carapace.ActionMessage("unknown kind %s", arg.TypeDef.Kind)
				}

				if arg.TypeDef.Optional {
					subCmd.Flag(arg.Name).NoOptDefVal = " "
				}

				localArg := arg
				carapace.Gen(subCmd).FlagCompletion(carapace.ActionMap{
					arg.Name: carapace.ActionCallback(func(c carapace.Context) carapace.Action {
						if localArg.TypeDef.AsObject == nil {
							return carapace.ActionValues()
						}

						switch localArg.TypeDef.AsObject.Name { // TODO more
						case "Directory":
							return carapace.ActionDirectories()
						case "File":
							return carapace.ActionFiles()
						default:
							return carapace.ActionValues()
						}
					}),
				})
			}

			cmd.AddCommand(subCmd)
		}

		return carapace.ActionExecute(cmd)
	})
}

func toKebab(s string) string {
	runes := make([]rune, 0)
	for _, r := range []rune(s) {
		if unicode.IsUpper(r) {
			runes = append(runes, '-')
		}
		runes = append(runes, unicode.ToLower(r))
	}
	return string(runes)
}
