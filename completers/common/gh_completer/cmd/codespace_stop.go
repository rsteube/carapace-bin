package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_stopCmd).Standalone()
	codespace_stopCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_stopCmd.Flags().StringP("org", "o", "", "The `login` handle of the organization (admin-only)")
	codespace_stopCmd.Flags().StringP("user", "u", "", "The `username` to stop codespace for (used with --org)")
	codespaceCmd.AddCommand(codespace_stopCmd)

	carapace.Gen(codespace_stopCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"org":       gh.ActionOrganizations(gh.HostOpts{}),
		"user":      gh.ActionUsers(gh.HostOpts{}),
	})
}
