package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a resource from a stack's state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_deleteCmd).Standalone()
	state_deleteCmd.Flags().Bool("force", false, "Force deletion of protected resources")
	state_deleteCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	state_deleteCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")
	stateCmd.AddCommand(state_deleteCmd)

	carapace.Gen(state_deleteCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(state_deleteCmd, action.StackOpts{}),
	})

	carapace.Gen(state_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUrns(state_deleteCmd)
		}),
	)
}
