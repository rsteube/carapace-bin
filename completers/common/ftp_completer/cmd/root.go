package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ftp",
	Short: "File Transfer Protocol client",
	Long:  "https://linux.die.net/man/1/ftp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("active", "A", false, "enable active mode transfer")
	rootCmd.Flags().BoolP("debug", "d", false, "enable debugging output")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().BoolP("ipv4", "4", false, "contact IPv4 hosts")
	rootCmd.Flags().BoolP("ipv6", "6", false, "contact IPv6 hosts")
	rootCmd.Flags().StringP("netrc", "N", "", "select a specific initialization file")
	rootCmd.Flags().BoolP("no-edit", "e", false, "disable command line editing")
	rootCmd.Flags().BoolP("no-glob", "g", false, "turn off file name globbing")
	rootCmd.Flags().BoolP("no-login", "n", false, "do not automatically login to the remote system")
	rootCmd.Flags().BoolP("no-prompt", "i", false, "do not prompt during multiple file transfers")
	rootCmd.Flags().BoolP("passive", "p", false, "enable passive mode transfer, default for `pftp'")
	rootCmd.Flags().String("prompt", "", "print a command line PROMPT (optionally), even if not on a tty")
	rootCmd.Flags().BoolP("trace", "t", false, "enable packet tracing")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")

	rootCmd.Flag("prompt").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"netrc": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		net.ActionHosts(),
	)
}
