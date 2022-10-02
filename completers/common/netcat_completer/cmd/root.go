package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netcat",
	Short: "simple Unix utility which reads and writes data across network connections",
	Long:  "https://nc110.sourceforge.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteNc() error {
	rootCmd.Use = "nc"
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("close", "c", false, "close connection on EOF from stdin")
	rootCmd.Flags().BoolP("dont-resolve", "n", false, "numeric-only IP addresses, no DNS")
	rootCmd.Flags().StringP("exec", "e", "", "program to exec after connect")
	rootCmd.Flags().StringP("gateway", "g", "", "source-routing hop point[s], up to 8")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("hexdump", "x", false, "hexdump incoming and outgoing traffic")
	rootCmd.Flags().StringP("interval", "i", "", "delay interval for lines sent, ports scanned")
	rootCmd.Flags().BoolP("listen", "l", false, "listen mode, for inbound connects")
	rootCmd.Flags().StringP("local-port", "p", "", "local port number")
	rootCmd.Flags().StringP("output", "o", "", "output hexdump traffic to FILE (implies -x)")
	rootCmd.Flags().StringP("pointer", "G", "", "source-routing pointer: 4, 8, 12, ...")
	rootCmd.Flags().BoolP("randomize", "r", false, "randomize local and remote ports")
	rootCmd.Flags().StringP("source", "s", "", "local source address (ip or hostname)")
	rootCmd.Flags().BoolP("tcp", "t", false, "TCP mode (default)")
	rootCmd.Flags().BoolP("telnet", "T", false, "answer using TELNET negotiation")
	rootCmd.Flags().StringP("tunnel", "L", "", "forward local port to remote address")
	rootCmd.Flags().BoolP("udp", "u", false, "UDP mode")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose (use twice to be more verbose)")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
	rootCmd.Flags().StringP("wait", "w", "", "timeout for connects and final net reads")
	rootCmd.Flags().BoolP("zero", "z", false, "zero-I/O mode (used for scanning)")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exec":       carapace.ActionFiles(),
		"local-port": net.ActionPorts(),
		"output":     carapace.ActionFiles(),
		"tunnel": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return net.ActionHosts()
			case 1:
				return net.ActionPorts()
			default:
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("L").Changed {
				return carapace.ActionValues()
			}
			return net.ActionHosts()
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("L").Changed {
				return carapace.ActionValues()
			}
			return net.ActionPorts()
		}),
	)
}
