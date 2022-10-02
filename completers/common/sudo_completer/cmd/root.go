package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                "sudo",
	Short:              "execute a command as another user",
	Long:               "https://linux.die.net/man/8/sudo",
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
			// sudo stops parsing it's own flags on the first positional argument.
			// To simulate this with pflag arguments are first parsed with unknown flags being allowed.
			// If a positional argument is given (and not the current word being completed) it is used to add a fake subcommand.
			// The subcommand then handles completion of the sudo'ed command.
			firstPass := sudoCmd()
			firstPass.FParseErrWhitelist = cobra.FParseErrWhitelist{
				UnknownFlags: true,
			}
			args := []string{}
			args = append(args, c.Args...)
			args = append(args, c.CallbackValue)
			if err := firstPass.ParseFlags(args); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			secondPass := sudoCmd()
			if args := firstPass.Flags().Args(); len(args) > 0 && !(len(args) == 1 && args[0] == c.CallbackValue) {
				subcmd := subCmd(args[0])
				secondPass.AddCommand(subcmd)
			}
			return carapace.ActionExecute(secondPass)
		}),
	)
}

func sudoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sudo",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()

	cmd.Flags().BoolP("askpass", "A", false, "use a helper program for password prompting")
	cmd.Flags().BoolP("background", "b", false, "run command in the background")
	cmd.Flags().BoolP("bell", "B", false, "ring bell when prompting")
	cmd.Flags().StringP("close-from", "C", "", "close all file descriptors >= num")
	cmd.Flags().BoolP("edit", "e", false, "edit files instead of running a command")
	cmd.Flags().StringP("group", "g", "", "run command as the specified group name or ID")
	cmd.Flags().BoolP("help", "h", false, "display help message and exit")
	cmd.Flags().String("host", "", "run command on host (if supported by plugin)")
	cmd.Flags().BoolP("list", "l", false, "list user's privileges or check a specific command; use twice for longer format")
	cmd.Flags().BoolP("login", "i", false, "run login shell as the target user; a command may also be specified")
	cmd.Flags().BoolP("non-interactive", "n", false, "non-interactive mode, no prompts are used")
	cmd.Flags().StringP("other-user", "U", "", "in list mode, display privileges for user")
	cmd.Flags().StringP("preserve-env", "E", "", "preserve user environment when running command")
	cmd.Flags().BoolP("preserve-groups", "P", false, "preserve group vector instead of setting to target's")
	cmd.Flags().StringP("prompt", "p", "", "use the specified password prompt")
	cmd.Flags().BoolP("remove-timestamp", "K", false, "remove timestamp file completely")
	cmd.Flags().BoolP("reset-timestamp", "k", false, "invalidate timestamp file")
	cmd.Flags().BoolP("set-home", "H", false, "set HOME variable to target user's home dir")
	cmd.Flags().BoolP("shell", "s", false, "run shell as the target user; a command may also be specified")
	cmd.Flags().BoolP("stdin", "S", false, "read password from standard input")
	cmd.Flags().StringP("user", "u", "", "run command (or edit file) as specified user name or ID")
	cmd.Flags().BoolP("validate", "v", false, "update user's timestamp without running a command")
	cmd.Flags().BoolP("version", "V", false, "display version information and exit")

	cmd.Flag("preserve-env").NoOptDefVal = " "

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"group":      os.ActionGroups(),
		"other-user": os.ActionUsers(),
		"preserve-env": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionEnvironmentVariables().Invoke(c).Filter(c.Parts).ToA()
		}),
		"user": os.ActionUsers(),
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
