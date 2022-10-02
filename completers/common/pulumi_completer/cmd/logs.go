package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "[PREVIEW] Show aggregated resource logs for a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()
	logsCmd.PersistentFlags().String("config-file", "", "Use the configuration values in the specified file rather than detecting the file name")
	logsCmd.PersistentFlags().BoolP("follow", "f", false, "Follow the log stream in real time (like tail -f)")
	logsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	logsCmd.PersistentFlags().StringP("resource", "r", "", "Only return logs for the requested resource ('name', 'type::name' or full URN).  Defaults to returning all logs.")
	logsCmd.PersistentFlags().String("since", "1h", "Only return logs newer than a relative duration ('5s', '2m', '3h') or absolute timestamp.  Defaults to returning the last 1 hour of logs.")
	logsCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	rootCmd.AddCommand(logsCmd)

	carapace.Gen(logsCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"resource": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(logsCmd)
		}),
		"stack": action.ActionStacks(logsCmd, action.StackOpts{}),
	})
}
