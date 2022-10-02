package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Export a stack's dependency graph to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_graphCmd).Standalone()
	stack_graphCmd.PersistentFlags().String("dependency-edge-color", "#246C60", "Sets the color of dependency edges in the graph")
	stack_graphCmd.PersistentFlags().Bool("ignore-dependency-edges", false, "Ignores edges introduced by dependency resource relationships")
	stack_graphCmd.PersistentFlags().Bool("ignore-parent-edges", false, "Ignores edges introduced by parent/child resource relationships")
	stack_graphCmd.PersistentFlags().String("parent-edge-color", "#AA6639", "Sets the color of parent edges in the graph")
	stackCmd.AddCommand(stack_graphCmd)

	carapace.Gen(stack_graphCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
