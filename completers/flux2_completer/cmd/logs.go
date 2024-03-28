package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Display formatted logs for Flux components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logsCmd).Standalone()
	logsCmd.Flags().BoolP("all-namespaces", "A", false, "displays logs for objects across all namespaces")
	logsCmd.Flags().String("flux-namespace", "flux-system", "the namespace where the Flux components are running")
	logsCmd.Flags().BoolP("follow", "f", false, "specifies if the logs should be streamed")
	logsCmd.Flags().String("kind", "", "displays errors of a particular toolkit kind e.g GitRepository")
	logsCmd.Flags().String("level", "", "log level, available options are: (debug, info, error)")
	logsCmd.Flags().String("name", "", "specifies the name of the object logs to be displayed")
	logsCmd.Flags().String("since", "", "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().String("since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().Int64("tail", -1, "lines of recent log file to display")
	rootCmd.AddCommand(logsCmd)
}
