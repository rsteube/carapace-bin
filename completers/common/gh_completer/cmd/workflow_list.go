package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var workflow_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List workflows",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workflow_listCmd).Standalone()
	workflow_listCmd.Flags().BoolP("all", "a", false, "Show all workflows, including disabled workflows")
	workflow_listCmd.Flags().IntP("limit", "L", 50, "Maximum number of workflows to fetch")
	workflowCmd.AddCommand(workflow_listCmd)
}
