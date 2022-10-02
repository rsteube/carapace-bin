package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_resourceClassCmd = &cobra.Command{
	Use:   "resource-class",
	Short: "Operate on runner resource-classes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_resourceClassCmd).Standalone()
	runnerCmd.AddCommand(runner_resourceClassCmd)
}
