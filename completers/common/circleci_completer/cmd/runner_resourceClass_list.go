package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_resourceClass_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resource-classes for a namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_resourceClass_listCmd).Standalone()
	runner_resourceClassCmd.AddCommand(runner_resourceClass_listCmd)
}
