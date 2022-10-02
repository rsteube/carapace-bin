package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cloud_boxCmd = &cobra.Command{
	Use:   "box",
	Short: "For managing a Vagrant box entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_boxCmd).Standalone()

	cloudCmd.AddCommand(cloud_boxCmd)
}
