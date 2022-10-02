package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadConfig_kubeletCmd = &cobra.Command{
	Use:   "kubelet",
	Short: "Upload the kubelet component config to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadConfig_kubeletCmd).Standalone()
	init_phase_uploadConfig_kubeletCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_uploadConfig_kubeletCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_uploadConfigCmd.AddCommand(init_phase_uploadConfig_kubeletCmd)

	carapace.Gen(init_phase_uploadConfig_kubeletCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
