package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var connect_proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Runs a Consul Connect proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_proxyCmd).Standalone()
	addClientFlags(connect_proxyCmd)
	addServerFlags(connect_proxyCmd)

	connect_proxyCmd.Flags().String("listen", "", "Address to listen for inbound connections to the proxied service.")
	connect_proxyCmd.Flags().Bool("log-json", false, "Output logs in JSON format.")
	connect_proxyCmd.Flags().String("log-level", "", "Specifies the log level.")
	connect_proxyCmd.Flags().String("pprof-addr", "", "Enable debugging via pprof.")
	connect_proxyCmd.Flags().String("proxy-id", "", "The proxy's ID on the local agent.")
	connect_proxyCmd.Flags().Bool("register", false, "Self-register with the local Consul agent.")
	connect_proxyCmd.Flags().String("register-id", "", "ID suffix for the service.")
	connect_proxyCmd.Flags().String("service", "", "Name of the service this proxy is representing.")
	connect_proxyCmd.Flags().String("service-addr", "", "Address of the local service to proxy.")
	connect_proxyCmd.Flags().String("sidecar-for", "", "The ID of a service instance on the local agent that this proxy should become a sidecar for.")
	connect_proxyCmd.Flags().String("upstream", "", "Upstream service to support connecting to.")
	connectCmd.AddCommand(connect_proxyCmd)

	carapace.Gen(connect_proxyCmd).FlagCompletion(carapace.ActionMap{
		// TODO flag completion
		"log-level":   carapace.ActionValues("trace", "debug", "info", "warn", "err").StyleF(style.ForLogLevel),
		"sidecar-for": action.ActionServices(connect_proxyCmd), // TODO local service
	})
}
