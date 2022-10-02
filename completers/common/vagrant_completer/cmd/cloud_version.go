package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "For managing a Vagrant box's versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_versionCmd).Standalone()

	cloudCmd.AddCommand(cloud_versionCmd)
}
