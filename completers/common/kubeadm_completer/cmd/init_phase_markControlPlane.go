package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_markControlPlaneCmd = &cobra.Command{
	Use:   "mark-control-plane",
	Short: "Mark a node as a control-plane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_markControlPlaneCmd).Standalone()
	init_phase_markControlPlaneCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_markControlPlaneCmd.Flags().String("node-name", "", "Specify the node name.")
	init_phaseCmd.AddCommand(init_phase_markControlPlaneCmd)

	carapace.Gen(init_phase_markControlPlaneCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
