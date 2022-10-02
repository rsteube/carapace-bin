package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var node_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(node_updateCmd).Standalone()

	node_updateCmd.Flags().String("availability", "", "Availability of the node (\"active\"|\"pause\"|\"drain\")")
	node_updateCmd.Flags().String("label-add", "", "Add or update a node label (key=value)")
	node_updateCmd.Flags().String("label-rm", "", "Remove a node label if exists")
	node_updateCmd.Flags().String("role", "", "Role of the node (\"worker\"|\"manager\")")
	nodeCmd.AddCommand(node_updateCmd)

	carapace.Gen(node_updateCmd).FlagCompletion(carapace.ActionMap{
		"availability": carapace.ActionValues("active", "pause", "drain"),
		"role":         carapace.ActionValues("worker", "manager"),
	})

	carapace.Gen(node_updateCmd).PositionalCompletion(
		docker.ActionNodes(),
	)
}
