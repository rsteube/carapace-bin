package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_unprotectCmd = &cobra.Command{
	Use:   "unprotect",
	Short: "Unprotect resources in a stack's state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_unprotectCmd).Standalone()
	state_unprotectCmd.Flags().Bool("all", false, "Unprotect all resources in the checkpoint")
	state_unprotectCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	state_unprotectCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")
	stateCmd.AddCommand(state_unprotectCmd)

	carapace.Gen(state_unprotectCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(state_unprotectCmd, action.StackOpts{}),
	})

	carapace.Gen(state_unprotectCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(state_unprotectCmd)
		}),
	)
}
