package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_version_releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Releases a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_releaseCmd).Standalone()

	cloud_version_releaseCmd.Flags().BoolP("force", "f", false, "Release without confirmation")
	cloud_version_releaseCmd.Flags().Bool("no-force", false, "Do not release without confirmation")
	cloud_versionCmd.AddCommand(cloud_version_releaseCmd)

	carapace.Gen(cloud_version_releaseCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
