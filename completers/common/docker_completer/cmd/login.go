package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to a Docker registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loginCmd).Standalone()

	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.Flags().Bool("password-stdin", false, "Take the password from stdin")
	loginCmd.Flags().StringP("username", "u", "", "Username")
	rootCmd.AddCommand(loginCmd)
}
