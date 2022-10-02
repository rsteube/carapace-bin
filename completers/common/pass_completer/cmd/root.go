package cmd

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pass",
	Short: "stores, retrieves, generates, and synchronizes passwords securely",
	Long:  "https://www.passwordstore.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		// TODO support locally installed extension
		if _, err := os.Stat("/usr/lib/password-store/extensions/otp.bash"); os.IsNotExist(err) {
			otpCmd.Hidden = true
		}
	})
}
