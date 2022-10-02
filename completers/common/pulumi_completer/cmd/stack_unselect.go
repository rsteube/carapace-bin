package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_unselectCmd = &cobra.Command{
	Use:   "unselect",
	Short: "Resets stack selection from the current workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_unselectCmd).Standalone()
	stackCmd.AddCommand(stack_unselectCmd)
}
