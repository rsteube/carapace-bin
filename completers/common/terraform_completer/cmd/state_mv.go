package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/terraform_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move an item in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_mvCmd).Standalone()

	state_mvCmd.Flags().BoolS("dry-run", "dry-run", false, "only print out what would've been moved")
	state_mvCmd.Flags().BoolS("ignore-remote-version", "ignore-remote-version", false, "A rare option used for the remote backend only")
	state_mvCmd.Flags().StringS("lock", "lock", "", "Don't hold a state lock during the operation")
	state_mvCmd.Flags().StringS("lock-timeout", "lock-timeout", "", "Duration to retry a state lock")
	stateCmd.AddCommand(state_mvCmd)

	state_mvCmd.Flag("lock").NoOptDefVal = " "
	state_mvCmd.Flag("lock-timeout").NoOptDefVal = " "

	carapace.Gen(state_mvCmd).FlagCompletion(carapace.ActionMap{
		"lock": action.ActionBool(),
	})

	// TODO positional completion
}
