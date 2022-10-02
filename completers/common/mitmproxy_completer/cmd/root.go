package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/mitmproxy"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mitmproxy",
	Short: "interactive, SSL/TLS-capable intercepting proxy",
	Long:  "https://mitmproxy.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringArray("allow-hosts", []string{}, "Opposite of --ignore-hosts")
	rootCmd.Flags().Bool("anticache", false, "Strip out request headers that might cause the server to return 304-not-modified")
	rootCmd.Flags().Bool("anticomp", false, "Try to convince servers to send us un-compressed data")
	rootCmd.Flags().String("cert-passphrase", "", "Passphrase for decrypting the private key provided in the --cert option")
	rootCmd.Flags().String("certs", "", "SSL certificates of the form \"[domain=]path\"")
	rootCmd.Flags().StringP("client-replay", "C", "", "Replay client requests from a saved file")
	rootCmd.Flags().Bool("commands", false, "Show all commands and their signatures")
	rootCmd.Flags().String("console-layout", "", "Console layout")
	rootCmd.Flags().Bool("console-layout-headers", false, "Show layout component headers")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("http2", false, "Enable HTTP/2 support")
	rootCmd.Flags().StringArray("ignore-hosts", []string{}, "Ignore host and forward all traffic without processing it")
	rootCmd.Flags().String("intercept", "", "Intercept filter expression")
	rootCmd.Flags().String("key-size", "", "TLS key size for certificates and CA")
	rootCmd.Flags().String("listen-host", "", "Address to bind proxy to")
	rootCmd.Flags().StringP("listen-port", "p", "", "Proxy service port")
	rootCmd.Flags().String("map-local", "", "Map remote resources to a local file")
	rootCmd.Flags().StringP("map-remote", "M", "", "Map remote resources to another remote URL")
	rootCmd.Flags().StringP("mode", "m", "", "Set mode")
	rootCmd.Flags().StringP("modify-body", "B", "", "Replacement pattern")
	rootCmd.Flags().StringP("modify-headers", "H", "", "Header modify pattern")
	rootCmd.Flags().Bool("no-anticache", false, "Do not strip out request headers that might cause the server to return 304-not-modified")
	rootCmd.Flags().Bool("no-anticomp", false, "Do not try to convince servers to send us un-compressed data")
	rootCmd.Flags().Bool("no-console-layout-headers", false, "Do not show layout component headers")
	rootCmd.Flags().Bool("no-http2", false, "Disable HTTP/2 support")
	rootCmd.Flags().Bool("no-rawtcp", false, "Disable raw TCP connections")
	rootCmd.Flags().BoolP("no-server", "n", false, "Do not start a proxy server")
	rootCmd.Flags().Bool("no-server-replay-kill-extra", false, "Do not kill extra requests during replay")
	rootCmd.Flags().Bool("no-server-replay-nopop", false, "Remove flows from server replay state after use")
	rootCmd.Flags().Bool("no-server-replay-refresh", false, "")
	rootCmd.Flags().Bool("no-showhost", false, "Do not use the Host header to construct URLs for display")
	rootCmd.Flags().Bool("no-ssl-insecure", false, "Verify upstream server SSL/TLS certificates")
	rootCmd.Flags().Bool("options", false, "Show all options and their default values")
	rootCmd.Flags().String("proxyauth", "", "Require proxy authentication")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet")
	rootCmd.Flags().Bool("rawtcp", false, "Enable raw TCP connections")
	rootCmd.Flags().StringP("rfile", "r", "", "Read flows from file")
	rootCmd.Flags().StringP("save-stream-file", "w", "", "Stream flows to file as they arrive")
	rootCmd.Flags().StringArrayP("scripts", "s", []string{}, "Execute a script")
	rootCmd.Flags().Bool("server", false, "Start a proxy server")
	rootCmd.Flags().StringP("server-replay", "S", "", "Replay server responses from a saved file")
	rootCmd.Flags().Bool("server-replay-kill-extra", false, "Kill extra requests during replay")
	rootCmd.Flags().Bool("server-replay-nopop", false, "Don't remove flows from server replay state after use")
	rootCmd.Flags().Bool("server-replay-refresh", false, "Refresh server replay responses by adjusting date, expires and last-modified headers")
	rootCmd.Flags().StringArray("set", []string{}, "Set an option")
	rootCmd.Flags().Bool("showhost", false, "Use the Host header to construct URLs for display")
	rootCmd.Flags().BoolP("ssl-insecure", "k", false, "Do not verify upstream server SSL/TLS certificates")
	rootCmd.Flags().String("stickyauth", "", "Set sticky auth filter")
	rootCmd.Flags().String("stickycookie", "", "Set sticky cookie filter")
	rootCmd.Flags().String("tcp-hosts", "", "Generic TCP SSL proxy mode for all hosts that match the pattern")
	rootCmd.Flags().String("upstream-auth", "", "Add HTTP Basic authentication to upstream proxy and reverse proxy requests")
	rootCmd.Flags().BoolP("verbose", "v", false, "Increase log verbosity")
	rootCmd.Flags().Bool("version", false, "show version number and exit")
	rootCmd.Flags().String("view-filter", "", "Limit the view to matching flows")

	// TODO complete flag completions
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"certs": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			if len(c.Parts) < 2 {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
		"client-replay":    carapace.ActionFiles(),
		"console-layout":   mitmproxy.ActionConsoleLayouts(),
		"listen-port":      net.ActionPorts(),
		"mode":             mitmproxy.ActionModes(),
		"modify-body":      mitmproxy.ActionModifyBodyPattern(),
		"modify-headers":   mitmproxy.ActionModifyHeaderPattern(),
		"rfile":            carapace.ActionFiles(),
		"save-stream-file": mitmproxy.ActionAppendableFiles(),
		"scripts":          carapace.ActionFiles(),
		"server-replay":    carapace.ActionFiles(),
		"set": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if splitted := strings.SplitN(c.CallbackValue, "=", 2); len(splitted) > 1 {
				c.CallbackValue = splitted[1]
				return mitmproxy.ActionOptionValues(splitted[0]).Invoke(c).Prefix(splitted[0] + "=").ToA().NoSpace()
			}
			return mitmproxy.ActionOptionNames().NoSpace()
		}),
	})
}
