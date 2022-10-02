package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_kubeletFinalize_experimentalCertRotationCmd = &cobra.Command{
	Use:   "experimental-cert-rotation",
	Short: "Enable kubelet client certificate rotation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeletFinalize_experimentalCertRotationCmd).Standalone()
	init_phase_kubeletFinalize_experimentalCertRotationCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_kubeletFinalize_experimentalCertRotationCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeletFinalizeCmd.AddCommand(init_phase_kubeletFinalize_experimentalCertRotationCmd)

	carapace.Gen(init_phase_kubeletFinalize_experimentalCertRotationCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
