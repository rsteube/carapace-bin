package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_replaceProviderCmd = &cobra.Command{
	Use:   "replace-provider",
	Short: "Replace provider in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_replaceProviderCmd).Standalone()

	state_replaceProviderCmd.Flags().BoolS("auto-approve", "auto-approve", false, "Skip interactive approval.")
	state_replaceProviderCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only.")
	state_replaceProviderCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation.")
	state_replaceProviderCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock.")
	stateCmd.AddCommand(state_replaceProviderCmd)

	state_replaceProviderCmd.Flag("lock").NoOptDefVal = " "
	state_replaceProviderCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(state_replaceProviderCmd).FlagCompletion(carapace.ActionMap{
		"lock":         action.ActionBool(),
		"lock-timeout": action.ActionBool(),
	})

	// TODO positional completion
}
