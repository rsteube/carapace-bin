package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadConfig_kubeadmCmd = &cobra.Command{
	Use:   "kubeadm",
	Short: "Upload the kubeadm ClusterConfiguration to a ConfigMap",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadConfig_kubeadmCmd).Standalone()
	init_phase_uploadConfig_kubeadmCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_uploadConfig_kubeadmCmd.Flags().String("kubeconfig", "/etc/kubernetes/admin.conf", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_uploadConfigCmd.AddCommand(init_phase_uploadConfig_kubeadmCmd)

	carapace.Gen(init_phase_uploadConfig_kubeadmCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
