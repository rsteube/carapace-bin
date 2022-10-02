package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var server_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Manually run the builtin server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_runCmd).Standalone()

	server_runCmd.Flags().Bool("accept-tos", false, "Pass to accept the Terms of Service and Privacy Policy to use the Waypoint URL Service.")
	server_runCmd.Flags().String("advertise-addr", "", "Address to advertise for the server.")
	server_runCmd.Flags().Bool("advertise-tls", false, "If true, the advertised address should be connected to with TLS")
	server_runCmd.Flags().Bool("advertise-tls-skip-verify", false, "Do not verify the TLS certificate presented by the server.")
	server_runCmd.Flags().String("db", "", "Path to the database file.")
	server_runCmd.Flags().Bool("disable-ui", false, "Disable the embedded web interface The default is false.")
	server_runCmd.Flags().String("listen-grpc", "", "Address to bind to for gRPC connections.")
	server_runCmd.Flags().String("listen-http", "", "Address to bind to for HTTP connections.")
	server_runCmd.Flags().String("listen-http-insecure", "", "Address to bind to for insecure HTTP connections.")
	server_runCmd.Flags().String("telemetry-dd-trace-addr", "", "Address of a DataDog agent available to accept traces.")
	server_runCmd.Flags().String("telemetry-oc-agent-addr", "", "Address of an OpenCensus agent or collector.")
	server_runCmd.Flags().Bool("telemetry-oc-agent-insecure", false, "Disables client transport security for the OpenCensus agent exporter's gRPC connection.")
	server_runCmd.Flags().String("telemetry-oc-zpages-addr", "", "If set, Waypoint will run an OpenCensus zPages server at this address.")
	server_runCmd.Flags().String("tls-cert-file", "", "Path to a PEM-encoded certificate file for TLS.")
	server_runCmd.Flags().String("tls-key-file", "", "Path to a PEM-encoded private key file for the TLS certificate specified with -tls-cert-file.")
	server_runCmd.Flags().String("url-api-addr", "", "Address to Waypoint URL service API.")
	server_runCmd.Flags().Bool("url-api-insecure", false, "True if TLS is not enabled for the Waypoint URL service API.")
	server_runCmd.Flags().Bool("url-auto-app-hostname", false, "Whether apps automatically get a hostname on deploy.")
	server_runCmd.Flags().String("url-control-addr", "", "Address to Waypoint URL service control API.")
	server_runCmd.Flags().String("url-control-token", "", "Token for the Waypoint URL server control API.")
	server_runCmd.Flags().Bool("url-enabled", false, "Enable the URL service.")

	addGlobalOptions(server_runCmd)

	serverCmd.AddCommand(server_runCmd)

	carapace.Gen(server_runCmd).FlagCompletion(carapace.ActionMap{
		"db":            carapace.ActionFiles(),
		"tls-cert-file": carapace.ActionFiles(),
		"tls-key-file":  carapace.ActionFiles(),
	})
}
