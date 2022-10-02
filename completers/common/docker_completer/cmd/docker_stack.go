package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Manage Docker stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()

	stackCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	rootCmd.AddCommand(stackCmd)

	carapace.Gen(stackCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})
}
