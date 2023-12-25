package cmd

//go:generate go run ../../carapace-generate/gen.go

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/action"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/lazyinit"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/shim"
	"github.com/rsteube/carapace-bin/pkg/actions"
	spec "github.com/rsteube/carapace-spec"
	"github.com/rsteube/carapace/pkg/ps"
	"github.com/rsteube/carapace/pkg/xdg"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carapace [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]",
	Short: "multi-shell multi-command argument completer",
	Example: fmt.Sprintf(`  All completers and specs:
    bash:       source <(carapace _carapace bash)
    elvish:     eval (carapace _carapace elvish | slurp)
    fish:       carapace _carapace fish | source
    nushell:    carapace _carapace nushell
    oil:        source <(carapace _carapace oil)
    powershell: carapace _carapace powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace _carapace tcsh`+"`"+`
    xonsh:      exec($(carapace _carapace xonsh))
    zsh:        source <(carapace _carapace zsh)

  Single completer or spec:
    bash:       source <(carapace chmod bash)
    elvish:     eval (carapace chmod elvish | slurp)
    fish:       carapace chmod fish | source
    nushell:    carapace chmod nushell
    oil:        source <(carapace chmod oil)
    powershell: carapace chmod powershell | Out-String | Invoke-Expression
    tcsh:       eval `+"`"+`carapace chmod tcsh`+"`"+`
    xonsh:      exec($(carapace chmod xonsh))
    zsh:        source <(carapace chmod zsh)
  
  Style:
    set:        carapace --style 'carapace.Value=bold,magenta'
    clear:      carapace --style 'carapace.Description='

  Shell parameter is optional and if left out carapace will try to detect it by parent process name.
  Some completions are cached at [%v/carapace].
  Config is written to [%v/carapace].
  Specs are loaded from [%v/carapace/specs].
  `, suppressErr(xdg.UserCacheDir), suppressErr(xdg.UserConfigDir), suppressErr(xdg.UserConfigDir)),
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		// since flag parsing is disabled do this manually
		switch args[0] {
		case "--macro":
			macroCmd.SetArgs(args[1:])
			macroCmd.Execute()
		case "-h", "--help":
			cmd.Help()
		case "-v", "--version":
			println(cmd.Version)
		case "--list":
			listCmd.SetArgs(args[1:])
			listCmd.Execute()
		case "--run":
			runCmd.SetArgs(args[1:])
			runCmd.Execute()
		case "--schema":
			schemaCmd.SetArgs(args[1:])
			schemaCmd.Execute()
		case "--codegen":
			codegenCmd.SetArgs(args[1:])
			codegenCmd.Execute()
		case "--style":
			styleCmd.SetArgs(args[1:])
			styleCmd.Execute()
		default:
			invokeCmd.SetArgs(args)
			invokeCmd.Execute()
		}
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func suppressErr(f func() (string, error)) string { s, _ := f(); return s }

func createOverlayDir() error {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(fmt.Sprintf("%v/carapace/overlays", configDir), os.ModePerm)
}

func overlayPath(command string) (string, error) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	overlayPath := fmt.Sprintf("%v/carapace/overlays/%v.yaml", configDir, command)
	if _, err = os.Stat(overlayPath); err != nil {
		return "", err
	}
	return overlayPath, nil
}

func overlayCompletion(overlayPath string, args ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args[0] = "export" // TODO length check
		out, err := specCompletion(overlayPath, args...)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionImport([]byte(out))
	})
}

func updateSchema() error {
	confDir, err := xdg.UserConfigDir()
	if err != nil {
		return err
	}
	path := fmt.Sprintf("%v/carapace/schema.json", confDir)

	infoSchema, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	executable, err := os.Executable()
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	infoCarapace, err := os.Stat(executable)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if infoSchema == nil || !infoSchema.ModTime().Equal(infoCarapace.ModTime()) {
		schema, err := spec.Schema()
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		carapace.LOG.Printf("writing json schema to %#v", path)
		if err := os.WriteFile(path, []byte(schema), os.ModePerm); err != nil {
			return err
		}

		if err := os.Chtimes(path, time.Now(), infoCarapace.ModTime()); err != nil {
			return err
		}
	}
	return nil
}

func Execute(version string) error {
	rootCmd.Version = version

	if len(os.Args) > 1 {
		if os.Args[1] == "carapace" {
			os.Args[1] = "_carapace"
		} else if len(os.Args) < 4 && os.Args[1] == "_carapace" {
			shell := ps.DetermineShell()
			if len(os.Args) > 2 {
				shell = os.Args[2]
			}

			if err := shim.Update(); err != nil {
				println(err.Error()) // TODO fail / exit 1 ?
			}

			if err := updateSchema(); err != nil { // TODO do this only if needed
				println(err.Error())
			}

			if err := createOverlayDir(); err != nil { // TODO do this only if needed
				println(err.Error())
			}

			completers := completers.Names()
			if os.Getenv("CARAPACE_ENV") == "0" {
				filtered := make([]string, 0, len(completers))
				for _, name := range completers {
					switch name {
					case "get-env", "set-env", "unset-env":
					default:
						filtered = append(filtered, name)

					}
				}
				completers = filtered
			}

			println(shell)

			switch shell {
			case "bash":
				fmt.Println(lazyinit.Bash(completers))
			case "bash-ble":
				fmt.Println(lazyinit.BashBle(completers))
			case "elvish":
				fmt.Println(lazyinit.Elvish(completers))
			case "fish":
				fmt.Println(lazyinit.Fish(completers))
			case "nushell":
				fmt.Println(lazyinit.Nushell(completers))
			case "oil":
				fmt.Println(lazyinit.Oil(completers))
			case "powershell":
				fmt.Println(lazyinit.Powershell(completers))
			case "tcsh":
				fmt.Println(lazyinit.Tcsh(completers))
			case "xonsh":
				fmt.Println(lazyinit.Xonsh(completers))
			case "zsh":
				fmt.Println(lazyinit.Zsh(completers))
			default:
				// TODO
				// println("could not determine shell")
				return rootCmd.Execute()
			}
			return nil
		}
	}
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().Bool("codegen", false, codegenCmd.Short)
	rootCmd.Flags().BoolP("help", "h", false, "help for carapace")
	rootCmd.Flags().Bool("list", false, listCmd.Short)
	rootCmd.Flags().Bool("macro", false, macroCmd.Short)
	rootCmd.Flags().Bool("run", false, runCmd.Short)
	rootCmd.Flags().Bool("schema", false, schemaCmd.Short)
	rootCmd.Flags().Bool("style", false, styleCmd.Short)
	rootCmd.Flags().BoolP("version", "v", false, "version for carapace")

	rootCmd.MarkFlagsMutuallyExclusive(
		"codegen",
		"help",
		"list",
		"macro",
		"run",
		"schema",
		"style",
		"version",
	)

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "-") {
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(rootCmd.Flags())
				return carapace.ActionExecute(cmd)
			}
			return action.ActionCompleters()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !strings.HasPrefix(c.Args[0], "-") {
				return carapace.ActionExecute(invokeCmd)
			}
			switch c.Args[0] {
			case "--codegen":
				return carapace.ActionExecute(codegenCmd).Shift(1)
			case "--help":
				return carapace.ActionValues()
			case "--list":
				return carapace.ActionExecute(listCmd).Shift(1).Usage("list")
			case "--macro":
				return carapace.ActionExecute(macroCmd).Shift(1)
			case "--run":
				return carapace.ActionExecute(runCmd).Shift(1)
			case "--schema":
				return carapace.ActionExecute(schemaCmd).Shift(1)
			case "--style":
				return carapace.ActionExecute(styleCmd).Shift(1)
			case "-v", "--version":
				return carapace.ActionValues()
			default:
				return carapace.ActionValues()
			}
		}),
	)

	for m, f := range actions.MacroMap {
		spec.AddMacro(m, f)
	}
	spec.Register(rootCmd)
}
