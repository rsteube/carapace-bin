package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_auth_loginCmd = &cobra.Command{
	Use:   "login",
	Short: "authenticate with Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_auth_loginCmd).Standalone()

	cloud_auth_loginCmd.Flags().BoolP("check", "c", false, "Checks if currently logged in")
	cloud_auth_loginCmd.Flags().StringP("description", "d", "", "Set description for the Vagrant Cloud token")
	cloud_auth_loginCmd.Flags().StringP("token", "t", "", "Set the Vagrant Cloud token")
	cloud_authCmd.AddCommand(cloud_auth_loginCmd)
}
