package cmd

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions"
	spec "github.com/rsteube/carapace-spec"
	"github.com/spf13/cobra"
)

var macroCmd = &cobra.Command{
	Use:   "--macro [macro] ...",
	Short: "list or execute macros",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			printMacros()
		case 1:
			printMacro(args[0])
		default:
			// TODO macro args
		}
	},
}

func init() {
	carapace.Gen(macroCmd).Standalone()
	macroCmd.Flags().SetInterspersed(false)

	carapace.Gen(macroCmd).PositionalCompletion(
		carapace.ActionMultiPartsN("(", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionExecCommand("carapace", "--macro")(func(output []byte) carapace.Action {
					lines := strings.Split(string(output), "\n")

					vals := make([]string, 0)
					for _, line := range lines[:len(lines)-1] {
						if fields := strings.Fields(line); len(fields) > 1 {
							vals = append(vals, fields[0], strings.Join(fields[1:], " "))
						} else {
							vals = append(vals, fields[0], "")
						}
					}
					return carapace.ActionValuesDescribed(vals...).Invoke(carapace.Context{}).ToMultiPartsA(".")
				})
			default:
				if m, ok := actions.MacroMap[c.Parts[0]]; ok {
					return carapace.ActionValues(m.Signature() + ")")
				}
				return carapace.ActionValues(")")
			}
		}),
	)

	carapace.Gen(macroCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return spec.ActionMacro("$carapace." + c.Args[0]).Shift(1) // TODO macro prefix
		}),
	)
}

func printMacros() {
	maxlen := 0
	names := make([]string, 0)
	for name := range actions.MacroMap {
		names = append(names, name)
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", name, actions.MacroDescriptions[name])
	}
}

func printMacro(name string) {
	if m, ok := actions.MacroMap[name]; ok {
		path := strings.Replace(name, ".", "/", -1)
		signature := ""
		if s := m.Signature(); s != "" {
			signature = fmt.Sprintf("(%v)", s)
		}

		fmt.Printf(`signature:   $carapace.%v%v
description: %v
reference:   https://pkg.go.dev/github.com/rsteube/carapace-bin/pkg/actions/%v#Action%v
`, name, signature, actions.MacroDescriptions[name], filepath.Dir(path), filepath.Base(path))
	}
}
