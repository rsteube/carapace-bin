package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_provider_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a provider entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_updateCmd).Standalone()

	cloud_provider_updateCmd.Flags().StringP("checksum", "c", "", "Checksum of the box for this provider. --checksum-type option is required.")
	cloud_provider_updateCmd.Flags().StringP("checksum-type", "C", "", "Type of checksum used (md5, sha1, sha256, sha384, sha512). --checksum option is required.")
	cloud_providerCmd.AddCommand(cloud_provider_updateCmd)

	carapace.Gen(cloud_provider_updateCmd).FlagCompletion(carapace.ActionMap{
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256", "sha384", "sha512"),
	})

	carapace.Gen(cloud_provider_updateCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
	)
}
