package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_box_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a box entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_box_updateCmd).Standalone()

	cloud_box_updateCmd.Flags().StringP("description", "d", "", "Full description of the box")
	cloud_box_updateCmd.Flags().Bool("no-private", false, "Makes box private")
	cloud_box_updateCmd.Flags().BoolP("private", "p", false, "Makes box private")
	cloud_box_updateCmd.Flags().StringP("short-description", "s", "", "Short description of the box")
	cloud_boxCmd.AddCommand(cloud_box_updateCmd)

	carapace.Gen(cloud_box_updateCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
	)
}
