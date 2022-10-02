package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadConfigCmd = &cobra.Command{
	Use:   "upload-config",
	Short: "Upload the kubeadm and kubelet configuration to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadConfigCmd).Standalone()
	init_phaseCmd.AddCommand(init_phase_uploadConfigCmd)
}
