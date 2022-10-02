package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_auth_whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Display currently logged in user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_auth_whoamiCmd).Standalone()

	cloud_authCmd.AddCommand(cloud_auth_whoamiCmd)
}
