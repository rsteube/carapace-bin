package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Runs a Consul agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agentCmd).Standalone()

	agentCmd.Flags().String("advertise", "", "Sets the advertise address to use.")
	agentCmd.Flags().String("advertise-wan", "", "Sets address to advertise on WAN instead of -advertise address.")
	agentCmd.Flags().StringArray("allow-write-http-from", []string{}, "Only allow write endpoint calls from given network.")
	agentCmd.Flags().String("alt-domain", "", "Alternate domain to use for DNS interface.")
	agentCmd.Flags().String("bind", "", "Sets the bind address for cluster communication.")
	agentCmd.Flags().Bool("bootstrap", false, "Sets server to bootstrap mode.")
	agentCmd.Flags().String("bootstrap-expect", "", "Sets server to expect bootstrap mode.")
	agentCmd.Flags().String("check_output_max_size", "", "Sets the maximum output size for checks on this agent")
	agentCmd.Flags().String("client", "", "Sets the address to bind for client access.")
	agentCmd.Flags().String("config-dir", "", "Path to a directory to read configuration files from.")
	agentCmd.Flags().StringArray("config-file", []string{}, "Path to a file in JSON or HCL format with a matching file extension.")
	agentCmd.Flags().String("config-format", "", "Config files are in this format irrespective of their extension.")
	agentCmd.Flags().String("data-dir", "", "Path to a data directory to store agent state.")
	agentCmd.Flags().String("default-query-time", "", "the amount of time a blocking query will wait before Consul will force a response.")
	agentCmd.Flags().Bool("dev", false, "Starts the agent in development mode.")
	agentCmd.Flags().Bool("disable-host-node-id", false, "generate a random node ID.")
	agentCmd.Flags().Bool("disable-keyring-file", false, "Disables the backing up of the keyring to a file.")
	agentCmd.Flags().String("dns-port", "", "DNS port to use.")
	agentCmd.Flags().String("domain", "", "Domain to use for DNS interface.")
	agentCmd.Flags().Bool("enable-local-script-checks", false, "Enables health check scripts from configuration file.")
	agentCmd.Flags().Bool("enable-script-checks", false, "Enables health check scripts.")
	agentCmd.Flags().String("encrypt", "", "Provides the gossip encryption key.")
	agentCmd.Flags().String("grpc-port", "", "Sets the gRPC API port to listen on (currently needed for Envoy xDS only).")
	agentCmd.Flags().StringArray("hcl", []string{}, "hcl config fragment. Can be specified multiple times.")
	agentCmd.Flags().String("http-port", "", "Sets the HTTP API port to listen on.")
	agentCmd.Flags().String("https-port", "", "Sets the HTTPS API port to listen on.")
	agentCmd.Flags().StringArray("join", []string{}, "Address of an agent to join at start time. Can be specified multiple times.")
	agentCmd.Flags().StringArray("join-wan", []string{}, "Address of an agent to join -wan at start time. Can be specified multiple times.")
	agentCmd.Flags().String("log-file", "", "Path to the file the logs get written to")
	agentCmd.Flags().Bool("log-json", false, "Output logs in JSON format.")
	agentCmd.Flags().String("log-level", "", "Log level of the agent.")
	agentCmd.Flags().String("log-rotate-bytes", "", "Maximum number of bytes that should be written to a log file")
	agentCmd.Flags().String("log-rotate-duration", "", "Time after which log rotation needs to be performed")
	agentCmd.Flags().String("log-rotate-max-files", "", "Maximum number of log file archives to keep")
	agentCmd.Flags().String("node", "", "Name of this node. Must be unique in the cluster.")
	agentCmd.Flags().StringArray("node-meta", []string{}, "An arbitrary metadata key/value pair for this node, of the format `key:value`.")
	agentCmd.Flags().Bool("non-voting-server", false, "(Enterprise-only) DEPRECATED: -read-replica should be used instead")
	agentCmd.Flags().String("pid-file", "", "Path to file to store agent PID.")
	agentCmd.Flags().String("primary-gateway", "", "Address of a mesh gateway in the primary datacenter to use to bootstrap WAN federation at start time with retries enabled.")
	agentCmd.Flags().String("protocol", "", "Sets the protocol version.")
	agentCmd.Flags().String("raft-protocol", "", "Sets the Raft protocol version. Defaults to latest.")
	agentCmd.Flags().Bool("read-replica", false, "(Enterprise-only) This flag is used to make the server not participate in the Raft quorum")
	agentCmd.Flags().StringArray("recursor", []string{}, "Address of an upstream DNS server. Can be specified multiple times.")
	agentCmd.Flags().Bool("rejoin", false, "Ignores a previous leave and attempts to rejoin the cluster.")
	agentCmd.Flags().String("retry-interval", "", "Time to wait between join attempts.")
	agentCmd.Flags().String("retry-interval-wan", "", "Time to wait between join -wan attempts.")
	agentCmd.Flags().StringArray("retry-join", []string{}, "Address of an agent to join at start time with retries enabled.")
	agentCmd.Flags().StringArray("retry-join-wan", []string{}, "Address of an agent to join -wan at start time with retries enabled.")
	agentCmd.Flags().String("retry-max", "", "Maximum number of join attempts. Defaults to 0, which will retry indefinitely.")
	agentCmd.Flags().String("retry-max-wan", "", "Maximum number of join -wan attempts. Defaults to 0, which will retry indefinitely.")
	agentCmd.Flags().String("segment", "", "(Enterprise-only) Sets the network segment to join.")
	agentCmd.Flags().StringArray("serf-lan-allowed-cidrs", []string{}, "Networks allowed for Serf LAN.")
	agentCmd.Flags().String("serf-lan-bind", "", "Address to bind Serf LAN listeners to.")
	agentCmd.Flags().String("serf-lan-port", "", "Sets the Serf LAN port to listen on.")
	agentCmd.Flags().StringArray("serf-wan-allowed-cidrs", []string{}, "Networks allowed for Serf WAN (other datacenters).")
	agentCmd.Flags().String("serf-wan-bind", "", "Address to bind Serf WAN listeners to.")
	agentCmd.Flags().String("serf-wan-port", "", "Sets the Serf WAN port to listen on.")
	agentCmd.Flags().Bool("server", false, "Switches agent to server mode.")
	agentCmd.Flags().String("server-port", "", "Sets the server port to listen on.")
	agentCmd.Flags().Bool("syslog", false, "Enables logging to syslog.")
	agentCmd.Flags().Bool("ui", false, "Enables the built-in static web UI server.")
	agentCmd.Flags().String("ui-content-path", "", "Sets the external UI path to a string.")
	agentCmd.Flags().String("ui-dir", "", "Path to directory containing the web UI resources.")
	rootCmd.AddCommand(agentCmd)

	carapace.Gen(agentCmd).FlagCompletion(carapace.ActionMap{
		"config-dir":    carapace.ActionDirectories(),
		"config-file":   carapace.ActionFiles(),
		"config-format": carapace.ActionValues("json", "hcl"),
		"data-dir":      carapace.ActionDirectories(),
		"log-file":      carapace.ActionFiles(),
		"log-level":     carapace.ActionValues("trace", "debug", "info", "warn", "err").StyleF(style.ForLogLevel),
		"pid-file":      carapace.ActionFiles(),
		"ui-dir":        carapace.ActionDirectories(),
	})
}
