package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var secret_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_setCmd).Standalone()
	secret_setCmd.Flags().StringP("app", "a", "", "Set the application for a secret: {actions|codespaces|dependabot}")
	secret_setCmd.Flags().StringP("body", "b", "", "The value for the secret (reads from standard input if not specified)")
	secret_setCmd.Flags().StringP("env", "e", "", "Set deployment `environment` secret")
	secret_setCmd.Flags().StringP("env-file", "f", "", "Load secret names and values from a dotenv-formatted `file`")
	secret_setCmd.Flags().Bool("no-store", false, "Print the encrypted, base64-encoded value instead of storing it on Github")
	secret_setCmd.Flags().StringP("org", "o", "", "Set `organization` secret")
	secret_setCmd.Flags().StringSliceP("repos", "r", []string{}, "List of `repositories` that can access an organization or user secret")
	secret_setCmd.Flags().BoolP("user", "u", false, "Set a secret for your user")
	secret_setCmd.Flags().StringP("visibility", "v", "private", "Set visibility for an organization secret: {all|private|selected}")
	secretCmd.AddCommand(secret_setCmd)

	carapace.Gen(secret_setCmd).FlagCompletion(carapace.ActionMap{
		"app":      carapace.ActionValues("actions", "codespaces", "dependabot"),
		"env":      action.ActionEnvironments(secret_setCmd),
		"env-file": carapace.ActionFiles(),
		"org":      gh.ActionOrganizations(gh.HostOpts{}),
		"repos": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionOwnerRepositories(secret_setCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"visibility": carapace.ActionValues("all", "private", "selected"),
	})

	carapace.Gen(secret_setCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionSecrets(secret_setCmd, action.SecretOpts{
				Org: secret_setCmd.Flag("org").Value.String(),
				Env: secret_setCmd.Flag("env").Value.String(),
			},
			)
		}),
	)
}
