package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_changeSecretsProviderCmd = &cobra.Command{
	Use:   "change-secrets-provider",
	Short: "Change the secrets provider for the current stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_changeSecretsProviderCmd).Standalone()
	stackCmd.AddCommand(stack_changeSecretsProviderCmd)

	carapace.Gen(stack_changeSecretsProviderCmd).PositionalCompletion(
		action.ActionSecretsProvider(),
	)
}
