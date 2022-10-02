package cmd

import (
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-watch",
	Short: "Watches over your Cargo project’s source",
	Long:  "https://github.com/watchexec/cargo-watch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("B", "B", "", "Inject RUST_BACKTRACE=VALUE (generally you want to set it to 1) into the environment")
	rootCmd.Flags().BoolP("clear", "c", false, "Clear the screen before each run")
	rootCmd.Flags().Bool("debug", false, "Show debug output")
	rootCmd.Flags().StringP("delay", "d", "", "File updates debounce delay in seconds [default: 0.5]")
	rootCmd.Flags().StringArrayP("exec", "x", []string{}, "Cargo command(s) to execute on changes [default: check]")
	rootCmd.Flags().String("features", "", "List of features passed to cargo invocations")
	rootCmd.Flags().BoolP("help", "h", false, "Display this message")
	rootCmd.Flags().StringP("ignore", "i", "", "Ignore a glob/gitignore-style pattern")
	rootCmd.Flags().Bool("ignore-nothing", false, "Ignore nothing, not even target/ and .git/")
	rootCmd.Flags().Bool("no-gitignore", false, "Don’t use .gitignore files")
	rootCmd.Flags().Bool("no-ignore", false, "Don’t use .ignore files")
	rootCmd.Flags().Bool("no-restart", false, "Don’t restart command while it’s still running")
	rootCmd.Flags().BoolP("notify", "N", false, "Send a desktop notification when watchexec notices a change (experimental, behaviour may")
	rootCmd.Flags().Bool("poll", false, "Force use of polling for file changes")
	rootCmd.Flags().Bool("postpone", false, "Postpone first run until a file changes")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress output from cargo-watch itself")
	rootCmd.Flags().StringArrayP("shell", "s", []string{}, "Shell command(s) to execute on changes")
	rootCmd.Flags().String("use-shell", "", "Use a different shell. E.g. --use-shell=bash")
	rootCmd.Flags().BoolP("version", "V", false, "Display version information")
	rootCmd.Flags().StringP("watch", "w", "", "Watch specific file(s) or folder(s) [default: .]")
	rootCmd.Flags().Bool("watch-when-idle", false, "Ignore events emitted while the commands run. Will become default behaviour in 8.0.")
	rootCmd.Flags().Bool("why", false, "Show paths that changed")
	rootCmd.Flags().StringP("workdir", "C", "", "Change working directory before running command [default: crate root]")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exec": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// very simple support for exec completion (breaks easy
			c.CallbackValue = c.CallbackValue + "FAKE_SUFFIX"
			fields := strings.Fields(c.CallbackValue) // TODO dumb split on space

			fields[len(fields)-1] = strings.TrimSuffix(fields[len(fields)-1], "FAKE_SUFFIX")

			c.CallbackValue = fields[len(fields)-1]
			c.Args = make([]string, 0)
			if len(fields) > 1 {
				c.Args = fields[:len(fields)-1]
			}

			prefix := ""
			if len(c.Args) > 0 {
				prefix = strings.Join(c.Args, " ") + " " // TODO won't work when multiple spaces or special characters were used
			}
			return bridge.ActionCarapaceBin("cargo").Invoke(c).Prefix(prefix).ToA()
		}),
		"features": carapace.ActionValues(), // TODO cargo feature completion
		"shell": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// very simple support for exec completion (breaks easy
			c.CallbackValue = c.CallbackValue + "FAKE_SUFFIX"
			fields := strings.Fields(c.CallbackValue) // TODO dumb split on space

			fields[len(fields)-1] = strings.TrimSuffix(fields[len(fields)-1], "FAKE_SUFFIX")

			c.CallbackValue = fields[len(fields)-1]
			c.Args = make([]string, 0)
			if len(fields) > 1 {
				c.Args = fields[:len(fields)-1]
			}

			prefix := ""
			if len(c.Args) > 0 {
				prefix = strings.Join(c.Args, " ") + " " // TODO won't work when multiple spaces or special characters were used
			}

			if len(c.Args) < 1 {
				return carapace.Batch(
					os.ActionPathExecutables(),
					carapace.ActionFiles(),
				).ToA()
			}

			cmd := filepath.Base(c.Args[0])
			if len(c.Args) > 1 {
				c.Args = c.Args[1:]
			} else {
				c.Args = make([]string, 0)
			}
			return bridge.ActionCarapaceBin(cmd).Invoke(c).Prefix(prefix).ToA()
		}),
		"use-shell": os.ActionShells(),
		"watch": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
		"workdir": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).DashCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			cmd := filepath.Base(c.Args[0])
			if len(c.Args) > 1 {
				c.Args = c.Args[1:]
			} else {
				c.Args = make([]string, 0)
			}
			return bridge.ActionCarapaceBin(cmd).Invoke(c).ToA()
		}),
	)

	carapace.Gen(rootCmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(rootCmd.Flag("workdir").Value.String())
	})
}
