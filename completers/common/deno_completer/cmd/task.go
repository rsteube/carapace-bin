package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/deno_completer/cmd/action"
	"github.com/spf13/cobra"
)

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Run a task defined in the configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(taskCmd).Standalone()

	taskCmd.Flags().StringP("config", "c", "", "The configuration file can be used to configure different aspects of deno")
	taskCmd.Flags().BoolP("help", "h", false, "Print help information")
	taskCmd.Flags().Bool("no-config", false, "Disable automatic loading of the configuration file.")
	taskCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	taskCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	rootCmd.AddCommand(taskCmd)

	carapace.Gen(taskCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})

	carapace.Gen(taskCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionTasks(taskCmd.Flag("config").Value.String())
		}),
	)
}
