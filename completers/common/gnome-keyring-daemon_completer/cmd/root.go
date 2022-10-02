package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gnome-keyring-daemon",
	Short: "The Gnome Keyring Daemon",
	Long:  "https://wiki.gnome.org/Projects/GnomeKeyring/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("components", "c", "", "The optional components to run")
	rootCmd.Flags().StringP("control-directory", "C", "", "The directory for sockets and control data")
	rootCmd.Flags().BoolP("daemonize", "d", false, "Run as a daemon")
	rootCmd.Flags().BoolP("foreground", "f", false, "Run in the foreground")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().BoolP("login", "l", false, "Run by PAM for a user login. Read login password from stdin")
	rootCmd.Flags().BoolP("replace", "r", false, "Replace the daemon for this desktop login environment.")
	rootCmd.Flags().BoolP("start", "s", false, "Start a dameon or initialize an already running daemon.")
	rootCmd.Flags().Bool("unlock", false, "Prompt for login keyring password, or read from stdin")
	rootCmd.Flags().BoolP("version", "V", false, "Show the version number and exit.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"components": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("pkcs11", "secrets", "ssh").Invoke(c).Filter(c.Parts).ToA()
		}),
		"control-directory": carapace.ActionDirectories(),
	})
}
