package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch the logs of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_logsCmd).Standalone()

	container_logsCmd.Flags().Bool("details", false, "Show extra details provided to logs")
	container_logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	container_logsCmd.Flags().String("since", "", "Show logs since timestamp (e.g. 2013-01-02T13:23:37) or relative (e.g.")
	container_logsCmd.Flags().String("tail", "", "Number of lines to show from the end of the logs (default \"all\")")
	container_logsCmd.Flags().BoolP("timestamps", "t", false, "Show timestamps")
	container_logsCmd.Flags().String("until", "", "Show logs before a timestamp (e.g. 2013-01-02T13:23:37) or relative")
	containerCmd.AddCommand(container_logsCmd)

	rootAlias(container_logsCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionContainers(),
		)
	})
}
