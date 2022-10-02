package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_authCmd = &cobra.Command{
	Use:   "auth",
	Short: "For various authorization operations on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_authCmd).Standalone()

	cloudCmd.AddCommand(cloud_authCmd)
}
