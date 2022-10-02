package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run a proxy to the Kubernetes API server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxyCmd).Standalone()

	proxyCmd.Flags().String("accept-hosts", "", "Regular expression for hosts that the proxy should accept.")
	proxyCmd.Flags().String("accept-paths", "", "Regular expression for paths that the proxy should accept.")
	proxyCmd.Flags().String("address", "", "The IP address on which to serve on.")
	proxyCmd.Flags().String("api-prefix", "", "Prefix to serve the proxied API under.")
	proxyCmd.Flags().Bool("disable-filter", false, "If true, disable request filtering in the proxy. This is dangerous, and can leave you vulnerable to ")
	proxyCmd.Flags().String("keepalive", "", "keepalive specifies the keep-alive period for an active network connection. Set to 0 to disable keep")
	proxyCmd.Flags().StringP("port", "p", "", "The port on which to run the proxy. Set to 0 to pick a random port.")
	proxyCmd.Flags().String("reject-methods", "", "Regular expression for HTTP methods that the proxy should reject (example --reject-methods='POST,PUT")
	proxyCmd.Flags().String("reject-paths", "", "Regular expression for paths that the proxy should reject. Paths specified here will be rejected eve")
	proxyCmd.Flags().StringP("unix-socket", "u", "", "Unix socket on which to run the proxy.")
	proxyCmd.Flags().StringP("www", "w", "", "Also serve static files from the given directory under the specified prefix.")
	proxyCmd.Flags().StringP("www-prefix", "P", "", "Prefix to serve static files under, if static file directory is specified.")
	rootCmd.AddCommand(proxyCmd)

	carapace.Gen(proxyCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
