package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch the logs of a service or task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_logsCmd).Standalone()

	service_logsCmd.Flags().Bool("details", false, "Show extra details provided to logs")
	service_logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	service_logsCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names in output")
	service_logsCmd.Flags().Bool("no-task-ids", false, "Do not include task IDs in output")
	service_logsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	service_logsCmd.Flags().Bool("raw", false, "Do not neatly format logs")
	service_logsCmd.Flags().String("since", "", "Show logs since timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g. 42m for 42 minutes)")
	service_logsCmd.Flags().String("tail", "", "Number of lines to show from the end of the logs (default \"all\")")
	service_logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	serviceCmd.AddCommand(service_logsCmd)

	carapace.Gen(service_logsCmd).PositionalCompletion(
		docker.ActionServices(),
	)
}
