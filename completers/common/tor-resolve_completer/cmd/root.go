package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tor-resolve",
	Short: "resolve a hostname to an IP address via tor",
	Long:  "https://www.torproject.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Use the SOCKS4a protocol rather than the default SOCKS5 protocol")
	rootCmd.Flags().BoolS("5", "5", false, "Use the SOCKS5 protocol")
	rootCmd.Flags().StringS("p", "p", "", "Override the default SOCKS port without setting the hostname")
	rootCmd.Flags().BoolS("v", "v", false, "Display verbose output")
	rootCmd.Flags().BoolS("x", "x", false, "Perform a reverse lookup")
}
