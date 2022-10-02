package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/consul_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connect_envoyCmd = &cobra.Command{
	Use:   "envoy",
	Short: "Runs or Configures Envoy as a Connect proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_envoyCmd).Standalone()
	addClientFlags(connect_envoyCmd)

	connect_envoyCmd.Flags().String("address", "", "LAN address to advertise in the gateway service registration")
	connect_envoyCmd.Flags().String("admin-access-log-path", "", "The path to write the access log for the administration server.")
	connect_envoyCmd.Flags().String("admin-bind", "", "The address:port to start envoy's admin server on.")
	connect_envoyCmd.Flags().String("bind-address", "", "Bind address to use instead of the default binding rules given as `<name>=<ip>:<port>` pairs.")
	connect_envoyCmd.Flags().Bool("bootstrap", false, "Generate the bootstrap.json but don't exec envoy")
	connect_envoyCmd.Flags().String("deregister-after-critical", "", "The amount of time the gateway services health check can be failing before being deregistered")
	connect_envoyCmd.Flags().String("envoy-binary", "", "The full path to the envoy binary to run.")
	connect_envoyCmd.Flags().String("envoy-version", "", "Sets the envoy-version that the envoy binary has.")
	connect_envoyCmd.Flags().Bool("expose-servers", false, "Expose the servers for WAN federation via this mesh gateway")
	connect_envoyCmd.Flags().String("gateway", "", "The type of gateway to register.")
	connect_envoyCmd.Flags().String("grpc-addr", "", "Set the agent's gRPC address and port (in http(s)://host:port format).")
	connect_envoyCmd.Flags().Bool("mesh-gateway", false, "Configure Envoy as a Mesh Gateway.")
	connect_envoyCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	connect_envoyCmd.Flags().Bool("no-central-config", false, "Require that the command run on the same agent as the proxy will.")
	connect_envoyCmd.Flags().Bool("omit-deprecated-tags", false, "Disable deprecated tags.")
	connect_envoyCmd.Flags().String("proxy-id", "", "The proxy's ID on the local agent.")
	connect_envoyCmd.Flags().Bool("register", false, "Register a new gateway service before configuring and starting Envoy")
	connect_envoyCmd.Flags().String("service", "", "Service name to use for the registration")
	connect_envoyCmd.Flags().String("sidecar-for", "", "The ID of a service instance on the local agent that this proxy should become a sidecar for.")
	connect_envoyCmd.Flags().String("wan-address", "", "WAN address to advertise in the gateway service registration.")
	connectCmd.AddCommand(connect_envoyCmd)

	carapace.Gen(connect_envoyCmd).FlagCompletion(carapace.ActionMap{
		// TODO complete flags
		"envoy-binary": carapace.ActionFiles(),
		"sidecar-for":  action.ActionServices(connect_envoyCmd), // TODO must be local service
	})
}
