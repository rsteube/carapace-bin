package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a stack and its configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_rmCmd).Standalone()
	stack_rmCmd.PersistentFlags().BoolP("force", "f", false, "Forces deletion of the stack, leaving behind any resources managed by the stack")
	stack_rmCmd.PersistentFlags().Bool("preserve-config", false, "Do not delete the corresponding Pulumi.<stack-name>.yaml configuration file for the stack")
	stack_rmCmd.PersistentFlags().BoolP("yes", "y", false, "Skip confirmation prompts, and proceed with removal anyway")
	stackCmd.AddCommand(stack_rmCmd)

	carapace.Gen(stack_rmCmd).PositionalCompletion(
		action.ActionStacks(stack_rmCmd, action.StackOpts{}),
	)
}
