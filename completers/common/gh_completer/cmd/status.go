package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print information about relevant issues, pull requests, and notifications across repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()
	statusCmd.Flags().StringSliceP("exclude", "e", []string{}, "Comma separated list of repos to exclude in owner/name format")
	statusCmd.Flags().StringP("org", "o", "", "Report status within an organization")
	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"exclude": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionOwnerRepositories(statusCmd)
		}),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})
}
