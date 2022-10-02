package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "xargs",
	Short:              "build and execute command lines from standard input",
	Long:               "https://linux.die.net/man/1/xargs",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			firstPass := xargsCmd()
			firstPass.FParseErrWhitelist = cobra.FParseErrWhitelist{
				UnknownFlags: true,
			}
			args := []string{}
			args = append(args, c.Args...)
			args = append(args, c.CallbackValue)
			if err := firstPass.ParseFlags(args); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			secondPass := xargsCmd()
			if args := firstPass.Flags().Args(); len(args) > 0 && !(len(args) == 1 && args[0] == c.CallbackValue) {
				subcmd := subCmd(args[0])
				secondPass.AddCommand(subcmd)
			}
			return carapace.ActionExecute(secondPass)
		}),
	)
}

func xargsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "xargs",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()

	cmd.Flags().StringS("E", "E", "", "set logical EOF string")
	cmd.Flags().StringS("I", "I", "", "same as --replace=R")
	cmd.Flags().StringP("arg-file", "a", "", "read arguments from FILE")
	cmd.Flags().StringP("delimiter", "d", "", "items in input stream are separated by CHARACTER")
	cmd.Flags().StringP("eof", "e", "", "equivalent to -E END if END is specified")
	cmd.Flags().BoolP("exit", "x", false, "exit if the size (see -s) is exceeded")
	cmd.Flags().Bool("help", false, "display this help and exit")
	cmd.Flags().BoolP("interactive", "p", false, "prompt before running commands")
	cmd.Flags().StringS("l", "l", "", "similar to -L but defaults to at most one non- blank input line")
	cmd.Flags().StringP("max-args", "n", "", "use at most MAX-ARGS arguments per command line")
	cmd.Flags().StringP("max-chars", "s", "", "limit length of command line to MAX-CHARS")
	cmd.Flags().StringP("max-lines", "L", "", "use at most MAX-LINES non-blank input lines per command line")
	cmd.Flags().StringP("max-procs", "P", "", "run at most MAX-PROCS processes at a time")
	cmd.Flags().BoolP("no-run-if-empty", "r", false, "if there are no arguments, then do not run COMMAND")
	cmd.Flags().BoolP("null", "0", false, "items are separated by a null")
	cmd.Flags().BoolP("open-tty", "o", false, "Reopen stdin as /dev/tty in the child process")
	cmd.Flags().String("process-slot-var", "", "set environment variable VAR in child processes")
	cmd.Flags().StringP("replace", "i", "", "replace R in INITIAL-ARGS with names")
	cmd.Flags().Bool("show-limits", false, "show limits on command-line length")
	cmd.Flags().BoolP("verbose", "t", false, "print commands before executing them")
	cmd.Flags().Bool("version", false, "output version information and exit")

	cmd.Flag("eof").NoOptDefVal = " "
	cmd.Flag("replace").NoOptDefVal = "{}"

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"arg-file": carapace.ActionFiles(),
	})

	carapace.Gen(cmd).PositionalCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	return cmd
}

func subCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                name,
		Run:                func(cmd *cobra.Command, args []string) {},
		DisableFlagParsing: true,
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(name),
	)
	return cmd
}
