package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List the tasks in the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_psCmd).Standalone()

	stack_psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	stack_psCmd.Flags().String("format", "", "Pretty-print tasks using a Go template")
	stack_psCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names")
	stack_psCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	stack_psCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stack_psCmd.Flags().BoolP("quiet", "q", false, "Only display task IDs")
	stackCmd.AddCommand(stack_psCmd)

	carapace.Gen(stack_psCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})
}
