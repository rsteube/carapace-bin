package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_resourceClass_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a resource-class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_resourceClass_deleteCmd).Standalone()
	runner_resourceClassCmd.AddCommand(runner_resourceClass_deleteCmd)
}
