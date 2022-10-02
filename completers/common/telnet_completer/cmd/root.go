package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telnet",
	Short: "User interface to TELNET",
	Long:  "https://linux.die.net/man/1/telnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("binary", "8", false, "use an 8-bit data transmission")
	rootCmd.Flags().BoolP("binary-output", "L", false, "use an 8-bit data transmission for output only")
	rootCmd.Flags().BoolP("debug", "d", false, "turn on debugging")
	rootCmd.Flags().StringP("escape", "e", "", "use CHAR as an escape character")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().BoolP("ipv4", "4", false, "use only IPv4")
	rootCmd.Flags().BoolP("ipv6", "6", false, "use only IPv6")
	rootCmd.Flags().BoolP("login", "a", false, "attempt automatic login")
	rootCmd.Flags().BoolP("no-escape", "E", false, "use no escape character")
	rootCmd.Flags().BoolP("no-login", "K", false, "do not automatically login to the remote system")
	rootCmd.Flags().BoolP("no-rc", "c", false, "do not read the user's .telnetrc file")
	rootCmd.Flags().BoolP("rlogin", "r", false, "use a user-interface similar to rlogin")
	rootCmd.Flags().StringP("trace", "n", "", "record trace information into FILE")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().StringP("user", "l", "", "attempt automatic login as USER")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"trace": carapace.ActionFiles(),
		"user":  os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
