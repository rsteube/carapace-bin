package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "For various provider actions with Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_providerCmd).Standalone()

	cloudCmd.AddCommand(cloud_providerCmd)
}
