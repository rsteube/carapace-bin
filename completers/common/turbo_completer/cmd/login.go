package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your Vercel account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().String("api", "", "api url")
	loginCmd.Flags().String("sso-team", "", "attempt to authenticate to the specified team using SSO")
	loginCmd.Flags().String("url", "", "login url")
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flag("api").NoOptDefVal = " "
	loginCmd.Flag("sso-team").NoOptDefVal = " "
	loginCmd.Flag("url").NoOptDefVal = " "
}
