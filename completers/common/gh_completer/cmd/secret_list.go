package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var secret_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List secrets",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_listCmd).Standalone()
	secret_listCmd.Flags().StringP("app", "a", "", "List secrets for a specific application: {actions|codespaces|dependabot}")
	secret_listCmd.Flags().StringP("env", "e", "", "List secrets for an environment")
	secret_listCmd.Flags().StringP("org", "o", "", "List secrets for an organization")
	secret_listCmd.Flags().BoolP("user", "u", false, "List a secret for your user")
	secretCmd.AddCommand(secret_listCmd)

	carapace.Gen(secret_listCmd).FlagCompletion(carapace.ActionMap{
		"app": carapace.ActionValues("actions", "codespaces", "dependabot"),
		"env": action.ActionEnvironments(secret_listCmd),
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})
}
