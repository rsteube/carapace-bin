package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgrade_node_phaseCmd = &cobra.Command{
	Use:   "phase",
	Short: "Use this command to invoke single phase of the node workflow",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgrade_node_phaseCmd).Standalone()
	upgrade_nodeCmd.AddCommand(upgrade_node_phaseCmd)
}
