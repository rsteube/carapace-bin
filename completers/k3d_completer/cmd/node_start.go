package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var node_startCmd = &cobra.Command{
	Use:   "start NODE",
	Short: "Start an existing k3d node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_startCmd).Standalone()

	nodeCmd.AddCommand(node_startCmd)
}
