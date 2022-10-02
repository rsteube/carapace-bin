package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_provider_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a box file to Vagrant Cloud for a specific provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_uploadCmd).Standalone()

	cloud_provider_uploadCmd.Flags().BoolP("direct", "D", false, "Upload asset directly to backend storage")
	cloud_provider_uploadCmd.Flags().Bool("no-direct", false, "Do not upload asset directly to backend storage")
	cloud_providerCmd.AddCommand(cloud_provider_uploadCmd)

	carapace.Gen(cloud_provider_uploadCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
		carapace.ActionFiles(".box", ".json"),
	)
}
