package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cloud_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search Vagrant Cloud for available boxes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud_searchCmd).Standalone()

	cloud_searchCmd.Flags().Bool("auth", false, "Authenticate with Vagrant Cloud if required before searching")
	cloud_searchCmd.Flags().BoolP("json", "j", false, "Formats results in JSON")
	cloud_searchCmd.Flags().StringP("limit", "l", "", "Max number of search results Default: 25")
	cloud_searchCmd.Flags().Bool("no-auth", false, "Do not authenticate with Vagrant Cloud if required before searching")
	cloud_searchCmd.Flags().StringP("order", "o", "", "Order to display results ('desc' or 'asc') Default: 'desc'")
	cloud_searchCmd.Flags().String("page", "", "The page to display Default: 1")
	cloud_searchCmd.Flags().StringP("provider", "p", "", "Filter search results to a single provider. Defaults to all.")
	cloud_searchCmd.Flags().BoolP("short", "s", false, "Shows a simple list of box names")
	cloud_searchCmd.Flags().String("sort-by", "", "Field to sort results on (created, downloads, updated) Default: downloads")
	cloudCmd.AddCommand(cloud_searchCmd)

	carapace.Gen(cloud_searchCmd).FlagCompletion(carapace.ActionMap{
		"order":    carapace.ActionValues("desc", "asc"),
		"provider": action.ActionProviders(),
	})
}
