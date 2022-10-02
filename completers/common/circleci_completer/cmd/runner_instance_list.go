package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_instance_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List runner instances",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_instance_listCmd).Standalone()
	runner_instanceCmd.AddCommand(runner_instance_listCmd)
}
