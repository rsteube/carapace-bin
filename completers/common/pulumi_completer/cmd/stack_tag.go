package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage stack tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_tagCmd).Standalone()
	stackCmd.AddCommand(stack_tagCmd)
}
