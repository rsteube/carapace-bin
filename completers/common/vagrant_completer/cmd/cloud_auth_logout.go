package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_auth_logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_auth_logoutCmd).Standalone()

	cloud_authCmd.AddCommand(cloud_auth_logoutCmd)
}
