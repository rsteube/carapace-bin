package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty stack with the given name, ready for updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_initCmd).Standalone()
	stack_initCmd.PersistentFlags().String("copy-config-from", "", "The name of the stack to copy existing config from")
	stack_initCmd.PersistentFlags().String("secrets-provider", "default", "The type of the provider that should be used to encrypt and decrypt secrets")
	stackCmd.AddCommand(stack_initCmd)

	carapace.Gen(stack_initCmd).FlagCompletion(carapace.ActionMap{
		"copy-config-from": action.ActionStacks(stack_initCmd, action.StackOpts{All: true}),
		"secrets-provider": action.ActionSecretsProvider(),
	})
}
