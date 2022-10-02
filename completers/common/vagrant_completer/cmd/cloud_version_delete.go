package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_version_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_deleteCmd).Standalone()

	cloud_version_deleteCmd.Flags().BoolP("force", "f", false, "Force deletion without confirmation")
	cloud_version_deleteCmd.Flags().Bool("no-force", false, "Do not force deletion without confirmation")
	cloud_versionCmd.AddCommand(cloud_version_deleteCmd)

	carapace.Gen(cloud_version_deleteCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
