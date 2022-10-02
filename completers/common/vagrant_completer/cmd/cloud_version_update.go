package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_version_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_updateCmd).Standalone()

	cloud_version_updateCmd.Flags().StringP("description", "d", "", "A description for this version")
	cloud_versionCmd.AddCommand(cloud_version_updateCmd)

	carapace.Gen(cloud_version_updateCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
