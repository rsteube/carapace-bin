package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletStartCmd = &cobra.Command{
	Use:   "kubelet-start",
	Short: "Write kubelet settings and (re)start the kubelet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletStartCmd).Standalone()
	init_phase_kubeletStartCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeletStartCmd.Flags().String("cri-socket", "", "Path to the CRI socket to connect. If empty kubeadm will try to auto-detect this value; use this option only if you have more than one CRI installed or if you have non-standard CRI socket.")
	init_phase_kubeletStartCmd.Flags().String("node-name", "", "Specify the node name.")
	init_phaseCmd.AddCommand(init_phase_kubeletStartCmd)

	carapace.Gen(init_phase_kubeletStartCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
