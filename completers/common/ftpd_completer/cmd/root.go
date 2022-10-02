package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ftpd",
	Short: "File Transfer Protocol daemon",
	Long:  "https://linux.die.net/man/8/ftpd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("anonymous-only", "A", false, "server configured for anonymous service only")
	rootCmd.Flags().StringP("auth", "a", "", "use AUTH for authentication")
	rootCmd.Flags().BoolP("daemon", "D", false, "start the ftpd standalone")
	rootCmd.Flags().BoolP("debug", "d", false, "debug mode")
	rootCmd.Flags().BoolP("help", "?", false, "give this help list")
	rootCmd.Flags().BoolP("ipv4", "4", false, "restrict daemon to IPv4")
	rootCmd.Flags().BoolP("ipv6", "6", false, "restrict daemon to IPv6")
	rootCmd.Flags().BoolP("logging", "l", false, "increase verbosity of syslog messages")
	rootCmd.Flags().StringP("max-timeout", "T", "", "set maximum value of timeout allowed")
	rootCmd.Flags().BoolP("no-version", "q", false, "do not display version in banner")
	rootCmd.Flags().Bool("non-rfc2577", false, "neglect RFC 2577 by giving info on missing users")
	rootCmd.Flags().StringP("pidfile", "p", "", "change default location of pidfile")
	rootCmd.Flags().StringP("timeout", "t", "", "set default idle timeout")
	rootCmd.Flags().StringP("umask", "u", "", "set default umask")
	rootCmd.Flags().Bool("usage", false, "give a short usage message")
	rootCmd.Flags().BoolP("version", "V", false, "print program version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"auth": carapace.ActionValuesDescribed(
			"default", "passwd authentication",
			"pam", "using PAM service 'ftp'",
		),
		"pidfile": carapace.ActionFiles(),
		"umask":   fs.ActionFileModesNumeric(),
	})
}
