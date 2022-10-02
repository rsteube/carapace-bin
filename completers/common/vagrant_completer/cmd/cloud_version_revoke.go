package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_version_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Revokes a version entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_version_revokeCmd).Standalone()

	cloud_version_revokeCmd.Flags().BoolP("force", "f", false, "Force revocation without confirmation")
	cloud_version_revokeCmd.Flags().Bool("no-force", false, "Do not force revocation without confirmation")
	cloud_versionCmd.AddCommand(cloud_version_revokeCmd)

	carapace.Gen(cloud_version_revokeCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], "")
		}),
	)
}
