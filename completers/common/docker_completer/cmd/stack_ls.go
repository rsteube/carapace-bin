package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_lsCmd).Standalone()

	stack_lsCmd.Flags().String("format", "", "Pretty-print stacks using a Go template")
	stack_lsCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stackCmd.AddCommand(stack_lsCmd)

	carapace.Gen(stack_lsCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})
}
