package cmd

import (
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	os "github.com/rsteube/carapace-bin/pkg/actions/os"
)

var rootCmd = &cobra.Command{
	Use:                "carapace",
	Short:              "multi-shell multi-command argument completer",
	Long:               "https://github.com/rsteube/carapace-bin",
	DisableFlagParsing: true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func flagCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "carapace",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cmd.Flags().String("bridge", "", "generic completion bridge")
	cmd.Flags().BoolP("help", "h", false, "help for carapace")
	cmd.Flags().Bool("list", false, "list completers")
	cmd.Flags().String("macros", "", "list spec macros")
	cmd.Flags().String("spec", "", "spec completion")
	cmd.Flags().String("style", "", "set style")
	cmd.Flags().BoolP("version", "v", false, "version for carapace")

	if len(args) > 0 {
		if f := cmd.Flag(strings.TrimPrefix(args[0], "--")); len(args) > 1 || f != nil && f.Value.Type() == "bool" {
			cmd.Flags().VisitAll(func(f *pflag.Flag) {
				f.Changed = true // only one flag shall be completed so fake the changed state
			})
		}
	}

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"bridge": carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return os.ActionPathExecutables().Invoke(c).Suffix("/").ToA()
			case 1:
				return carapace.ActionValuesDescribed(
					"argcomplete", "kislyuk/argcomplete based completion like gcloud",
					"cobra", "spf13/cobra based completions like docker and kubectl",
					"fish", "completions registered in fish shell",
					"posener", "posener/complete based completion like hashicorp tools",
				)
			default:
				return carapace.ActionValues()
			}
		}),
		"macros": carapace.ActionExecCommand("carapace", "--macros")(func(output []byte) carapace.Action {
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
		}),
		"spec":  carapace.ActionFiles(".yaml"),
		"style": carapace.ActionStyleConfig(),
	})

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(args) > 0 && args[0] == "--macros" {
				c.Args = args[2:]
				return spec.ActionMacro("$_" + args[1]).Invoke(c).ToA()
			}
			return carapace.ActionValues()
		}),
	)

	return cmd
}

func posCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "carapace",
		DisableFlagParsing: true,
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		Run: func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).PositionalCompletion(
		ActionCompleters(),
		carapace.ActionStyledValues(
			"bash", "#d35673",
			"bash-ble", "#c2039a",
			"elvish", "#ffd6c9",
			"export", style.Default,
			"fish", "#7ea8fc",
			"ion", "#0e5d6d",
			"nushell", "#29d866",
			"oil", "#373a36",
			"powershell", "#e8a16f",
			"tcsh", "#412f09",
			"xonsh", "#a8ffa9",
			"zsh", "#efda53",
		),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(c.Args[0])
		}),
	)

	carapace.Gen(cmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{c.Args[0], "export", c.Args[0]}
			args = append(args, c.Args[3:]...)
			args = append(args, c.CallbackValue)
			return carapace.ActionExecCommand("carapace", args...)(func(output []byte) carapace.Action {
				if string(output) == "" {
					return carapace.ActionValues()
				}
				return carapace.ActionImport(output)
			})
		}),
	)

	return cmd
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.Batch(
					carapace.ActionExecute(flagCmd(c.Args)),
					carapace.ActionExecute(posCmd()),
				).ToA()
			}

			if strings.HasPrefix(c.Args[0], "-") {
				return carapace.ActionExecute(flagCmd(c.Args))
			}
			return carapace.ActionExecute(posCmd())
		}),
	)
}

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			re := regexp.MustCompile("^(?P<name>[^ ]+) +(?P<description>.*)$")

			vals := make([]string, 0)
			for _, line := range lines {
				if re.MatchString(line) {
					matched := re.FindStringSubmatch(line)
					vals = append(vals, matched[1])
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
