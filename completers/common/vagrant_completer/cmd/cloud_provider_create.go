package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_provider_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a provider entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_createCmd).Standalone()

	cloud_provider_createCmd.Flags().StringP("checksum", "c", "", "Checksum of the box for this provider. --checksum-type option is required.")
	cloud_provider_createCmd.Flags().StringP("checksum-type", "C", "", "Type of checksum used (md5, sha1, sha256, sha384, sha512). --checksum option is required.")
	cloud_providerCmd.AddCommand(cloud_provider_createCmd)

	carapace.Gen(cloud_provider_createCmd).FlagCompletion(carapace.ActionMap{
		"checksum-type": carapace.ActionValues("md5", "sha1", "sha256", "sha384", "sha512"),
	})

	carapace.Gen(cloud_provider_createCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		action.ActionProviders(),
	)
}
