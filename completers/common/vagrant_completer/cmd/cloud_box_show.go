package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_box_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays a boxes attributes on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_box_showCmd).Standalone()

	cloud_box_showCmd.Flags().Bool("auth", false, "Authenticate with Vagrant Cloud if required before searching")
	cloud_box_showCmd.Flags().Bool("no-auth", false, "Do not authenticate with Vagrant Cloud if required before searching")
	cloud_boxCmd.AddCommand(cloud_box_showCmd)

	carapace.Gen(cloud_box_showCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
	)
}
