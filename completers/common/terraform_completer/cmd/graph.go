package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Generate a Graphviz graph of the steps in an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graphCmd).Standalone()

	graphCmd.Flags().BoolS("draw-cycles", "draw-cycles", false, "Highlight any cycles in the graph with colored edges")
	graphCmd.Flags().StringS("plan", "plan", "", "Render graph using the specified plan file")
	graphCmd.Flags().StringS("type", "type", "", "Type of graph to output")
	rootCmd.AddCommand(graphCmd)

	graphCmd.Flag("plan").NoOptDefVal = " "
	graphCmd.Flag("type").NoOptDefVal = " "

	carapace.Gen(graphCmd).FlagCompletion(carapace.ActionMap{
		"plan": carapace.ActionFiles(),
		"type": carapace.ActionValues("plan", "plan-destroy", "apply", "validate", "input", "refresh"),
	})
}
