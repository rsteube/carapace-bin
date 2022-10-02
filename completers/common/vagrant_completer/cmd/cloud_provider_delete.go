package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_provider_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a provider entry on Vagrant Cloud",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_provider_deleteCmd).Standalone()

	cloud_provider_deleteCmd.Flags().BoolP("force", "f", false, "Force deletion of box version provider without confirmation")
	cloud_provider_deleteCmd.Flags().Bool("no-force", false, "Do not force deletion of box version provider without confirmation")
	cloud_providerCmd.AddCommand(cloud_provider_deleteCmd)

	carapace.Gen(cloud_provider_deleteCmd).PositionalCompletion(
		action.ActionCloudBoxSearch(""),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxProviders(c.Args[0])
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionCloudBoxVersions(c.Args[0], c.Args[1])
		}),
	)
}
