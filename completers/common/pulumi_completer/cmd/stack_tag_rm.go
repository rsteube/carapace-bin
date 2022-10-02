package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_tag_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a stack tag",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_tag_rmCmd).Standalone()
	stack_tagCmd.AddCommand(stack_tag_rmCmd)
}
