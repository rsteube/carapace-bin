package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secrets_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Change the name of a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_renameCmd).Standalone()

	secretsCmd.AddCommand(secrets_renameCmd)

	carapace.Gen(secrets_renameCmd).PositionalCompletion(
		action.ActionSecrets(secrets_renameCmd),
	)
}
