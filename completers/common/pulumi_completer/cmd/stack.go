package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()
	stackCmd.Flags().BoolP("show-ids", "i", false, "Display each resource's provider-assigned unique ID")
	stackCmd.Flags().Bool("show-name", false, "Display only the stack name")
	stackCmd.Flags().Bool("show-secrets", false, "Display stack outputs which are marked as secret in plaintext")
	stackCmd.Flags().BoolP("show-urns", "u", false, "Display each resource's Pulumi-assigned globally unique URN")
	stackCmd.PersistentFlags().StringP("stack", "s", "", "The name of the stack to operate on. Defaults to the current stack")
	rootCmd.AddCommand(stackCmd)

	carapace.Gen(stackCmd).FlagCompletion(carapace.ActionMap{
		"stack": action.ActionStacks(stackCmd, action.StackOpts{}),
	})
}
