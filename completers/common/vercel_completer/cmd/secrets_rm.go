package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secrets_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secrets_rmCmd).Standalone()

	secretsCmd.AddCommand(secrets_rmCmd)

	carapace.Gen(secrets_rmCmd).PositionalCompletion(
		action.ActionSecrets(secrets_rmCmd),
	)
}
