package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Print the public port for a port binding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(portCmd).Standalone()

	portCmd.Flags().String("index", "", "index of the container if there are multiple")
	portCmd.Flags().String("protocol", "", "tcp or udp [default: tcp]")
	rootCmd.AddCommand(portCmd)

	carapace.Gen(portCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValues("tcp", "udp"),
	})

	carapace.Gen(portCmd).PositionalCompletion(
		action.ActionServices(portCmd),
		carapace.ActionMessage("TODO: read ports form serice (first agument)"),
	)
}
