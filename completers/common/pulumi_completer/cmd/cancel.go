package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel a stack's currently running update, if any",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cancelCmd).Standalone()
	cancelCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	cancelCmd.PersistentFlags().BoolP("yes", "y", false, "Skip confirmation prompts, and proceed with cancellation anyway")
	rootCmd.AddCommand(cancelCmd)

	carapace.Gen(cancelCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(cancelCmd, action.StackOpts{}),
	})

	carapace.Gen(cancelCmd).PositionalCompletion(
		action.ActionStacks(cancelCmd, action.StackOpts{UpdateInProgress: true}),
	)
}
