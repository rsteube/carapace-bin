package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var join_phase_controlPlaneJoin_markControlPlaneCmd = &cobra.Command{
	Use:   "mark-control-plane",
	Short: "Mark a node as a control-plane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(join_phase_controlPlaneJoin_markControlPlaneCmd).Standalone()
	join_phase_controlPlaneJoin_markControlPlaneCmd.Flags().String("config", "", "Path to kubeadm config file.")
	join_phase_controlPlaneJoin_markControlPlaneCmd.Flags().Bool("control-plane", false, "Create a new control plane instance on this node")
	join_phase_controlPlaneJoin_markControlPlaneCmd.Flags().String("node-name", "", "Specify the node name.")
	join_phase_controlPlaneJoinCmd.AddCommand(join_phase_controlPlaneJoin_markControlPlaneCmd)

	carapace.Gen(join_phase_controlPlaneJoin_markControlPlaneCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
