package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "watchgnupg",
	Short: "Read and print logs from a socket",
	Long:  "https://www.gnupg.org/documentation/manuals/gnupg/watchgnupg.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("force", false, "delete an already existing socket file")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("tcp", false, "listen on a TCP port and optionally on a local socket")
	rootCmd.Flags().Bool("time-only", false, "print only the time; not a full timestamp")
	rootCmd.Flags().Bool("verbose", false, "enable extra informational output")
	rootCmd.Flags().Bool("version", false, "print version of the program and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiParts("://", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("socket", "tcp").Invoke(c).Suffix("://").ToA()
			case 1:
				switch c.Parts[0] {
				case "socket":
					return carapace.ActionFiles()
				case "tcp":
					return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
						switch len(c.Parts) {
						case 0:
							return net.ActionIpv4Addresses()
						case 1:
							return net.ActionPorts()
						default:
							return carapace.ActionValues()
						}
					})
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
