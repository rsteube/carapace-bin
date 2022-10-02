package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote one or more nodes to manager in the swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_promoteCmd).Standalone()

	nodeCmd.AddCommand(node_promoteCmd)

	carapace.Gen(node_promoteCmd).PositionalAnyCompletion(docker.ActionNodes())
}
