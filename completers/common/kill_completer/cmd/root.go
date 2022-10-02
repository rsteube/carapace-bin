package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kill",
	Short: "Forcibly terminate a process",
	Long:  "https://linux.die.net/man/1/kill",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "do not restrict the name-to-pid conversion to processes")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("list", "l", "", "list signal names, or convert a signal number to a name")
	rootCmd.Flags().BoolP("pid", "p", false, "print pids without signaling them")
	rootCmd.Flags().StringP("queue", "q", "", "use sigqueue(2), not kill(2), and pass <value> as data")
	rootCmd.Flags().StringP("signal", "s", "", "send this <signal> instead of SIGTERM")
	rootCmd.Flags().BoolP("table", "L", false, "list signal names and numbers")
	rootCmd.Flags().String("timeout", "", "wait up to timeout and send follow-up signal")
	rootCmd.Flags().Bool("verbose", false, "print pids that will be signaled")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("list").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"list":   ps.ActionKillSignals(),
		"signal": ps.ActionKillSignals(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("timeout").Changed {
				// timout takes two arguments so when flag set assume first positional argument as the second flag value
				return ps.ActionKillSignals()
			} else {
				return actionProcessIdsAndNames()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		actionProcessIdsAndNames(),
	)
}

func actionProcessIdsAndNames() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.Batch(
			ps.ActionProcessIds(),
			ps.ActionProcessExecutables(),
		).Invoke(c).Merge().Filter(c.Args).ToA()
	})
}
