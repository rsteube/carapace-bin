package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "torsocks",
	Short: "Shell wrapper to simplify the use of the torsocks(8) library to transparently torify an application",
	Long:  "https://www.torproject.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "Set debug mode.")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("isolate", "i", false, "Automatic tor isolation. Can't be used with -u/-p")
	rootCmd.Flags().StringP("pass", "p", "", "Password for the SOCKS5 authentication")
	rootCmd.Flags().StringP("port", "P", "", "Specify Tor port")
	rootCmd.Flags().Bool("shell", false, "Spawn a torified shell")
	rootCmd.Flags().StringP("user", "u", "", "Username for the SOCKS5 authentication")
	rootCmd.Flags().Bool("version", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			carapace.ActionValuesDescribed(
				"on", "adds torsocks to the LD_PRELOAD environment variable",
				"off", "removes torsocks from the LD_PRELOAD environment variable",
				"show", "Show the current value of the LD_PRELOAD environment variable",
				"sh", "Show the current value of the LD_PRELOAD environment variable",
			),
		).ToA(),
	)
}
