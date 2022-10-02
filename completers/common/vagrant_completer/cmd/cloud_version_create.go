package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_version_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_createCmd).Standalone()

	cloud_version_createCmd.Flags().StringP("description", "d", "", "A description for this version")
	cloud_versionCmd.AddCommand(cloud_version_createCmd)

	carapace.Gen(cloud_version_createCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
