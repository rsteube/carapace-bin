package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Print the public port for a port binding.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()
	portCmd.Flags().Int("index", 1, "index of the container if service has multiple replicas")
	portCmd.Flags().String("protocol", "tcp", "tcp or udp")
	rootCmd.AddCommand(portCmd)

	carapace.Gen(portCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValues("tcp", "udp"),
	})

	carapace.Gen(portCmd).PositionalCompletion(
		action.ActionServices(portCmd),
		carapace.ActionMessage("TODO: read ports form serice (first agument)"),
	)
}
