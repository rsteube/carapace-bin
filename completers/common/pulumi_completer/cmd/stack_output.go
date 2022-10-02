package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Show a stack's output properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_outputCmd).Standalone()
	stack_outputCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	stack_outputCmd.PersistentFlags().Bool("show-secrets", false, "Display outputs which are marked as secret in plaintext")
	stackCmd.AddCommand(stack_outputCmd)
}
