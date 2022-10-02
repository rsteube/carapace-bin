package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connect_exposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Expose a Connect-enabled service through an Ingress gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_exposeCmd).Standalone()
	addClientFlags(connect_exposeCmd)
	addServerFlags(connect_exposeCmd)

	connect_exposeCmd.Flags().String("host", "", "Additional DNS hostname to use for routing to this service.")
	connect_exposeCmd.Flags().String("ingress-gateway", "", "(Required) The name of the ingress gateway service to use.")
	connect_exposeCmd.Flags().String("port", "", "(Required) The listener port to use for the service on the Ingress gateway.")
	connect_exposeCmd.Flags().String("protocol", "", "The protocol for the service. Defaults to 'tcp'.")
	connect_exposeCmd.Flags().String("service", "", "(Required) The name of destination service to expose.")
	connectCmd.AddCommand(connect_exposeCmd)

	carapace.Gen(connect_exposeCmd).FlagCompletion(carapace.ActionMap{
		"protocol": carapace.ActionValues("tcp", "udp"),
		"service":  action.ActionServices(connect_exposeCmd),
	})
}
