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
	stack_lsCmd.PersistentFlags().BoolP("all", "a", false, "List all stacks instead of just stacks for the current project")
	stack_lsCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	stack_lsCmd.PersistentFlags().StringP("organization", "o", "", "Filter returned stacks to those in a specific organization")
	stack_lsCmd.PersistentFlags().StringP("project", "p", "", "Filter returned stacks to those with a specific project name")
	stack_lsCmd.PersistentFlags().StringP("tag", "t", "", "Filter returned stacks to those in a specific tag (tag-name or tag-name=tag-value)")
	stackCmd.AddCommand(stack_lsCmd)
}
