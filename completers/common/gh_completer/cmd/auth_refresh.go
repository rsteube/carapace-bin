package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var auth_refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh stored authentication credentials",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_refreshCmd).Standalone()
	auth_refreshCmd.Flags().StringP("hostname", "h", "", "The GitHub host to use for authentication")
	auth_refreshCmd.Flags().StringSliceP("scopes", "s", []string{}, "Additional authentication scopes for gh to have")
	authCmd.AddCommand(auth_refreshCmd)

	carapace.Gen(auth_refreshCmd).FlagCompletion(carapace.ActionMap{
		"hostname": action.ActionConfigHosts(),
		"scopes": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAuthScopes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
