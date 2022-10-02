package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var state_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Renames a resource from a stack's state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_renameCmd).Standalone()
	state_renameCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	state_renameCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompts")
	stateCmd.AddCommand(state_renameCmd)

	carapace.Gen(state_renameCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(state_renameCmd, action.StackOpts{All: true}),
	})

	carapace.Gen(state_renameCmd).PositionalCompletion(
		action.ActionUrns(state_renameCmd),
	)
}
