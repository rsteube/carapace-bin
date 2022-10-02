package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_demoteCmd = &cobra.Command{
	Use:   "demote",
	Short: "Demote one or more nodes from manager in the swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_demoteCmd).Standalone()

	nodeCmd.AddCommand(node_demoteCmd)

	carapace.Gen(node_demoteCmd).PositionalAnyCompletion(docker.ActionNodes())
}
