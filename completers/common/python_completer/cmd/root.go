package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/python_completer/cmd/action"
	"github.com/rsteube/carapace-bin/completers/python_completer/cmd/module"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "python",
	Short: "an interpreted, interactive, object-oriented programming language",
	Long:  "https://www.python.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	moduleFlagIndex := -1
	for index, arg := range os.Args {
		if arg == "-m" {
			moduleFlagIndex = index
			break
		}
	}

	if moduleFlagIndex != -1 && moduleFlagIndex < len(os.Args)-2 {
		patchedArgs := make([]string, 0)
		patchedArgs = append(patchedArgs, os.Args[:moduleFlagIndex+2]...)
		patchedArgs = append(patchedArgs, "--") // fake dash arg
		patchedArgs = append(patchedArgs, os.Args[moduleFlagIndex+2:]...)
		os.Args = patchedArgs
	}

	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "don't write .pyc files on import")
	rootCmd.Flags().BoolS("E", "E", false, "ignore PYTHON* environment variables")
	rootCmd.Flags().BoolS("I", "I", false, "isolate Python from the user's environment")
	rootCmd.Flags().BoolSliceS("O", "O", []bool{}, "remove assert and __debug__-dependent statements")
	rootCmd.Flags().BoolS("S", "S", false, "don't imply 'import site' on initialization")
	rootCmd.Flags().BoolSliceS("V", "V", []bool{}, "print the Python version number and exit")
	rootCmd.Flags().StringS("W", "W", "", "warning control")
	rootCmd.Flags().StringArrayS("X", "X", []string{}, "set implementation-specific option")
	rootCmd.Flags().BoolS("b", "b", false, "issue warnings about str(bytes_instance), str(bytearray_instance) and comparing bytes/bytearray with str")
	rootCmd.Flags().StringS("c", "c", "", "program passed in as string")
	rootCmd.Flags().String("check-hash-based-pycs", "", "control how Python invalidates hash-based .pyc files")
	rootCmd.Flags().BoolS("d", "d", false, "turn on parser debugging output")
	rootCmd.Flags().BoolS("h", "h", false, "print this help message and exit")
	rootCmd.Flags().BoolS("i", "i", false, "inspect interactively after running script")
	rootCmd.Flags().StringS("m", "m", "", "run library module as a script")
	rootCmd.Flags().BoolS("q", "q", false, "don't print version and copyright messages on interactive startup")
	rootCmd.Flags().BoolS("s", "s", false, "don't add user site directory to sys.path")
	rootCmd.Flags().BoolS("u", "u", false, "force the stdout and stderr streams to be unbuffered")
	rootCmd.Flags().BoolSliceS("v", "v", []bool{}, "verbose (trace import statements)")
	rootCmd.Flags().BoolS("x", "x", false, "skip first line of source")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"X": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValuesDescribed(
					"faulthandler", "enable faulthandler",
					"oldparser", "enable the traditional LL(1) parser; also PYTHONOLDPARSER",
					"showrefcount", "output the total reference count and number of used memory",
					"tracemalloc", "start tracing Python memory allocations using the tracemalloc module",
					"importtime", "show how long each import takes",
					"dev", "enable CPython's \"development mode\"",
					"utf8", "enable UTF-8 mode for operating system interfaces",
					"pycache_prefix=", "enable writing .pyc files to a parallel tree",
				)
			case 1:
				switch c.Parts[0] {
				case "pycache_prefix":
					return carapace.ActionDirectories()
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"check-hash-based-pycs": carapace.ActionValues("always", "default", "never").StyleF(style.ForKeyword),
		"m":                     action.ActionModules(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch rootCmd.Flag("m").Value.String() {
			case "http.server":
				return module.ActionInvokeHttpServer()
			case "venv":
				return module.ActionInvokeVenv()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
