package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var node_stopCmd = &cobra.Command{
	Use:   "stop NAME",
	Short: "Stop an existing k3d node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_stopCmd).Standalone()

	nodeCmd.AddCommand(node_stopCmd)
}
