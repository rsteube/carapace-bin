package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_tag_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a single stack tag value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_tag_getCmd).Standalone()
	stack_tagCmd.AddCommand(stack_tag_getCmd)
}
