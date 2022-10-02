package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an existing stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_renameCmd).Standalone()
	stackCmd.AddCommand(stack_renameCmd)
}
