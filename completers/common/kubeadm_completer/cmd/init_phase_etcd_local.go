package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var init_phase_etcd_localCmd = &cobra.Command{
	Use:   "local",
	Short: "Generate the static Pod manifest file for a local, single-node local etcd instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_etcd_localCmd).Standalone()
	init_phase_etcd_localCmd.Flags().String("cert-dir", "/etc/kubernetes/pki", "The path where to save and store the certificates.")
	init_phase_etcd_localCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_etcd_localCmd.Flags().String("image-repository", "k8s.gcr.io", "Choose a container registry to pull control plane images from")
	init_phase_etcd_localCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\".")
	init_phase_etcdCmd.AddCommand(init_phase_etcd_localCmd)

	carapace.Gen(init_phase_etcd_localCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"patches":  carapace.ActionDirectories(),
	})
}
