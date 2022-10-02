package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/circleci_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var context_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_listCmd).Standalone()
	contextCmd.AddCommand(context_listCmd)

	// TODO org-name
	carapace.Gen(context_listCmd).PositionalCompletion(
		action.ActionVcsTypes(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if c.Args[0] == "github" {
				return gh.ActionOrganizations(gh.HostOpts{})
			}
			return carapace.ActionValues()
		}),
	)
}
