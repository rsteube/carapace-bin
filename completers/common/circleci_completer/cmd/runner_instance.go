package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Operate on runner instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_instanceCmd).Standalone()
	runnerCmd.AddCommand(runner_instanceCmd)
}
